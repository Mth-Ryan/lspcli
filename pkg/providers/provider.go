package providers

import (
	"fmt"

	"github.com/Mth-Ryan/lspcli/pkg/models"
)

type Provider interface {
	Install() error
	Update() error
	Remove() error
	LatestVersion() (string, error)
	InstaledVersion() (string, error)
}

type ProviderConstructor = func(models.Tool) Provider

var providersConstructors = map[string]ProviderConstructor{
	models.RECIPE_GIT_RELEASE: NewGitReleaseProvider,
	models.RECIPE_GO:          NewGoProvider,
	models.RECIPE_NPM:         NewNpmProvider,
}

func GetProvider(tool models.Tool) (Provider, error) {
	kind, ok := tool.Recipe["kind"]
	if !ok {
		return &ErrProvider{}, fmt.Errorf("unable to get the tool's recipe kind")
	}

	if constructors, ok := providersConstructors[fmt.Sprint(kind)]; ok {
		return constructors(tool), nil
	}
	return &ErrProvider{}, fmt.Errorf("Invalid recipe kind: %s", kind)
}
