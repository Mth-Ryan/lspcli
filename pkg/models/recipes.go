package models

type RecipeKind = string

const (
	RECIPE_GIT_RELEASE = "git"
	RECIPE_GO          = "go"
	RECIPE_NPM         = "npm"
)

type Recipe struct {
	Kind             RecipeKind        `json:"kind" yaml:"kind"`
	GoRecipe         *GoRecipe         `json:"go,omitempty" yaml:"go,omitempty"`
	NpmRecipe        *NpmRecipe        `json:"npm,omitempty" yaml:"npm,omitempty"`
	GitReleaseRecipe *GitReleaseRecipe `json:"git,omitempty" yaml:"git,omitempty"`
}

type GoRecipe struct {
	Package string `json:"package" yaml:"package"`
}

type NpmRecipe struct {
	Package string `json:"package" yaml:"package"`
}

type GitReleaseRecipe struct {
	DownloadUrl     string `json:"download_url" yaml:"download_url"`
	BinaryInnerPath string `json:"binary_inner_path" yaml:"binary_inner_path"`
}
