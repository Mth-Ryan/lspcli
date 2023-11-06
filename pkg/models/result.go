package models

type ResultKind = string

const (
	RESULT_ERR   = "err"
	RESULT_OK    = "ok"
	RESULT_NO_OP = "no_op"
)

type Result struct {
	Kind    ResultKind `json:"kind" yaml:"kind"`
	Message string     `json:"message" yaml:"message"`
}
