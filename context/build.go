package context

// Build is a container for the application build information.
type Build struct {
	AppName     string
	Version     string
	Description string
	Service     string
	Domain      string
	Cloud       string
	Revision    string
	BuildDate   string
	Dev         bool
}

func BuildFromConfig(config *Config) Build {
	return Build{
		AppName:     config.V.GetString("Build.AppName"),
		Version:     config.V.GetString("Build.Version"),
		Description: config.V.GetString("Build.Description"),
		Service:     config.V.GetString("Build.Service"),
		Domain:      config.V.GetString("Build.Domain"),
		Cloud:       config.V.GetString("Build.Cloud"),
		Revision:    config.V.GetString("Build.Revision"),
		BuildDate:   config.V.GetString("Build.BuildDate"),
		Dev:         config.V.GetBool("Build.Dev"),
	}
}
