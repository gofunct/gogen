package google

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

type App struct {
	Name string
}

func (a *App) GetManifest() string {
	return a.Name + ".yaml.in"
}

func (a *App) ReadManifest(dir string, g *Gcloud) (string, error) {
	gbyin, err := ioutil.ReadFile(filepath.Join(dir, a.GetManifest()))
	if err != nil {
		return "", err
	}

	gby := string(gbyin)
	replacements := map[string]string{
		"{{IMAGE}}":             g.GetImageName(a),
		"{{bucket}}":            g.Bucket.Value,
		"{{database_instance}}": g.DatabaseInstance.Value,
		"{{database_region}}":   g.DatabaseRegion.Value,
		"{{motd_var_config}}":   g.MotdVarConfig.Value,
		"{{motd_var_name}}":     g.MotdVarName.Value,
	}
	for old, new := range replacements {
		gby = strings.Replace(gby, old, new, -1)
	}

	return gby, nil
}
