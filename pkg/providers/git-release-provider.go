package providers

import (
	"fmt"

	"github.com/Mth-Ryan/lspcli/pkg/handlers"
	"github.com/Mth-Ryan/lspcli/pkg/loggers"
	"github.com/Mth-Ryan/lspcli/pkg/models"
	"github.com/Mth-Ryan/lspcli/pkg/recipes"
)

type GitReleaseProvider struct {
	tool         models.Tool
	recipeParser *recipes.GitReleaseRecipeParser
	handler      *handlers.GithubReleaseHandler
}

func (e *GitReleaseProvider) getRecipe() (*models.GitReleaseRecipe, error) {
	return e.recipeParser.Parse(e.tool.Recipe)
}

func NewGitReleaseProvider(tool models.Tool, logger loggers.Logger) Provider {
	return &GitReleaseProvider{
		tool:         tool,
		recipeParser: recipes.NewGitReleaseRecipeParser(),
		handler:      handlers.NewGitReleaseHandler(logger),
	}
}

func (e *GitReleaseProvider) Install() error {
	recipe, err := e.getRecipe()
	if err != nil {
		return err
	}

	return fmt.Errorf(recipe.Package)
}

func (e *GitReleaseProvider) Update() error {
	return nil
}

func (e *GitReleaseProvider) Remove() error {
	return nil
}

func (e *GitReleaseProvider) LatestVersion() (string, error) {
	recipe, err := e.getRecipe()
	if err != nil {
		return "", err
	}

	release, err := e.handler.LatestVersion(recipe.Repository)
	return release.TagName, err
}

func (e *GitReleaseProvider) InstaledVersion() (string, error) {
	return "", nil
}
