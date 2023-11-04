package providers

import "github.com/Mth-Ryan/lspcli/pkg/models"

type GoProvider struct {
	tool models.Tool
}

func NewGoProvider(tool models.Tool) Provider {
	return &GoProvider{
		tool,
	}
}

func (e *GoProvider) Install() error {
	return nil
}

func (e *GoProvider) Update() error {
	return nil
}

func (e *GoProvider) Remove() error {
	return nil
}

func (e *GoProvider) LatestVersion() error {
	return nil
}

func (e *GoProvider) InstaledVersion() error {
	return nil
}
