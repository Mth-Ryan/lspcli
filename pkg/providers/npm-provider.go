package providers

import (
	"fmt"

	"github.com/Mth-Ryan/lspcli/pkg/models"
	"github.com/Mth-Ryan/lspcli/pkg/recipes"
)

type NpmProvider struct {
	tool         models.Tool
	recipeParser *recipes.NpmRecipeParser
}

func NewNpmProvider(tool models.Tool) Provider {
	return &NpmProvider{
		tool,
		recipes.NewNpmRecipeParser(),
	}
}

func (e *NpmProvider) getRecipe() (*models.NpmRecipe, error) {
	return e.recipeParser.Parse(e.tool.Recipe)
}

func (e *NpmProvider) Install() error {
	recipe, err := e.getRecipe()
	if err != nil {
		return err
	}

	return fmt.Errorf(recipe.Package)
}

func (e *NpmProvider) Update() error {
	return nil
}

func (e *NpmProvider) Remove() error {
	return nil
}

func (e *NpmProvider) LatestVersion() (string, error) {
	return "", nil
}

func (e *NpmProvider) InstaledVersion() (string, error) {
	return "", nil
}
