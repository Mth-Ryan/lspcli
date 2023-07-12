package core

import (
	"os"

	yaml "gopkg.in/yaml.v3"
)

type Recipe struct {
	Lsp string          `yaml:"lsp"`
	Lang string         `yaml:"lang"`
	Provider string     `yaml:"provider"`
	Ref string          `yaml:"ref"`
	Steps []string `yaml:"steps"`
}

func loadRecipesFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func parseRecipes(ymlData []byte) ([]Recipe, error) {
	recipes := []Recipe{}
	err := yaml.Unmarshal(ymlData, &recipes)
	return recipes, err
}

func LoadAndParseRecipes(path string) ([]Recipe, error) {
	data, err := loadRecipesFile(path)
	if err != nil {
		return []Recipe{}, err
	}
	return parseRecipes(data)
}
