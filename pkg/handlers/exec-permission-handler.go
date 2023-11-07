package handlers

import "os/exec"

type ExecPermissionHandler interface {
	SetPermission(filepath string) error
}

type UnixExecPermissionHandler struct{}

func NewUnixExecPermissionHandler() *UnixExecPermissionHandler {
	return &UnixExecPermissionHandler{}
}

func (s *UnixExecPermissionHandler) SetPermission(filepath string) error {
	cmd := exec.Command("chmod", "+x", filepath)
	return cmd.Run()
}
