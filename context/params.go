package context

type Params struct {
	Name    string
	Package string
}

func (c *Ctx) NewParams() *Params {
	return &Params{
		Name:    c.Build.AppName,
		Package: c.GetImportPath(),
	}
}
