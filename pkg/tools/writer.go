package tools

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Mth-Ryan/lspcli/internal/utils"
	"github.com/Mth-Ryan/lspcli/pkg/models"
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type Writer interface {
	Write([]models.Tool)
}

type TableWriter struct {
	padding   int
	headerFmt func(string, ...any) string
	idFmt     func(string, ...any) string
}

func NewTableWriter() *TableWriter {
	return &TableWriter{
		padding:   2,
		headerFmt: color.New(color.FgBlue, color.Bold).Sprintf,
		idFmt:     color.New(color.FgCyan).Sprintf,
	}
}

func (w *TableWriter) createTable() table.Table {
	tbl := table.New("id", "name", "kind", "languages", "installed version")
	tbl.WithPadding(w.padding)
	tbl.WithHeaderFormatter(w.headerFmt)
	tbl.WithFirstColumnFormatter(w.idFmt)

	return tbl
}

func (w *TableWriter) addRow(tbl table.Table, tool models.Tool) {
	languages := strings.Join(tool.Languages, ", ")
	version := ""
	if tool.InstalledVersion != nil {
		version = *tool.InstalledVersion
	}

	tbl.AddRow(
		tool.ID,
		tool.Name,
		tool.Kind,
		languages,
		version,
	)
}

func (w *TableWriter) Write(tools []models.Tool) {
	tbl := w.createTable()

	for _, tool := range tools {
		w.addRow(tbl, tool)
	}

	fmt.Println()
	tbl.Print()
	fmt.Println()
}

type JsonWriter struct{}

func NewJsonWriter() *JsonWriter {
	return &JsonWriter{}
}

func (w *JsonWriter) Write(tools []models.Tool) {
	raw, err := json.Marshal(utils.Map(tools, models.ToShort))
	if err != nil {
		panic("Unable to unmarshal tools slice to json. This should be unreachable")
	}
	json := string(raw[:])
	fmt.Println(json)
}
