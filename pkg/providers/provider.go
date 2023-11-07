package providers

import (
	"fmt"

	"github.com/Mth-Ryan/lspcli/pkg/loggers"
	"github.com/Mth-Ryan/lspcli/pkg/models"
	"github.com/Mth-Ryan/lspcli/pkg/runtime"
)

type Provider interface {
	Install() error
	Update() error
	Remove() error
	LatestVersion() (string, error)
	InstalledVersion() (string, error)
}

type ProviderConstructor = func(runtime.Conf, models.Tool, loggers.Logger) Provider

var providersConstructors = map[string]ProviderConstructor{
	models.RECIPE_GIT_RELEASE: NewGitReleaseProvider,
	models.RECIPE_GO:          NewGoProvider,
	models.RECIPE_NPM:         NewNpmProvider,
}

func GetProvider(runConf runtime.Conf, tool models.Tool, logger loggers.Logger) (Provider, error) {
	kind, ok := tool.Recipe["kind"]
	if !ok {
		return &ErrProvider{}, fmt.Errorf("unable to get the tool's recipe kind")
	}

	if constructors, ok := providersConstructors[fmt.Sprint(kind)]; ok {
		return constructors(runConf, tool, logger), nil
	}
	return &ErrProvider{}, fmt.Errorf("Invalid recipe kind: %s", kind)
}
