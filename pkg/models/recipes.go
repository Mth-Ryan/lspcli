package models

type RecipeKind = string

const (
	RECIPE_GIT_RELEASE = "git_release"
	RECIPE_GO          = "go"
	RECIPE_NPM         = "npm"
)

type RecipeContextReplaces = *(map[string](map[string]string))

type Recipe struct {
	Kind            RecipeKind            `json:"kind" yaml:"kind" mapstructure:"kind"`
	ContextReplaces RecipeContextReplaces `json:"context_replaces" yaml:"context_replaces" mapstructure:"context_replaces"`
}

type GoRecipe struct {
	Recipe  `mapstructure:",squash"`
	Package string `json:"package" yaml:"package" mapstructure:"package"`
}

type NpmRecipe struct {
	Recipe  `mapstructure:",squash"`
	Package string `json:"package" yaml:"package" mapstructure:"package"`
}

type GitReleaseRecipe struct {
	Recipe          `mapstructure:",squash"`
	Repository      string `json:"repository" yaml:"repository" mapstructure:"repository"`
	BinaryName      string `json:"binary_name" yaml:"binary_name" mapstructure:"binary_name"`
	BinaryInnerPath string `json:"binary_inner_path" yaml:"binary_inner_path" mapstructure:"binary_inner_path"`
	Package         string `json:"package" yaml:"package" mapstructure:"package"`
}
