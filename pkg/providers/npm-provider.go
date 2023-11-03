package providers

import "github.com/Mth-Ryan/lspcli/pkg/models"

type NpmProvider struct{}

func (e *NpmProvider) Install(tool models.Tool) error {
	return nil
}

func (e *NpmProvider) Update(tool models.Tool) error {
	return nil
}

func (e *NpmProvider) Remove(tool models.Tool) error {
	return nil
}

func (e *NpmProvider) LatestVersion(tool models.Tool) error {
	return nil
}

func (e *NpmProvider) InstaledVersion(tool models.Tool) error {
	return nil
}
