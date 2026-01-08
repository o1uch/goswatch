package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/o1uch/goswatch/internal/service"
	"gopkg.in/yaml.v3"
)

func SaveJSON(s *service.Stopwatch) error {
	// создаём Persistence file рядом с бинарным
	execFile, err := os.Executable()

	if err != nil {
		return fmt.Errorf("error getting executable file path: %w", err)
	}

	fileName := "state.json"
	stateFile := path.Join(filepath.Dir(execFile), fileName)

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
	execFile, err := os.Executable()

	if err != nil {
		return fmt.Errorf("error getting executable file path: %w", err)
	}

	fileName := "state.yaml"
	stateFile := path.Join(path.Dir(execFile), fileName)

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
