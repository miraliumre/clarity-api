package certs

import (
	"os"
	"path/filepath"

	"github.com/miraliumre/clarity-api/settings"
)

func ReadAll(s *settings.Settings) (*[]string, error) {
	dir := s.CertsDir
	contents := []string{}

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		filename := file.Name()
		ext := filepath.Ext(filename)

		if ext == ".cer" || ext == ".crt" || ext == ".pem" {
			data, err := os.ReadFile(filepath.Join(dir, filename))
			if err != nil {
				return nil, err
			}
			contents = append(contents, string(data))
		}
	}

	return &contents, nil
}
