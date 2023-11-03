package result

import (
	"encoding/json"
	"fmt"

	"github.com/Mth-Ryan/lspcli/pkg/models"
	"github.com/fatih/color"
)

type Writer interface {
	Write(models.Result)
}

type PlainWriter struct{}

func NewPlainWriter() *PlainWriter {
	return &PlainWriter{}
}

func (w *PlainWriter) Write(result models.Result) {
	printFmt := color.New(color.FgRed).PrintfFunc()
	printFmt("Error: %s.\n", result.Message)
}

type JsonWriter struct{}

func NewJsonWriter() *JsonWriter {
	return &JsonWriter{}
}

func (w *JsonWriter) Write(result models.Result) {
	raw, err := json.Marshal(result)
	if err != nil {
		panic("Unable to marshal the result struct to json. This should be unreachable")
	}
	json := string(raw[:])
	fmt.Println(json)
}
