package providers

import "github.com/Mth-Ryan/lspcli/pkg/models"

type GitReleaseProvider struct{}

func (e *GitReleaseProvider) Install(tool models.Tool) error {
	return nil
}

func (e *GitReleaseProvider) Update(tool models.Tool) error {
	return nil
}

func (e *GitReleaseProvider) Remove(tool models.Tool) error {
	return nil
}

func (e *GitReleaseProvider) LatestVersion(tool models.Tool) error {
	return nil
}

func (e *GitReleaseProvider) InstaledVersion(tool models.Tool) error {
	return nil
}
