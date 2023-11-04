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
