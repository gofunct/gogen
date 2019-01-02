package build

// Build is a container for the application build information.
type Build struct {
	AppName   string
	Version   string
	Description string
	Revision  string
	BuildDate string
}
