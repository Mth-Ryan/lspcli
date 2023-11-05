package recipes

import (
	"github.com/Mth-Ryan/lspcli/pkg/models"
	"github.com/mitchellh/mapstructure"
)

type NpmRecipeParser struct{}

func NewNpmRecipeParser() *NpmRecipeParser {
	return &NpmRecipeParser{}
}

func (p *NpmRecipeParser) Parse(raw map[string]any) (*models.NpmRecipe, error) {
	recipe := new(models.NpmRecipe)
	if err := mapstructure.Decode(raw, recipe); err != nil {
		return recipe, err
	}

	itemParser := NewTemplateParser(recipe.ContextReplaces)

	newPackage, err := itemParser.Parse(recipe.Package)
	if err != nil {
		return recipe, err
	}
	recipe.Package = newPackage

	return recipe, nil
}
