package google

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type Key struct {
	PrivateKeyID string `json:"private_key_id"`
}

type Db struct {
	ServiceAccount, Instance, Name, Password, SchemaPath string
}

func (g *Gcloud) ProvisionDB(db *Db, app *App) error {
	log.Printf("Downloading Docker images...")
	const mySQLImage = "mysql:5.6"
	cloudSQLProxyImage := "gcr.io/cloudsql-docker/gce-proxy:1.11"
	images := []string{mySQLImage, cloudSQLProxyImage}
	for _, img := range images {
		if _, err := Runb("docker", "pull", img); err != nil {
			return err
		}
	}

	log.Printf("Getting connection string from database metadata...")

	dbConnStr := g.Cmd("sql", "instances", "describe", "--format", "value(connectionName)", db.Instance).Run()

	// Create a temporary directory to hold the service account key.
	// We resolve all symlinks to avoid Docker on Mac issues, see
	// https://github.com/google/go-cloud/issues/110.
	serviceAccountVolDir, err := ioutil.TempDir("", app.Name+"-service-acct")
	if err != nil {
		return fmt.Errorf("creating temp dir to hold service account key: %v", err)
	}
	serviceAccountVolDir, err = filepath.EvalSymlinks(serviceAccountVolDir)
	if err != nil {
		return fmt.Errorf("evaluating any symlinks: %v", err)
	}
	defer os.RemoveAll(serviceAccountVolDir)
	log.Printf("Created %v", serviceAccountVolDir)

	// Furnish a new service account key.
	if err := g.Cmd("iam", "service-accounts", "keys", "create", "--iam-account="+db.ServiceAccount, serviceAccountVolDir+"/key.json").Run(); err != nil {
		return fmt.Errorf("creating new service account key: %v", err)
	}
	keyJSONb, err := ioutil.ReadFile(filepath.Join(serviceAccountVolDir, "key.json"))
	if err != nil {
		return fmt.Errorf("reading key.json file: %v", err)
	}
	var k Key
	if err := json.Unmarshal(keyJSONb, &k); err != nil {
		return fmt.Errorf("parsing key.json: %v", err)
	}
	serviceAccountKeyID := k.PrivateKeyID
	defer func() {
		if err := g.Cmd("iam", "service-accounts", "keys", "delete", "--iam-account", db.ServiceAccount, serviceAccountKeyID).Run(); err != nil {

			log.Printf("deleting service account key: %v", err)
		}
	}()
	log.Printf("Created service account key %s", serviceAccountKeyID)

	log.Printf("Starting Cloud SQL proxy...")
	proxyContainerID, err := Runb("docker", "run", "--detach", "--rm", "--volume", serviceAccountVolDir+":/creds", "--publish", "3306", cloudSQLProxyImage, "/cloud_sql_proxy", "-instances", dbConnStr.Error()+"=tcp:0.0.0.0:3306", "-credential_file=/creds/key.json")
	if err != nil {
		return err
	}
	defer func() {
		if _, err := Runb("docker", "kill", string(proxyContainerID)); err != nil {
			log.Printf("failed to kill docker container for proxy: %v", err)
		}
	}()

	log.Print("Sending schema to database...")
	mySQLCmd := fmt.Sprintf(`mysql --wait -h"$PROXY_PORT_3306_TCP_ADDR" -P"$PROXY_PORT_3306_TCP_PORT" -uroot -p'%s' '%s'`, db.Password, db.Name)
	connect := exec.Command("docker", "run", "--rm", "--interactive", "--link", string(proxyContainerID)+":proxy", mySQLImage, "sh", "-c", mySQLCmd)
	schema, err := os.Open(db.SchemaPath)
	if err != nil {
		return err
	}
	defer schema.Close()
	connect.Stdin = schema
	connect.Stderr = os.Stderr
	if err := connect.Run(); err != nil {
		return fmt.Errorf("running %v: %v", connect.Args, err)
	}

	return nil
}
