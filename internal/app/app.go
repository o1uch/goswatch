package app

import (
	"errors"
	"os"

	"github.com/o1uch/goswatch/internal/service"
)

// смысл пакета в том, чтобы запустить "таймер" как сценарий. Не знает json это или yaml. Не знает, откуда пришла команда.
// интерфейс для того, чтобы app не знало про storage
type StateInterface interface {
	Load() (*service.Stopwatch, error)
	Save(*service.Stopwatch) error
}

func StartApp(state StateInterface) error {
	sw, err := state.Load()

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			sw = &service.Stopwatch{}
		} else {
			return err
		}
	}

	if err := sw.Start(); err != nil {
		return err
	}

	if err := state.Save(sw); err != nil {
		return err
	}

	return nil
}
