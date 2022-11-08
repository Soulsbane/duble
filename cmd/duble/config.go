package main

import gap "github.com/muesli/go-app-paths"

const (
	vendor          string = "Raijinsoft"
	applicationName string = "duble"
)

type Config struct {
}

func getConfigFilePath() (string, error) {
	scope := gap.NewVendorScope(gap.User, vendor, applicationName)
	path, err := scope.ConfigPath("config.toml")

	return path, err
}
