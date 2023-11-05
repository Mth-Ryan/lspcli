package recipes

import (
	"bytes"
	"runtime"
	"strings"
	"text/template"

	"github.com/Mth-Ryan/lspcli/pkg/models"
)

type ItemsParser interface {
	Parse(item string) (string, error)
}

type TemplateContext = map[string]string

type TemplateParser struct {
	context map[string]string
}

func replaceContext(replaces models.RecipeContextReplaces, ctx TemplateContext) TemplateContext {
	if replaces == nil {
		return ctx
	}

	newContext := TemplateContext{}

	for key, value := range ctx {
		if replaceMap, ok := (*replaces)[key]; ok {
			if newValue, ok := replaceMap[value]; ok {
				newContext[key] = newValue
			} else {
				newContext[key] = value
			}
		} else {
			newContext[key] = value
		}
	}

	return newContext
}

func NewTemplateParser(replaces models.RecipeContextReplaces) *TemplateParser {
	context := replaceContext(replaces, TemplateContext{
		"OS":   runtime.GOOS,
		"Arch": runtime.GOARCH,
	})

	return &TemplateParser{
		context,
	}
}

func (p *TemplateParser) Parse(item string) (string, error) {
	templ, err := template.New("item").Parse(item)

	if err != nil {
		return item, err
	}

	buf := new(bytes.Buffer)
	if err := templ.Execute(buf, p.context); err != nil {
		return item, err
	}

	return strings.TrimSpace(buf.String()), nil
}

func getReplacesFromRecipe(raw map[string]any) models.RecipeContextReplaces {
	var replaces models.RecipeContextReplaces
	if repl, ok := raw["context_replaces"]; ok {
		switch r := repl.(type) {
		case models.RecipeContextReplaces:
			replaces = r
		}
	}
	return replaces
}
