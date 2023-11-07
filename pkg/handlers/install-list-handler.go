package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/Mth-Ryan/lspcli/pkg/runtime"
)

type InstallsListHandler interface {
	GetInstalls() (map[string](*string), error)
	SetVersion(id string, version *string) error
}

type JsonInstallsHandler struct {
	runtimeConf runtime.Conf
}

func NewJsonInstallsHandler(runtimeConf runtime.Conf) *JsonInstallsHandler {
	return &JsonInstallsHandler{
		runtimeConf,
	}
}

func (j *JsonInstallsHandler) initialize() (*os.File, error) {
	listPath := j.runtimeConf.InstallsListPath()
	if _, err := os.Stat(listPath); err == nil {
		return os.Open(listPath)
	} else if errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(listPath)
		if err != nil {
			return file, err
		}

		_, err = file.Write([]byte("{}"))
		if err != nil {
			return nil, err
		}
		file.Close()

		return os.Open(listPath)
	} else {
		return nil, err
	}
}

func (j *JsonInstallsHandler) GetInstalls() (map[string](*string), error) {
	installs := map[string](*string){}

	file, err := j.initialize()
	if err != nil {
		return installs, err
	}
	defer file.Close()

	raw, err := io.ReadAll(file)
	if err != nil {
		return installs, err
	}

	err = json.Unmarshal(raw, &installs)
	return installs, err
}

func (j *JsonInstallsHandler) SetVersion(id string, version *string) error {
	installs, err := j.GetInstalls()
	if err != nil {
		return err
	}
	installs[id] = version

	raw, err := json.Marshal(installs)
	if err != nil {
		return err
	}

	return os.WriteFile(j.runtimeConf.InstallsListPath(), raw, 0644)
}
