package tools

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/Mth-Ryan/lspcli/pkg/models"
	"gopkg.in/yaml.v3"
)

type Reader interface {
	GetAll() ([]models.Tool, error)
	Get(id string) (models.Tool, error)
}

type RuntimeReader struct {
	toolsPath string
	cache     *[]models.Tool // cli usualy run and die, but who knows...
}

func NewRuntimeReader(runtimePath string) *RuntimeReader {
	return &RuntimeReader{
		toolsPath: path.Join(runtimePath, "tools"),
		cache:     nil,
	}
}

func getAllYamlPathsInTheFolder(folderPath string) ([]string, error) {
	entries, err := os.ReadDir(folderPath)
	if err != nil {
		return []string{}, err
	}

	files := []string{}

	for _, entity := range entries {
		if !entity.IsDir() {
			filename := entity.Name()
			ext := filepath.Ext(filename)
			if ext == ".yml" || ext == ".yaml" {
				files = append(files, path.Join(folderPath, filename))
			}
		}
	}

	return files, nil
}

func parseTool(toolRaw []byte) (models.Tool, error) {
	tool := new(models.Tool)
	err := yaml.Unmarshal(toolRaw, tool)
	return *tool, err
}

func (r *RuntimeReader) GetAll() ([]models.Tool, error) {
	if r.cache != nil {
		return *r.cache, nil
	}

	paths, err := getAllYamlPathsInTheFolder(r.toolsPath)
	if err != nil {
		return []models.Tool{}, err
	}

	tools := []models.Tool{}
	for _, filePath := range paths {
		toolRaw, err := os.ReadFile(filePath)
		if err != nil {
			continue
		}

		tool, err := parseTool(toolRaw)
		if err == nil {
			tools = append(tools, tool)
		}
	}

	return tools, nil
}

func (r *RuntimeReader) Get(id string) (models.Tool, error) {
	tools, err := r.GetAll()
	if err != nil {
		return models.Tool{}, err
	}

	for _, tool := range tools {
		if tool.ID == id {
			return tool, nil
		}
	}

	return models.Tool{}, fmt.Errorf("Tool with id: %s not found", id)
}
