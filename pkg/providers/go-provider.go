package providers

import "github.com/Mth-Ryan/lspcli/pkg/models"

type GoProvider struct{}

func (e *GoProvider) Install(tool models.Tool) error {
	return nil
}

func (e *GoProvider) Update(tool models.Tool) error {
	return nil
}

func (e *GoProvider) Remove(tool models.Tool) error {
	return nil
}

func (e *GoProvider) LatestVersion(tool models.Tool) error {
	return nil
}

func (e *GoProvider) InstaledVersion(tool models.Tool) error {
	return nil
}
