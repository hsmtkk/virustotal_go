package apikey

import (
	"fmt"
	"os"
)

const VirusTotalAPIKey = "VIRUS_TOTAL_API_KEY"

type APIKeyLoader interface {
	LoadAPIKey() (string, error)
}

func New() APIKeyLoader {
	return &apiKeyLoaderImpl{}
}

type apiKeyLoaderImpl struct{}

func (imp *apiKeyLoaderImpl) LoadAPIKey() (string, error) {
	val := os.Getenv(VirusTotalAPIKey)
	if val == "" {
		return "", fmt.Errorf("environment variable %s is not defined", VirusTotalAPIKey)
	}
	return val, nil
}
