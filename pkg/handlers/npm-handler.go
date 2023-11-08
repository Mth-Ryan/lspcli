package handlers

import (
	"os/exec"
	"strings"

	"github.com/Mth-Ryan/lspcli/pkg/loggers"
)

type NpmHandler struct {
	logger loggers.Logger
}

func NewNpmHandler(logger loggers.Logger) *NpmHandler {
	return &NpmHandler{
		logger,
	}
}

func (h *NpmHandler) GetLatestVersion(pkg string) (string, error) {
	cmd := exec.Command("npm", "view", pkg, "version")

	h.logger.Log("Get the latest version from npm")
	raw, err := cmd.Output()
	return strings.TrimSpace(string(raw)), err
}

func (h *NpmHandler) Install(pkg string) error {
	cmd := exec.Command("npm", "install", "-g", pkg)
	h.logger.Log("Installing package with npm...")
	return cmd.Run()
}

func (h *NpmHandler) Update(pkg string) error {
	cmd := exec.Command("npm", "update", "-g", pkg)
	h.logger.Log("Updating package with npm...")
	return cmd.Run()
}

func (h *NpmHandler) Remove(pkg string) error {
	cmd := exec.Command("npm", "uninstall", "-g", pkg)
	h.logger.Log("Removing package with npm...")
	return cmd.Run()
}
