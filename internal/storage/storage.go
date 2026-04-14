package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gabrielgomes/forgetful/internal/model"
)

func storePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not find home directory: %w", err)
	}
	return filepath.Join(home, ".forg", "commands.json"), nil
}

func Load() (*model.Store, error) {
	path, err := storePath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return &model.Store{NextID: 1}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("could not read store: %w", err)
	}

	var store model.Store
	if err := json.Unmarshal(data, &store); err != nil {
		return nil, fmt.Errorf("could not parse store: %w", err)
	}
	return &store, nil
}

func Save(store *model.Store) error {
	path, err := storePath()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("could not create store directory: %w", err)
	}

	data, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		return fmt.Errorf("could not serialize store: %w", err)
	}

	return os.WriteFile(path, data, 0o644)
}
