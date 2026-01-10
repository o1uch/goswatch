package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/o1uch/goswatch/internal/service"
	"gopkg.in/yaml.v3"
)

func stateFilePath(fName string) (string, error) {
	execFile, err := os.Executable()

	if err != nil {
		return "", fmt.Errorf("error getting executable file path: %w", err)
	}

	stateFilePath := filepath.Join(filepath.Dir(execFile), fName)

	return stateFilePath, nil
}

func SaveJSON(s *service.Stopwatch) error {
	// создаём persistence file рядом с бинарным

	if s == nil {
		return fmt.Errorf("Stopwatch is nil")
	}

	stateFile, err := stateFilePath("state.json")

	if err != nil {
		return err
	}

	data, err := json.Marshal(s)

	if err != nil {
		return fmt.Errorf("serialization error: %w", err)
	}

	// т.к. файл будет содержать исключительно актульное состояние Stopwatch, то нет смысла переусложнять логику. Файл будет просто пересоздаваться

	err = os.WriteFile(stateFile, data, 0644)

	if err != nil {
		return fmt.Errorf("file creation error: %w", err)
	}

	return nil
}

func SaveYAML(s *service.Stopwatch) error {

	if s == nil {
		return fmt.Errorf("Stopwatch is nil")
	}

	stateFile, err := stateFilePath("state.yaml")

	if err != nil {
		return err
	}

	data, err := yaml.Marshal(s)

	if err != nil {
		return fmt.Errorf("serialization error: %w", err)
	}

	err = os.WriteFile(stateFile, data, 0644)

	if err != nil {
		return fmt.Errorf("file creation error: %w", err)
	}

	return nil
}

func LoadJSON() (*service.Stopwatch, error) {
	sw := &service.Stopwatch{}

	stateFile, err := stateFilePath("state.json")

	if err != nil {
		return nil, err
	}

	raw, err := os.ReadFile(stateFile)
	if err != nil {
		return nil, fmt.Errorf("error reading the state-file:%w", err)
	}

	err = json.Unmarshal(raw, sw)

	if err != nil {
		return nil, fmt.Errorf("data conversion error:%w", err)
	}

	return sw, nil
}

func LoadYAML() (*service.Stopwatch, error) {
	sw := &service.Stopwatch{}

	stateFile, err := stateFilePath("state.yaml")

	if err != nil {
		return nil, err
	}

	raw, err := os.ReadFile(stateFile)

	if err != nil {
		return nil, fmt.Errorf("error reading the state-file: %w", err)
	}

	err = yaml.Unmarshal(raw, sw)

	if err != nil {
		return nil, fmt.Errorf("data conversion error:%w", err)
	}

	return sw, nil

}
