package cli

import (
	"fmt"
	"os"

	"github.com/o1uch/goswatch/internal/app"
	"github.com/o1uch/goswatch/internal/service"
	"github.com/o1uch/goswatch/internal/storage"
	"github.com/urfave/cli/v2"
)

var state app.StateInterface

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
			if c.Bool("yaml") {
				state = &storage.YamlLoader{}
			} else {
				state = &storage.JsonLoader{}
			}

			err := app.StartApp(state)

			if err == service.ErrTimerAlreadyRunning {
				fmt.Println("Таймер уже запущен")
				return nil
			}

			fmt.Println("таймер запущен")
			return nil
		},
	}
}
