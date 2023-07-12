package config

import (
	"log"
	"os"
	"path/filepath"
)

const (
	EXEC_DEBUG   = "DEBUG"
	EXEC_RELEASE = "RELEASE"
)

// This section is intentionally Hardcoded
// HACK: change this to release mode on publish
const _EXEC_MODE = EXEC_DEBUG
const _VERSION   = "0.1.0"

type Config struct {
	Version            string
    ExecMode           string
	RuntimeFolderPath  string
	RecipesFilePath    string
	InstalledsListPath string
	InstallsFolderPath string
	BinFolderPath      string
}


func getRuntimeFolder() string {
	if _EXEC_MODE == EXEC_DEBUG {
		return "runtime"
	}

	confDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("Unable to finde the user config dir:\n%v", err)
	}

	return filepath.Join(confDir, "lspcli")
}

// Singleton config variables
var (
	runtimeFolder = getRuntimeFolder()
	config = Config{
		Version: _VERSION,
		ExecMode: _EXEC_MODE,
		RuntimeFolderPath: runtimeFolder,
		RecipesFilePath: filepath.Join(runtimeFolder, "recipes.yml"),
		InstalledsListPath: filepath.Join(runtimeFolder, "installs.yml"),
		InstallsFolderPath: filepath.Join(runtimeFolder, "installs"),
		BinFolderPath: filepath.Join(runtimeFolder, "bin"),
	}
)

func Get() Config {
	structureCheckerSingleton.CreateFoldersStructure()
	return config
}
