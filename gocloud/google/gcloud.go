package google

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type TfItem struct {
	Sensitive bool
	Type      string
	Value     string
}

type Gcloud struct {
	Project          TfItem
	ClusterName      TfItem `json:"cluster_name"`
	ClusterZone      TfItem `json:"cluster_zone"`
	Bucket           TfItem
	DatabaseInstance TfItem `json:"database_instance"`
	DatabaseRegion   TfItem `json:"database_region"`
	MotdVarConfig    TfItem `json:"motd_var_config"`
	MotdVarName      TfItem `json:"motd_var_name"`
}

func NewGcloudFromStatePath(tfStatePath string) (*Gcloud, error) {
	tfStateb, err := Runb("terraform", "output", "-state", tfStatePath, "-json")
	if err != nil {
		return nil, err
	}
	var gcloud Gcloud

	if err := json.Unmarshal(tfStateb, &gcloud); err != nil {
		return nil, fmt.Errorf("parsing terraform state JSON: %v", err)
	}

	return &gcloud, nil
}

func (g *Gcloud) Cmd(args ...string) *exec.Cmd {
	args = append([]string{"--quiet", "--project", g.Project.Value}, args...)
	cmd := exec.Command("gcloud", args...)
	cmd.Env = append(cmd.Env, os.Environ()...)
	cmd.Stderr = os.Stderr
	return cmd
}

func (g *Gcloud) Runb(args ...string) (stdout string, err error) {
	stdoutb, err := Runb(args...)
	return strings.TrimSpace(string(stdoutb)), err
}

func (g *Gcloud) Deploy(dir, tfStatePath string, app *App) error {

	imageName := g.GetImageName(app)

	tempDir, err := ioutil.TempDir("", app.Name+"-k8s-")
	if err != nil {
		return fmt.Errorf("making temp dir: %v", err)
	}

	defer os.RemoveAll(tempDir)

	gby, err := app.ReadManifest(tempDir, g)

	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(filepath.Join(tempDir, app.Name+".yaml"), []byte(gby), 0666); err != nil {
		return err
	}

	if err := g.BuildImage(tempDir, app); err != nil {
		return err
	}

	if err := g.Cmd("builds", "submit", "-t", imageName, filepath.Join(dir)).Run(); err != nil {
		return err
	}

	if err = g.GetKubeKreds(tfStatePath, app); err != nil {
		return err
	}

	if err = g.KubeDeploy(tempDir, app); err != nil {
		return err
	}

	if err := g.GetLoadBalancer(app); err != nil {
		return err
	}

	return nil
}

func (g *Gcloud) BuildImage(dir string, app *App) error {
	proj := strings.Replace(g.Project.Value, ":", "/", -1)
	imageName := fmt.Sprintf("gcr.io/%s/"+app.Name, proj)

	// Build app Docker image.
	log.Printf("Building %s...", imageName)
	build := exec.Command("go", "build", "-o", app.Name)
	env := append(build.Env, "GOOS=linux", "GOARCH=amd64")
	env = append(env, os.Environ()...)
	build.Env = env
	absDir, err := filepath.Abs(dir)
	if err != nil {
		return fmt.Errorf("getting abs path to "+app.Name+" dir (%s): %v", dir, err)
	}
	build.Dir = absDir
	build.Stderr = os.Stderr
	if err := build.Run(); err != nil {
		return fmt.Errorf("building "+app.Name+" app by running %v: %v", build.Args, err)
	}
	return nil
}

func (g *Gcloud) FormatProject() string {
	return strings.Replace(g.Project.Value, ":", "/", -1)
}

func (g *Gcloud) GetImageName(app *App) string {
	return fmt.Sprintf("gcr.io/%s/"+app.Name, g.FormatProject())
}
