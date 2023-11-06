package runtime

import "path"

type Conf struct {
	rootPath string
}

func NewConf(runtimePath string) *Conf {
	return &Conf{
		rootPath: runtimePath,
	}
}

func (c *Conf) RootPath() string {
	return c.rootPath
}

func (c *Conf) ToolsPath() string {
	return path.Join(c.RootPath(), "tools")
}

func (c *Conf) InstallsPath() string {
	return path.Join(c.RootPath(), "installs")
}

func (c *Conf) InstallsListPath() string {
	return path.Join(c.RootPath(), "installs-list")
}

func (c *Conf) BinPath() string {
	return path.Join(c.RootPath(), "bin")
}

func (c *Conf) CachePath() string {
	return path.Join(c.RootPath(), "cache")
}
