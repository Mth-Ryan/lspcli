package recipes

import (
	"github.com/Mth-Ryan/lspcli/pkg/models"
	"github.com/mitchellh/mapstructure"
)

type GoRecipeParser struct{}

func NewGoRecipeParser() *GoRecipeParser {
	return &GoRecipeParser{}
}

func (p *GoRecipeParser) Parse(raw map[string]any) (*models.GoRecipe, error) {
	recipe := new(models.GoRecipe)
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
