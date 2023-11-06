package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/Mth-Ryan/lspcli/internal/utils"
	"github.com/Mth-Ryan/lspcli/pkg/loggers"
)

type GithubRealease struct {
	ID        int    `json:"id"`
	Url       string `json:"url"`
	AssetsUrl string `json:"assets_url"`
	TagName   string `json:"tag_name"`
}

type GithubReleaseAsset struct {
	ID                 int    `json:"id"`
	Url                string `json:"url"`
	Name               string `json:"name"`
	ContentType        string `json:"content_type"`
	Size               int    `json:"size"`
	BrowserDownloadUrl string `json:"browser_download_url"`
}

type GithubReleaseWithAsset struct {
	GithubRealease
	Asset GithubReleaseAsset
}

type GithubReleaseWithAssets struct {
	GithubRealease
	Assets []GithubReleaseAsset
}

type GithubReleaseHandler struct {
	baseUrl    string
	apiVersion string
	logger     loggers.Logger
}

func NewGitReleaseHandler(logger loggers.Logger) *GithubReleaseHandler {
	return &GithubReleaseHandler{
		baseUrl:    "https://api.github.com",
		apiVersion: "2022-11-28",
		logger:     logger,
	}
}

func (g *GithubReleaseHandler) get(p string) ([]byte, error) {
	requestPath, _ := url.JoinPath(g.baseUrl, p)

	request, err := http.NewRequest(http.MethodGet, requestPath, nil)
	if err != nil {
		return []byte{}, err
	}

	request.Header.Add("X-GitHub-Api-Version", g.apiVersion)

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	return io.ReadAll(res.Body)
}

func (g *GithubReleaseHandler) getJson(p string, out any) error {
	body, err := g.get(p)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, out)
}

func (g *GithubReleaseHandler) LatestVersion(repo string) (GithubRealease, error) {
	g.logger.Log("Trying to get the latest release tag for: https://github.com/%s", repo)

	release := &GithubRealease{}
	err := g.getJson(fmt.Sprintf("/repos/%s/releases/latest", repo), release)

	if err == nil {
		g.logger.Log("Release tag found: %s", release.TagName)
	}

	return *release, err
}

func (g *GithubReleaseHandler) LatestVersionWithAssets(repo string) (GithubReleaseWithAssets, error) {
	assets := &[]GithubReleaseAsset{}

	rawRelease, err := g.LatestVersion(repo)
	release := GithubReleaseWithAssets{
		GithubRealease: rawRelease,
	}
	if err != nil {
		return release, err
	}

	g.logger.Log("Finding the %s assets", release.TagName)

	err = g.getJson(fmt.Sprintf("/repos/%s/releases/%d/assets", repo, rawRelease.ID), assets)
	release.Assets = *assets

	if err == nil {
		g.logger.Log("Assets found")
	}

	return release, err
}

func downloadFile(filepath string, url string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	_, err = io.Copy(out, res.Body)
	return err
}

func (g *GithubReleaseHandler) DownloadAssetFromLatestVersion(repo string, assetName string, filepath string) (GithubReleaseWithAsset, error) {
	release, err := g.LatestVersionWithAssets(repo)
	if err != nil {
		return GithubReleaseWithAsset{}, err
	}

	asset, err := utils.First(release.Assets, func(asset GithubReleaseAsset) bool {
		return asset.Name == assetName
	})
	withAsset := GithubReleaseWithAsset{
		GithubRealease: release.GithubRealease,
		Asset:          asset,
	}
	if err != nil {
		return withAsset, err
	} else {
		g.logger.Log("Target asset found: %s", asset.Name)
	}

	g.logger.Log("Downloading asset from: %s...", asset.BrowserDownloadUrl)
	err = downloadFile(filepath, asset.BrowserDownloadUrl)

	if err == nil {
		g.logger.Log("Download completed successfully")
	}

	return withAsset, err
}
