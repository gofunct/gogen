package context

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	V *viper.Viper
}

func NewConfig() *Config {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(".")
	v.AutomaticEnv()
	v.SetDefault("Build.AppName", "app")
	v.SetDefault("Build.Version", "0.1.1")
	v.SetDefault("Build.Domain", "")
	v.SetDefault("Build.Service", "example")
	v.SetDefault("Build.Description", "the default gogen configuration settings!")
	v.SetDefault("Build.Cloud", "Google")
	v.SetDefault("Build.Revision", "")
	v.SetDefault("Build.BuildDate", "")
	v.SetDefault("Build.Dev", true)

	return &Config{
		V: v,
	}
}

func (c *Config) ReadConfig() error {
	// If a config file is found, read it in.
	if err := c.V.ReadInConfig(); err != nil {
		log.Printf("%s, %s", "failed to read config file, writing defaults...", err)
		if err = c.V.WriteConfigAs("config.yaml"); err != nil {
			return err
		}

	} else {
		log.Printf("%s, %s", "Using config file:", c.V.ConfigFileUsed())
		if err = c.V.WriteConfig(); err != nil {
			return err
		}
	}
	return nil
}
