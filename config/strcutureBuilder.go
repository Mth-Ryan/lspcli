package config

import (
	"os"
	"sync"
)

type structureChecker struct {
    hasStruct bool
	mu sync.Mutex
}

var structureCheckerSingleton = &structureChecker{
	hasStruct: false,
}

func createFolderIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, os.ModePerm)
	}
	return nil
}

func createFoldersStructure() error {
	if err := createFolderIfNotExists(config.RuntimeFolderPath); err != nil {
		return err
	}

	if err := createFolderIfNotExists(config.InstallsFolderPath); err != nil {
		return err
	}

	if err := createFolderIfNotExists(config.BinFolderPath); err != nil {
		return err
	}

	return nil
}

func (s *structureChecker)  CreateFoldersStructure() error {
	if !s.hasStruct {
		s.mu.Lock()
		defer s.mu.Unlock()

		if err := createFoldersStructure(); err != nil {
			return err
		}

		s.hasStruct = true
	}
	return nil
}
