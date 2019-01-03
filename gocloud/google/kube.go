package google

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type Service struct {
	Status *Status
}

type Status struct {
	LoadBalancer LoadBalancer
}

type LoadBalancer struct {
	Ingress []Ingress
}

type Ingress struct {
	IP string
}

func (g *Gcloud) GetKubeKreds(state string, app *App) error {
	zone := g.ClusterZone.Value
	if zone == "" {
		return fmt.Errorf("empty or missing cluster_zone in %s", state)
	}

	// Run on Kubernetes.
	log.Printf("Deploying to %s...", g.ClusterName.Value)
	getCreds := g.Cmd("container", "clusters", "get-credentials", "--zone", zone, g.ClusterName.Value)
	getCreds.Stderr = os.Stderr
	if err := getCreds.Run(); err != nil {
		return fmt.Errorf("getting credentials with %v: %v", getCreds.Args, err)
	}
	return nil
}
func (g *Gcloud) KubeDeploy(dir string, app *App) error {

	kubeCmds := [][]string{
		{"kubectl", "apply", "-f", filepath.Join(dir, app.Name+".yaml")},
		// Force pull the latest image.
		{"kubectl", "scale", "--replicas", "0", "deployment/" + app.Name},
		{"kubectl", "scale", "--replicas", "1", "deployment/" + app.Name},
	}
	for _, kcmd := range kubeCmds {
		cmd := exec.Command(kcmd[0], kcmd[1:]...)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("running %v: %v", cmd.Args, err)
		}
	}
	return nil
}

func (g *Gcloud) GetLoadBalancer(app *App) error {
	// Wait for endpoint then print it.
	log.Printf("Waiting for load balancer...")

	for {
		outb, err := Runb("kubectl", "get", "service", app.Name, "-o", "json")
		if err != nil {
			return err
		}
		var s Service

		if err := json.Unmarshal(outb, &s); err != nil {
			return fmt.Errorf("parsing JSON output: %v", err)
		}
		i := s.Status.LoadBalancer.Ingress
		if len(i) == 0 || i[0].IP == "" {
			dt := time.Second
			log.Printf("No ingress returned in %s. Trying again in %v", outb, dt)
			time.Sleep(dt)
			continue
		}
		endpoint := i[0].IP
		log.Printf("Deployed at http://%s:8080", endpoint)
		break
	}
	return nil
}
