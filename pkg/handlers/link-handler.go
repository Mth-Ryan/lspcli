package handlers

import (
	"os/exec"
	"path/filepath"
)

type LinkHander interface {
	CreateLink(source string, destination string) error
}

type SymbolicLinkHandler struct{}

func NewSymbolicLinkHandler() *SymbolicLinkHandler {
	return &SymbolicLinkHandler{}
}

func (s *SymbolicLinkHandler) CreateLink(source string, destination string) error {
	absSource, _ := filepath.Abs(source)
	absDestination, _ := filepath.Abs(destination)

	cmd := exec.Command("ln", "-sf", absSource, absDestination)
	return cmd.Run()
}
