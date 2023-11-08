package providers

import (
	"fmt"

	"github.com/Mth-Ryan/lspcli/pkg/handlers"
	"github.com/Mth-Ryan/lspcli/pkg/loggers"
	"github.com/Mth-Ryan/lspcli/pkg/models"
	"github.com/Mth-Ryan/lspcli/pkg/recipes"
	"github.com/Mth-Ryan/lspcli/pkg/runtime"
)

type NpmProvider struct {
	tool                models.Tool
	logger              loggers.Logger
	recipeParser        *recipes.NpmRecipeParser
	npmHandler          *handlers.NpmHandler
	installsListHandler handlers.InstallsListHandler
}

func NewNpmProvider(runConf runtime.Conf, tool models.Tool, logger loggers.Logger) Provider {
	return &NpmProvider{
		tool:                tool,
		logger:              logger,
		recipeParser:        recipes.NewNpmRecipeParser(),
		npmHandler:          handlers.NewNpmHandler(logger),
		installsListHandler: handlers.NewJsonInstallsHandler(runConf),
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

	latestVersion, err := e.npmHandler.GetLatestVersion(recipe.Package)
	if err != nil {
		return fmt.Errorf("Unable to fetch the latest version from npm")
	}

	err = e.npmHandler.Install(recipe.Package)
	if err != nil {
		return fmt.Errorf("Unable to install latest version with npm")
	}

	return e.installsListHandler.SetVersion(e.tool.ID, &latestVersion)
}

func (e *NpmProvider) Update() error {
	recipe, err := e.getRecipe()
	if err != nil {
		return err
	}

	latestVersion, err := e.npmHandler.GetLatestVersion(recipe.Package)
	if err != nil {
		return fmt.Errorf("Unable to fetch the latest version from npm")
	}

	err = e.npmHandler.Update(recipe.Package)
	if err != nil {
		return fmt.Errorf("Unable to update to latest version with npm")
	}

	return e.installsListHandler.SetVersion(e.tool.ID, &latestVersion)
}

func (e *NpmProvider) Remove() error {
	recipe, err := e.getRecipe()
	if err != nil {
		return err
	}

	err = e.npmHandler.Remove(recipe.Package)
	if err != nil {
		return fmt.Errorf("Unable to remove the package with npm")
	}

	return e.installsListHandler.SetVersion(e.tool.ID, nil)
}

func (e *NpmProvider) LatestVersion() (string, error) {
	recipe, err := e.getRecipe()
	if err != nil {
		return "", err
	}

	return e.npmHandler.GetLatestVersion(recipe.Package)
}

func (e *NpmProvider) InstalledVersion() (string, error) {
	installs, err := e.installsListHandler.GetInstalls()
	if err != nil {
		return "", err
	}

	if version, ok := installs[e.tool.ID]; ok {
		if version != nil {
			return *version, nil
		}
	}

	return "", nil
}
