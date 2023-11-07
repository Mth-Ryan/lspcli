package handlers

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type ArchiveHandlerFactory interface {
	GetHandler(filename string) (ArchiveHandler, error)
}

type NativeArchiveHandlerFactory struct{}

func mustMatch(s, pattern string) bool {
	return regexp.MustCompile(pattern).MatchString(s)
}

func (f *NativeArchiveHandlerFactory) GetHandler(filename string) (ArchiveHandler, error) {
	switch {
	case mustMatch(filename, `.zip$`):
		return NewZipHandler(), nil
	case mustMatch(filename, `(.tar$)|(.tar.((gz)|(bz2)|(br)|(xz))$)`):
		return NewTarHandler(filename), nil
	}

	return &ErrHandler{}, fmt.Errorf("Unable to get a handler for this extension")
}

func NewNativeArchiveHandlerFactory() *NativeArchiveHandlerFactory {
	return &NativeArchiveHandlerFactory{}
}

type ArchiveHandler interface {
	Extract(archivePath string, outputPath string) error
}

// Zip handler
type ZipHandler struct{}

func NewZipHandler() *ZipHandler {
	return &ZipHandler{}
}

func (h *ZipHandler) Extract(archivePath string, outputPath string) error {
	cmd := exec.Command("unzip", archivePath, "-d", outputPath)
	return cmd.Run()
}

// Tar handler
type TarHandler struct {
	filename string
}

func NewTarHandler(filename string) *TarHandler {
	return &TarHandler{
		filename,
	}
}

func (h *TarHandler) Extract(archivePath string, outputPath string) error {
	err := os.MkdirAll(outputPath, 0755)
	if err != nil {
		return err
	}

	args := []string{"-xf"}
	switch {
	case strings.HasSuffix(h.filename, ".gz"):
		args = []string{"xzf"}
	case strings.HasSuffix(h.filename, ".bz2"):
		args = []string{"xjf"}
	case strings.HasSuffix(h.filename, ".br"):
		args = []string{"-I", "lbzip2", "-xf"}
	case strings.HasSuffix(h.filename, ".xz"):
		args = []string{"-Jxf"}
	case strings.HasSuffix(h.filename, ".xz"):
		args = []string{"-I", "zstd", "-xf"}
	}

	args = append(args, archivePath, "-C", outputPath)

	cmd := exec.Command("tar", args...)
	return cmd.Run()
}

type ErrHandler struct{}

func NewErrHandler() *ErrHandler {
	return &ErrHandler{}
}

func (h *ErrHandler) Extract(archivePath string, outputPath string) error {
	return nil
}
