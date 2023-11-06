package providers

import (
	"os"
	"path"

	"github.com/Mth-Ryan/lspcli/pkg/handlers"
	"github.com/Mth-Ryan/lspcli/pkg/loggers"
	"github.com/Mth-Ryan/lspcli/pkg/models"
	"github.com/Mth-Ryan/lspcli/pkg/recipes"
	"github.com/Mth-Ryan/lspcli/pkg/runtime"
)

type GitReleaseProvider struct {
	tool           models.Tool
	runtimeConf    runtime.Conf
	recipeParser   *recipes.GitReleaseRecipeParser
	handler        *handlers.GithubReleaseHandler
	logger         loggers.Logger
	archiveFactory handlers.ArchiveHandlerFactory
}

func (e *GitReleaseProvider) getRecipe() (*models.GitReleaseRecipe, error) {
	return e.recipeParser.Parse(e.tool.Recipe)
}

func NewGitReleaseProvider(runConf runtime.Conf, tool models.Tool, logger loggers.Logger) Provider {

	return &GitReleaseProvider{
		tool:           tool,
		runtimeConf:    runConf,
		recipeParser:   recipes.NewGitReleaseRecipeParser(),
		handler:        handlers.NewGitReleaseHandler(logger),
		logger:         logger,
		archiveFactory: handlers.NewNativeArchiveHandlerFactory(),
	}
}

func (e *GitReleaseProvider) Install() error {
	recipe, err := e.getRecipe()
	if err != nil {
		return err
	}

	e.logger.Log("Removing the cached release file if exists")
	cachedAsset := path.Join(e.runtimeConf.CachePath(), recipe.Package)

	os.Remove(cachedAsset)

	_, err = e.handler.DownloadAssetFromLatestVersion(
		recipe.Repository,
		recipe.Package,
		cachedAsset,
	)
	if err != nil {
		return err
	}

	archiveHandler, err := e.archiveFactory.GetHandler(recipe.Package)
	if err != nil {
		return err
	}

	installPath := path.Join(e.runtimeConf.InstallsPath(), e.tool.ID)
	err = archiveHandler.Extract(cachedAsset, installPath)
	if err != nil {
		return err
	}

	return nil
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
