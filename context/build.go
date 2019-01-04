package context

import "go/build"

var (
	BuildContext = build.Default
)

// Build is a container for the application build information.
type Build struct {
	AppName       string
	Version       string
	Revision      string
	BuildDate     string
	Description   string
	Service       string
	Domain        string
	Cloud         string
	Dev           bool
	StorageBucket string
}

func BuildFromConfig(config *Config) Build {
	return Build{
		AppName:       config.V.GetString("Build.AppName"),
		Version:       config.V.GetString("Build.Version"),
		Description:   config.V.GetString("Build.Description"),
		Service:       config.V.GetString("Build.Service"),
		Domain:        config.V.GetString("Build.Domain"),
		Cloud:         config.V.GetString("Build.Cloud"),
		Dev:           config.V.GetBool("Build.Dev"),
		StorageBucket: config.V.GetString("Build.Bucket"),
	}
}
