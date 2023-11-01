package models

type ToolShort struct {
	ID               string `json:"id" yaml:"id"`
	Name             string `json:"name" yaml:"name"`
	InstalledVersion string `json:"installed_version" yaml:"installed_version"`
}

type Tool struct {
	ID               string   `json:"id" yaml:"id"`
	Name             string   `json:"name" yaml:"name"`
	LatestVersion    string   `json:"latest_version" yaml:"latest_version"`
	InstalledVersion string   `json:"installed_version" yaml:"installed_version"`
	Dependencies     []string `json:"dependencies" yaml:"dependencies"`
	Recipe           Recipe   `json:"recipe" yaml:"recipe"`
}
