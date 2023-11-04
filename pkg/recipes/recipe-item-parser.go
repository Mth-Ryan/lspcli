package recipes

import (
	"bytes"
	"runtime"
	"strings"
	"text/template"
)

type ItemContext struct {
	OS   string
	Arch string
}

type ItemsParser struct {
	context ItemContext
}

func NewItemsParser() *ItemsParser {
	context := ItemContext{
		OS:   runtime.GOOS,
		Arch: runtime.GOARCH,
	}

	return &ItemsParser{
		context,
	}
}

func (p *ItemsParser) Parse(item string) (string, error) {
	templ, err := template.New("item").
		Funcs(parserTmplFuncs).
		Parse(item)

	if err != nil {
		return item, err
	}

	buf := new(bytes.Buffer)
	if err := templ.Execute(buf, p.context); err != nil {
		return item, err
	}

	return strings.TrimSpace(buf.String()), nil
}

var parserTmplFuncs = map[string]any{
	"alterOS": func(o string) string {
		switch o {
		case "windows":
			return "win"
		case "darwin":
			return "osx"
		default:
			return o
		}
	},

	"alterArch": func(o string) string {
		switch o {
		case "386":
			return "x86"
		case "amd64":
			return "x64"
		case "arm":
			return "aarch"
		case "arm64":
			return "aarch64"
		default:
			return o
		}
	},
}
