package providers

import (
	"fmt"

	"github.com/Mth-Ryan/lspcli/pkg/models"
	"github.com/Mth-Ryan/lspcli/pkg/recipes"
)

type GitReleaseProvider struct {
	tool             models.Tool
	recipeItemParser recipes.ItemsParser
}

func NewGitReleaseProvider(tool models.Tool) Provider {
	return &GitReleaseProvider{
		tool:             tool,
		recipeItemParser: *recipes.NewItemsParser(),
	}
}

func (e *GitReleaseProvider) Install() error {
	pack, err := e.recipeItemParser.Parse(e.tool.Recipe.GitReleaseRecipe.Package)
	if err != nil {
		return err
	}

	return fmt.Errorf(pack)
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
