package cli

import (
	"errors"
	"fmt"
	"os"

	"github.com/o1uch/goswatch/internal/service"
	"github.com/o1uch/goswatch/internal/storage"
	"github.com/urfave/cli/v2"
)

func Run(args []string) int {
	app := &cli.App{
		Name:  "goswatch",
		Usage: "Простой таймер в терминале",
		Commands: []*cli.Command{
			startCommand(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return 1
	}

	return 0
}

func startCommand() *cli.Command {
	return &cli.Command{
		Name:  "start",
		Usage: "Запустить таймер",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "yaml",
				Usage: "если используется state файл в формате yaml",
			},
		},
		Action: func(c *cli.Context) error {
			sw := &service.Stopwatch{}
			var err error
			if c.Bool("yaml") {
				sw, err = ifTheStateFile("yaml", sw, err)
			} else {
				sw, err = ifTheStateFile("json", sw, err)
			}

			err = sw.Start()
			if err == service.ErrTimerAlreadyRunning {
				fmt.Println("Таймер уже запущен")
				return nil
			}

			fmt.Println("таймер запущен")

			if c.Bool("yaml") {
				storage.SaveYAML(sw)
			} else {
				storage.SaveJSON(sw)
			}

			return nil
		},
	}
}

func ifTheStateFile(name string, sw *service.Stopwatch, err error) (*service.Stopwatch, error) {
	if name == "yaml" {
		sw, err = storage.LoadYAML()
	} else {
		sw, err = storage.LoadJSON()
	}

	if errors.Is(err, os.ErrNotExist) {
		sw = &service.Stopwatch{}
	} else {
		fmt.Println(err)
		return nil, err
	}

	return sw, nil
}
