package models

type ToolShort struct {
	ID               string   `json:"id" yaml:"id"`
	Name             string   `json:"name" yaml:"name"`
	Kind             string   `json:"kind" yaml:"kind"`
	Languages        []string `json:"languages" yaml:"languages"`
	InstalledVersion *string  `json:"installed_version" yaml:"installed_version"`
}

type ToolDescribe struct {
	ID               string   `json:"id" yaml:"id"`
	Name             string   `json:"name" yaml:"name"`
	Description      string   `json:"description" yaml:"description"`
	Url              string   `json:"url" yaml:"url"`
	LatestVersion    *string  `json:"latest_version" yaml:"latest_version"`
	InstalledVersion *string  `json:"installed_version" yaml:"installed_version"`
	Dependencies     []string `json:"dependencies" yaml:"dependencies"`
	Languages        []string `json:"languages" yaml:"languages"`
	Kind             string   `json:"kind" yaml:"kind"`
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
	Kind             string   `json:"kind" yaml:"kind"`
	Recipe           Recipe   `json:"recipe" yaml:"recipe"`
}

func ToShort(m Tool) ToolShort {
	return ToolShort{
		ID:               m.ID,
		Name:             m.Name,
		Kind:             m.Kind,
		Languages:        m.Languages,
		InstalledVersion: m.InstalledVersion,
	}
}

func ToDescribe(m Tool) ToolDescribe {
	return ToolDescribe{
		ID:               m.ID,
		Name:             m.Name,
		Description:      m.Description,
		Url:              m.Url,
		LatestVersion:    m.LatestVersion,
		InstalledVersion: m.InstalledVersion,
		Dependencies:     m.Dependencies,
		Languages:        m.Languages,
		Kind:             m.Kind,
	}
}
