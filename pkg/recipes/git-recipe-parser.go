package recipes

import (
	"github.com/Mth-Ryan/lspcli/pkg/models"
	"github.com/mitchellh/mapstructure"
)

type GitReleaseRecipeParser struct{}

func NewGitReleaseRecipeParser() *GitReleaseRecipeParser {
	return &GitReleaseRecipeParser{}
}

func (p *GitReleaseRecipeParser) Parse(raw map[string]any) (*models.GitReleaseRecipe, error) {
	recipe := new(models.GitReleaseRecipe)
	if err := mapstructure.Decode(raw, recipe); err != nil {
		return recipe, err
	}

	itemParser := NewTemplateParser(recipe.ContextReplaces)

	newPackage, err := itemParser.Parse(recipe.Package)
	if err != nil {
		return recipe, err
	}
	recipe.Package = newPackage

	newBinaryPath, err := itemParser.Parse(recipe.BinaryInnerPath)
	if err != nil {
		return recipe, err
	}
	recipe.BinaryInnerPath = newBinaryPath

	return recipe, nil
}
