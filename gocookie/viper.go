package gocookie

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"github.com/gofunct/common/logging"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"path/filepath"
	"text/template"
	"text/template/parse"
)

var (
	logger = logging.Base()
	Cookie *GoCookieConfig
)

func init() {
	Cookie = NewGoCookieConfig()
}

type GoCookieConfig struct {
	Config 	*viper.Viper
	Data 	map[string]interface{}
	Os 		*afero.OsFs
	Buffer 	*bytes.Buffer
	Tree 	*parse.Tree
}


func  NewGoCookieConfig() *GoCookieConfig {
	g := &GoCookieConfig{
		Config:    viper.New(),
		Data:      make(map[string]interface{}),
		Os:        &afero.OsFs{},
		Buffer:    new(bytes.Buffer),
	}

	{
		g.Config.SetConfigName("gocookiecutter")
		g.Config.AddConfigPath(".")
		g.Config.AutomaticEnv()
		g.Config.SetDefault("app", "app")
		g.Config.SetDefault("serg.Configice", "user")
		g.Config.SetDefault("domain", "")
		g.Config.SetDefault("author", "Coleman Word")
		g.Config.SetDefault("summary", "the default gocookiecutter configuration settings!")
		g.Config.SetDefault("git_username", "gofunct")
		g.Config.SetDefault("docker_username", "colemanword")
		g.Config.SetDefault("transport", "grpc")
		g.Config.SetDefault("listen", ":8080")
		g.Config.SetDefault("makefile", "y")
		g.Config.SetDefault(".gitignore", "y")
		g.Config.SetDefault("kubernetes", "y")
		g.Config.SetDefault("postgres", "y")
	}

	{
		g.Data["App"] = g.Config.GetString("app")
		g.Data["Service"] = g.Config.GetString("service")
		g.Data["Domain"] = g.Config.GetString("domain")
		g.Data["Author"] = g.Config.GetString("author")
		g.Data["Summary"] = g.Config.GetString("summary")
		g.Data["GitUserName"] = g.Config.GetString("git_username")
		g.Data["DockerUserName"] = g.Config.GetString("docker_username")
		g.Data["Transport"] = g.Config.GetString("transport")
		g.Data["Listen"] = g.Config.GetString("listen")
	}


		// If a config file is found, read it in.
		if err := g.Config.ReadInConfig(); err != nil {
			logger.Debug("failed to read config file, writing defaults...")
			if err := g.Config.WriteConfigAs("gocookiecutter.yaml"); err != nil {
				logger.Fatal("failed to write config")
			}

		} else {
			logger.Debug("Using config file-->",  g.Config.ConfigFileUsed())
			if err := g.Config.WriteConfig(); err != nil {
				logger.Fatal("failed to write config file")
			}
		}
		g.Tree = parse.New("gocookiecutter", g.Data, sprig.TxtFuncMap())

	return g
}


func (g *GoCookieConfig) ExecTemplates(tmplPath string) {
	if g.Tree == nil {
		logger.Fatal("must initialize parse tree before generating template")
	}

	pattern := filepath.Join(tmplPath, "*.tmpl")
	var err error
	var tmpl = &template.Template{
		Tree: g.Tree,
	}
	tmpl, err = tmpl.ParseGlob(pattern)

	if err != nil {
		logger.Fatal("failed to parse template glob", err)
	}

	if err := tmpl.Execute(g.Buffer, g.Data); err != nil {
		logger.Fatal("failed to generate files from templates", err)
	}
}
