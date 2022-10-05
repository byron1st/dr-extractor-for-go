package lib

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
)

type Config struct {
	TargetSourceCodes []*TargetSourceCode `json:"targetSouceCodes"`
}

type TargetSourceCode struct {
	MainPkgName        string   `json:"mainPkgName"`
	SourceCodePkgNames []string `json:"sourceCodePkgNames"`
}

func SetConfig(path string) (*Config, error) {
	filePath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	configFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	fileContent, err := io.ReadAll(configFile)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(fileContent, &config); err != nil {
		return nil, err
	}

	if len(config.TargetSourceCodes) == 0 || !checkIfConfigValid(config.TargetSourceCodes) {
		return nil, errors.New("wrong config")
	}

	return &config, nil
}

func checkIfConfigValid(targetSourceCodes []*TargetSourceCode) bool {
	for _, targetSourceCode := range targetSourceCodes {
		if targetSourceCode.MainPkgName == "" || len(targetSourceCode.SourceCodePkgNames) == 0 {
			return false
		}
	}

	return true
}
