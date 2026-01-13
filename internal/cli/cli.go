package cli

import (
	"fmt"
	"os"

	"github.com/o1uch/goswatch/internal/app"
	"github.com/o1uch/goswatch/internal/service"
	"github.com/o1uch/goswatch/internal/storage"
	"github.com/urfave/cli/v2"
)

func stateFlagSelector(ctx *cli.Context) app.StateInterface {
	if ctx.Bool("yaml") {
		return &storage.YamlLoader{}
	}
	return &storage.JsonLoader{}
}

func Run(args []string) int {
	app := &cli.App{
		Name:  "goswatch",
		Usage: "Простой таймер в терминале",
		Commands: []*cli.Command{
			startCommand(),
			resetCommand(),
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
		Action: func(ctx *cli.Context) error {

			state := stateFlagSelector(ctx)

			err := app.StartApp(state)

			if err == service.ErrTimerAlreadyRunning {
				fmt.Println("Таймер уже запущен")
				return nil
			}

			if err != nil {
				fmt.Println("Ошибка выполнения команды: ", err)
				return err
			}

			fmt.Println("таймер запущен")
			return nil
		},
	}
}

func resetCommand() *cli.Command {
	return &cli.Command{
		Name:  "reset",
		Usage: "очистить состояние таймера",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "yaml",
				Usage: "если используется state файл в формате yaml",
			},
		},
		Action: func(ctx *cli.Context) error {

			state := stateFlagSelector(ctx)

			err := app.ResetApp(state)

			if err != nil {
				fmt.Println("При сбросе таймера произошла ошибка:", err)
				return err
			}

			fmt.Println("Таймер сброшен")
			return nil
		},
	}
}
