package app

import (
	"errors"
	"os"
	"time"

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

func ResetApp(state StateInterface) error {
	sw, err := state.Load()

	if err != nil {
		return err
	}

	sw.Reset()

	if err := state.Save(sw); err != nil {
		return err
	}

	return nil

}

func PauseApp(state StateInterface) error {
	sw, err := state.Load()

	if err != nil {
		return err
	}

	if err := sw.Pause(); err != nil {
		return err
	}

	if err := state.Save(sw); err != nil {
		return err
	}

	return nil

}

func ResumeApp(state StateInterface) error {
	sw, err := state.Load()

	if err != nil {
		return err
	}

	if err := sw.Resume(); err != nil {
		return err
	}

	if err := state.Save(sw); err != nil {
		return err
	}

	return nil
}

func SaveSplitApp(state StateInterface) error {
	sw, err := state.Load()

	if err != nil {
		return err
	}

	if err := sw.SaveSplit(); err != nil {
		return err
	}

	if err := state.Save(sw); err != nil {
		return err
	}

	return nil
}

func ElapsedApp(state StateInterface) (time.Duration, error) {

	sw, err := state.Load()

	if err != nil {
		return 0, err
	}

	duration := sw.Elapsed()

}
