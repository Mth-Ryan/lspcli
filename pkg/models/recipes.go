package models

type RecipeKind = string

const (
	RECIPE_GIT_RELEASE = "git_release"
	RECIPE_GO          = "go"
	RECIPE_NPM         = "npm"
)

type Recipe struct {
	Kind             RecipeKind        `json:"kind" yaml:"kind"`
	GoRecipe         *GoRecipe         `json:"go,omitempty" yaml:"go,omitempty"`
	NpmRecipe        *NpmRecipe        `json:"npm,omitempty" yaml:"npm,omitempty"`
	GitReleaseRecipe *GitReleaseRecipe `json:"git_release,omitempty" yaml:"git_release,omitempty"`
}

type GoRecipe struct {
	Package string `json:"package" yaml:"package"`
}

type NpmRecipe struct {
	Package string `json:"package" yaml:"package"`
}

type GitReleaseRecipe struct {
	Repository      string `json:"repository" yaml:"repository"`
	BinaryInnerPath string `json:"binary_inner_path" yaml:"binary_inner_path"`
}
