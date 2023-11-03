package models

type ResultKind = string

const (
	RESULT_ERR = "err"
	RESULT_OK  = "ok"
)

type Result struct {
	Kind    ResultKind `json:"kind" yaml:"kind"`
	Message string     `json:"message" yaml:"message"`
}
