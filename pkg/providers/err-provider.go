package providers

import "github.com/Mth-Ryan/lspcli/pkg/models"

type ErrProvider struct{}

func (e *ErrProvider) Install(tool models.Tool) error {
	return nil
}

func (e *ErrProvider) Update(tool models.Tool) error {
	return nil
}

func (e *ErrProvider) Remove(tool models.Tool) error {
	return nil
}

func (e *ErrProvider) LatestVersion(tool models.Tool) error {
	return nil
}

func (e *ErrProvider) InstaledVersion(tool models.Tool) error {
	return nil
}
