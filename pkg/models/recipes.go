package models

type RecipeKind = string

const (
	RECIPE_GIT_RELEASE = "git_release"
	RECIPE_GO          = "go"
	RECIPE_NPM         = "npm"
)

type Recipe struct {
	Kind            RecipeKind                        `json:"kind" yaml:"kind"`
	ContextReplaces *(map[string](map[string]string)) `json:"context_replaces" yaml:"context_replaces"`
}

type GoRecipe struct {
	Recipe
	Package string `json:"package" yaml:"package"`
}

type NpmRecipe struct {
	Recipe
	Package string `json:"package" yaml:"package"`
}

type GitReleaseRecipe struct {
	Recipe
	Repository      string `json:"repository" yaml:"repository"`
	BinaryInnerPath string `json:"binary_inner_path" yaml:"binary_inner_path"`
	Package         string `json:"package" yaml:"package"`
}
