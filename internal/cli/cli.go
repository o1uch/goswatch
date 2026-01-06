package cli

import (
	"fmt"
	"os"

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
		Action: func(c *cli.Context) error {
			fmt.Println("таймер запущен")
			return nil
		},
	}
}
