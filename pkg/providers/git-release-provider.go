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
	tool                  models.Tool
	runtimeConf           runtime.Conf
	logger                loggers.Logger
	recipeParser          *recipes.GitReleaseRecipeParser
	handler               *handlers.GithubReleaseHandler
	archiveFactory        handlers.ArchiveHandlerFactory
	linkHandler           handlers.LinkHander
	execPermissionHandler handlers.ExecPermissionHandler
}

func (e *GitReleaseProvider) getRecipe() (*models.GitReleaseRecipe, error) {
	return e.recipeParser.Parse(e.tool.Recipe)
}

func NewGitReleaseProvider(runConf runtime.Conf, tool models.Tool, logger loggers.Logger) Provider {

	return &GitReleaseProvider{
		tool:                  tool,
		runtimeConf:           runConf,
		logger:                logger,
		recipeParser:          recipes.NewGitReleaseRecipeParser(),
		handler:               handlers.NewGitReleaseHandler(logger),
		archiveFactory:        handlers.NewNativeArchiveHandlerFactory(),
		linkHandler:           handlers.NewSymbolicLinkHandler(),
		execPermissionHandler: handlers.NewUnixExecPermissionHandler(),
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

	e.logger.Log("Extracting the release asset")
	installPath := path.Join(e.runtimeConf.InstallsPath(), e.tool.ID)
	err = archiveHandler.Extract(cachedAsset, installPath)
	if err != nil {
		return err
	}

	e.logger.Log("Setting the binary exec permission")
	originalBinaryPath := path.Join(installPath, recipe.BinaryInnerPath)
	err = e.execPermissionHandler.SetPermission(originalBinaryPath)
	if err != nil {
		return err
	}

	e.logger.Log("Creating the main binary symbolic link")
	linkBinaryPath := path.Join(e.runtimeConf.BinPath(), recipe.BinaryName)
	err = e.linkHandler.CreateLink(originalBinaryPath, linkBinaryPath)
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
