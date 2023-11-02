package models

type ToolShort struct {
	ID               string   `json:"id" yaml:"id"`
	Name             string   `json:"name" yaml:"name"`
	Languages        []string `json:"languages" yaml:"languages"`
	InstalledVersion string   `json:"installed_version" yaml:"installed_version"`
}

type Tool struct {
	ID               string   `json:"id" yaml:"id"`
	Name             string   `json:"name" yaml:"name"`
	Description      string   `json:"description" yaml:"description"`
	Url              string   `json:"url" yaml:"url"`
	LatestVersion    *string  `json:"latest_version,omitempty" yaml:"latest_version,omitempty"`
	InstalledVersion *string  `json:"installed_version,omitempty" yaml:"installed_version,omitempty"`
	Dependencies     []string `json:"dependencies" yaml:"dependencies"`
	Languages        []string `json:"languages" yaml:"languages"`
	Recipe           Recipe   `json:"recipe" yaml:"recipe"`
}
