package commands

import (
	"github.com/Mth-Ryan/lspcli/pkg/models"
	"github.com/Mth-Ryan/lspcli/pkg/tools"
)

type ListCommand struct {
	reader tools.Reader
}

func NewListCommand(reader tools.Reader) *ListCommand {
	return &ListCommand{
		reader,
	}
}

func filter[T any](slice []T, f func(T) bool) []T {
	newValues := []T{}
	for _, value := range slice {
		if f(value) {
			newValues = append(newValues, value)
		}
	}

	return newValues
}

func fmap[T any](slice []T, f func(T) T) []T {
	newValues := []T{}
	for _, value := range slice {
		newValues = append(newValues, f(value))
	}

	return newValues
}

func (l *ListCommand) GetAll(where *(func(models.Tool) bool)) ([]models.Tool, error) {
	tools, err := l.reader.GetAll()
	if err != nil {
		return tools, err
	}

	if where != nil {
		return filter(tools, *where), nil
	}

	return tools, nil
}
