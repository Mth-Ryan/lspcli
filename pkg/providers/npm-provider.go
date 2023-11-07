package providers

import (
	"fmt"

	"github.com/Mth-Ryan/lspcli/pkg/loggers"
	"github.com/Mth-Ryan/lspcli/pkg/models"
	"github.com/Mth-Ryan/lspcli/pkg/recipes"
	"github.com/Mth-Ryan/lspcli/pkg/runtime"
)

type NpmProvider struct {
	tool         models.Tool
	recipeParser *recipes.NpmRecipeParser
}

func NewNpmProvider(runConf runtime.Conf, tool models.Tool, logger loggers.Logger) Provider {
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

func (e *NpmProvider) InstalledVersion() (string, error) {
	return "", nil
}
