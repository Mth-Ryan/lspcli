package providers

import (
	"fmt"

	"github.com/Mth-Ryan/lspcli/pkg/models"
)

type Provider interface {
	Install(tool models.Tool) error
	Update(tool models.Tool) error
	Remove(tool models.Tool) error
	LatestVersion(tool models.Tool) error
	InstaledVersion(tool models.Tool) error
}

var providers = map[string]Provider{
	models.RECIPE_GIT_RELEASE: new(GitReleaseProvider),
	models.RECIPE_GO:          new(GoProvider),
	models.RECIPE_NPM:         new(NpmProvider),
}

func GetProvider(tool models.Tool) (Provider, error) {
	if provider, ok := providers[tool.Recipe.Kind]; ok {
		return provider, nil
	}
	return &ErrProvider{}, fmt.Errorf("Invalid recipe kind: %s", tool.Recipe.Kind)
}
