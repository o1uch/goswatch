package app

import (
	"errors"
	"os"

	"github.com/o1uch/goswatch/internal/service"
	"github.com/o1uch/goswatch/internal/storage"
)

type StateInterface interface {
	Load() (*service.Stopwatch, error)
	Save(*service.Stopwatch) error
}

func StartApp(state StateInterface) error {
	sw := &service.Stopwatch{}
	var err error
	if format == "yaml" {
		sw, err = storage.LoadYAML()
	} else {
		sw, err = storage.LoadJSON()
	}

	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	err = sw.Start()

	if err != nil {
		return err
	}

	if format == "yaml" {
		err = storage.SaveYAML(sw)
	} else {
		err = storage.SaveJSON(sw)
	}

	return nil

}
