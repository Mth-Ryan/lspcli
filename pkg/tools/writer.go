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
	WriteAll([]models.Tool)
	Write(models.Tool)
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

func (w *TableWriter) createListTable() table.Table {
	tbl := table.New("id", "name", "kind", "languages", "installed version")
	tbl.WithPadding(w.padding)
	tbl.WithHeaderFormatter(w.headerFmt)
	tbl.WithFirstColumnFormatter(w.idFmt)

	return tbl
}

func (w *TableWriter) addListRow(tbl table.Table, tool models.Tool) {
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

func (w *TableWriter) WriteAll(tools []models.Tool) {
	tbl := w.createListTable()

	for _, tool := range tools {
		w.addListRow(tbl, tool)
	}

	fmt.Println()
	tbl.Print()
	fmt.Println()
}

func (w *TableWriter) Write(tool models.Tool) {
	tbl := table.New("", "")
	tbl.WithFirstColumnFormatter(w.headerFmt)

	tbl.AddRow("id", w.idFmt("%s", tool.ID))
	tbl.AddRow("kind", tool.Kind)
	tbl.AddRow("name", tool.Name)
	tbl.AddRow("languages", strings.Join(tool.Languages, ", "))
	tbl.AddRow("dependencies", strings.Join(tool.Dependencies, ", "))
	if tool.LatestVersion != nil {
		tbl.AddRow("latest version", tool.LatestVersion)
	}
	if tool.InstalledVersion != nil {
		tbl.AddRow("installed version", tool.InstalledVersion)
	}
	tbl.AddRow("Url", tool.Url)
	tbl.AddRow("Description", tool.Description)

	tbl.Print()
	fmt.Println()
}

type JsonWriter struct{}

func NewJsonWriter() *JsonWriter {
	return &JsonWriter{}
}

func (w *JsonWriter) WriteAll(tools []models.Tool) {
	raw, err := json.Marshal(utils.Map(tools, models.ToShort))
	if err != nil {
		panic("Unable to unmarshal tools slice to json. This should be unreachable")
	}
	json := string(raw[:])
	fmt.Println(json)
}

func (w *JsonWriter) Write(tool models.Tool) {
	raw, err := json.Marshal(models.ToDescribe(tool))
	if err != nil {
		panic("Unable to unmarshal tool struct to json. This should be unreachable")
	}
	json := string(raw[:])
	fmt.Println(json)
}
