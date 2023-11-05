package providers

import (
	"fmt"

	"github.com/Mth-Ryan/lspcli/pkg/models"
	"github.com/Mth-Ryan/lspcli/pkg/recipes"
)

type GoProvider struct {
	tool         models.Tool
	recipeParser *recipes.GoRecipeParser
}

func NewGoProvider(tool models.Tool) Provider {
	return &GoProvider{
		tool:         tool,
		recipeParser: recipes.NewGoRecipeParser(),
	}
}

func (e *GoProvider) getRecipe() (*models.GoRecipe, error) {
	return e.recipeParser.Parse(e.tool.Recipe)
}

func (e *GoProvider) Install() error {
	recipe, err := e.getRecipe()
	if err != nil {
		return err
	}

	return fmt.Errorf(recipe.Package)
}

func (e *GoProvider) Update() error {
	return nil
}

func (e *GoProvider) Remove() error {
	return nil
}

func (e *GoProvider) LatestVersion() (string, error) {
	return "", nil
}

func (e *GoProvider) InstaledVersion() (string, error) {
	return "", nil
}
