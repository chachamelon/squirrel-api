package parser

import (
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
	"api-test-engine/internal/model"
)

func LoadTestCases(path string) ([]*model.TestCase, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if !info.IsDir() {
		// Single file
		tc, err := loadTestCaseFromFile(path)
		if err != nil {
			return nil, err
		}
		return []*model.TestCase{tc}, nil
	}

	// Directory - load all yaml files
	var testCases []*model.TestCase
	err = filepath.WalkDir(path, func(filePath string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(strings.ToLower(filePath), ".yaml") {
			tc, err := loadTestCaseFromFile(filePath)
			if err != nil {
				return err
			}
			testCases = append(testCases, tc)
		}
		return nil
	})

	return testCases, err
}

func loadTestCaseFromFile(path string) (*model.TestCase, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var tc model.TestCase
	err = yaml.Unmarshal(data, &tc)
	return &tc, err
}
