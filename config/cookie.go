package config

import (
	"github.com/Masterminds/sprig"
	"github.com/gofunct/common/errors"
	"github.com/gofunct/common/logging"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"os"
	"text/template"
)

var (
	messenger = logging.NewMessenger()
	logger, _ = zap.NewDevelopment()
)

type GoCookieConfig struct {
	Config   *viper.Viper
	Data     map[string]interface{}
	Os       *afero.OsFs
	Template *template.Template
}

func NewGoCookieConfig() (*GoCookieConfig, error) {
	var err error
	g := &GoCookieConfig{
		Config: viper.New(),
		Data:   make(map[string]interface{}),
		Os:     &afero.OsFs{},
	}

	g.Config.SetConfigName("gogen")
	g.Config.AddConfigPath(".")
	g.Config.AutomaticEnv()
	g.Config.SetDefault("app", "app")
	g.Config.SetDefault("service", "user")
	g.Config.SetDefault("domain", "")
	g.Config.SetDefault("author", "Coleman Word")
	g.Config.SetDefault("summary", "the default gogen configuration settings!")
	g.Config.SetDefault("git_username", "gofunct")
	g.Config.SetDefault("docker_username", "colemanword")
	g.Config.SetDefault("transport", "grpc")
	g.Config.SetDefault("listen", ":8080")
	g.Config.SetDefault("makefile", "y")
	g.Config.SetDefault(".gitignore", "y")
	g.Config.SetDefault("kubernetes", "y")
	g.Config.SetDefault("postgres", "y")

	g.Data["App"] = g.Config.GetString("app")
	g.Data["Service"] = g.Config.GetString("service")
	g.Data["Domain"] = g.Config.GetString("domain")
	g.Data["Author"] = g.Config.GetString("author")
	g.Data["Summary"] = g.Config.GetString("summary")
	g.Data["GitUserName"] = g.Config.GetString("git_username")
	g.Data["DockerUserName"] = g.Config.GetString("docker_username")
	g.Data["Transport"] = g.Config.GetString("transport")
	g.Data["Listen"] = g.Config.GetString("listen")

	g.Template = template.Must(template.ParseGlob("templates/*")).Funcs(sprig.TxtFuncMap())

	// If a config file is found, read it in.
	if err = g.Config.ReadInConfig(); err != nil {
		log.Printf("%s, %s", "failed to read config file, writing defaults...", err)
		if err = g.Config.WriteConfigAs("gogen.yaml"); err != nil {
			return nil, err
		}

	} else {
		log.Printf("%s, %s", "Using config file:", g.Config.ConfigFileUsed())
		if err = g.Config.WriteConfig(); err != nil {
			return nil, err
		}
	}

	return g, err
}

func (g *GoCookieConfig) ExecTemplates() error {
	switch {
	case g.Template == nil:
		return errors.New("must initialize template before execution")
	case g.Data == nil:
		return errors.New("must initialize data before execution")
	case g.Config == nil:
		return errors.New("must initialize viper config before execution")
	}

	if err := g.Template.Execute(os.Stdout, g.Data); err != nil {
		return err
	}
	return nil
}
