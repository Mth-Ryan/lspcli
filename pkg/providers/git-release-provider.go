package providers

import (
	"fmt"

	"github.com/Mth-Ryan/lspcli/pkg/models"
	"github.com/Mth-Ryan/lspcli/pkg/recipes"
)

type GitReleaseProvider struct {
	tool         models.Tool
	recipeParser *recipes.GitReleaseRecipeParser
}

func (e *GitReleaseProvider) getRecipe() (*models.GitReleaseRecipe, error) {
	return e.recipeParser.Parse(e.tool.Recipe)
}

func NewGitReleaseProvider(tool models.Tool) Provider {
	return &GitReleaseProvider{
		tool:         tool,
		recipeParser: recipes.NewGitReleaseRecipeParser(),
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

func (e *GitReleaseProvider) LatestVersion() error {
	return nil
}

func (e *GitReleaseProvider) InstaledVersion() error {
	return nil
}
