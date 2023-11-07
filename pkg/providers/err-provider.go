package providers

type ErrProvider struct{}

func (e *ErrProvider) Install() error {
	return nil
}

func (e *ErrProvider) Update() error {
	return nil
}

func (e *ErrProvider) Remove() error {
	return nil
}

func (e *ErrProvider) LatestVersion() (string, error) {
	return "", nil
}

func (e *ErrProvider) InstalledVersion() (string, error) {
	return "", nil
}

func (e *ErrProvider) InstalledVersions() (map[string](*string), error) {
	return map[string](*string){}, nil
}
