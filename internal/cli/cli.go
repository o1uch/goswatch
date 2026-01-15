package cli

import (
	"errors"
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
			pauseCommand(),
			resumeCommand(),
			saveSplitCommand(),
			elapsedСommand(),
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

			if errors.Is(err, service.ErrTimerAlreadyRunning) {
				fmt.Println("Таймер уже запущен")
				return err
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

func pauseCommand() *cli.Command {
	return &cli.Command{
		Name:  "pause",
		Usage: "поставить таймер на паузу",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "yaml",
				Usage: "если используется state файл в формате yaml",
			},
		},
		Action: func(ctx *cli.Context) error {
			state := stateFlagSelector(ctx)

			err := app.PauseApp(state)

			if errors.Is(err, service.ErrTimerNotStarted) {
				fmt.Println("Не удалось остановить. Таймер не запущен")
				return err
			}

			if errors.Is(err, service.ErrTimerAlreadyPaused) {
				fmt.Println("Таймер уже остановлен")
				return err
			}

			if err != nil {
				return err
			}

			return nil
		},
	}

}

func resumeCommand() *cli.Command {
	return &cli.Command{
		Name:  "resume",
		Usage: "снять таймер с паузы",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "yaml",
				Usage: "если используется state файл в формате yaml",
			},
		},
		Action: func(ctx *cli.Context) error {

			state := stateFlagSelector(ctx)

			err := app.ResumeApp(state)

			if errors.Is(err, service.ErrTimerNotPaused) {
				fmt.Println("Таймер не остановлен")
				return err
			}

			if err != nil {
				return err
			}

			return nil
		},
	}
}

func saveSplitCommand() *cli.Command {
	return &cli.Command{
		Name:    "savesplit",
		Aliases: []string{"ss"},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "yaml",
				Usage: "если используется state файл в формате yaml",
			},
		},
		Action: func(ctx *cli.Context) error {
			state := stateFlagSelector(ctx)

			err := app.SaveSplitApp(state)

			if errors.Is(err, service.ErrTimerNotStarted) {
				fmt.Println("невозможно сохранить split. Таймер не запущен")
				return err
			}

			if errors.Is(err, service.ErrCannotSaveSplit) {
				fmt.Println("невозможно сохранить split. Таймер на паузе")
				return err
			}

			if err != nil {
				return err
			}

			return nil
		},
	}
}

func elapsedСommand() *cli.Command {
	return &cli.Command{
		Name:  "elapsed",
		Usage: "вывести время активности таймера",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "second",
				Usage:   "форматирования вывода времени команды elapsed",
				Aliases: []string{"s"},
			},
			&cli.BoolFlag{
				Name:    "minute",
				Usage:   "форматирования вывода времени команды elapsed",
				Aliases: []string{"m"},
			},
			&cli.BoolFlag{
				Name:    "hour",
				Usage:   "форматирования вывода времени команды elapsed",
				Aliases: []string{"h"},
			},
			&cli.BoolFlag{
				Name:  "yaml",
				Usage: "если используется state файл в формате yaml",
			},
		},
		Action: func(ctx *cli.Context) error {
			state := stateFlagSelector(ctx)

			d, err := app.ElapsedApp(state)

			if err != nil {
				return err
			}

			switch {
			case ctx.Bool("second"):
				output := Seconds(d)
				fmt.Println(output)
				return nil
			case ctx.Bool("minute"):
				output := Minutes(d)
				fmt.Println(output)
				return nil
			case ctx.Bool("hour"):
				output := Hours(d)
				fmt.Println(output)
				return nil
			default:
				output := DefaultFormat(d)
				fmt.Println(output)
				return nil
			}

		},
	}
}
