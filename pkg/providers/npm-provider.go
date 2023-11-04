package providers

import "github.com/Mth-Ryan/lspcli/pkg/models"

type NpmProvider struct {
	tool models.Tool
}

func NewNpmProvider(tool models.Tool) Provider {
	return &NpmProvider{
		tool,
	}
}

func (e *NpmProvider) Install() error {
	return nil
}

func (e *NpmProvider) Update() error {
	return nil
}

func (e *NpmProvider) Remove() error {
	return nil
}

func (e *NpmProvider) LatestVersion() error {
	return nil
}

func (e *NpmProvider) InstaledVersion() error {
	return nil
}
