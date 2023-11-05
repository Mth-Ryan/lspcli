package providers

import (
	"fmt"

	"github.com/Mth-Ryan/lspcli/pkg/models"
	"github.com/Mth-Ryan/lspcli/pkg/recipes"
)

type GitReleaseProvider struct {
	tool          models.Tool
	recipesParser *recipes.GitReleaseRecipeParser
}

func NewGitReleaseProvider(tool models.Tool) Provider {
	return &GitReleaseProvider{
		tool:          tool,
		recipesParser: recipes.NewGitReleaseRecipeParser(),
	}
}

func (e *GitReleaseProvider) Install() error {
	recipe, err := e.recipesParser.Parse(e.tool.Recipe)
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
