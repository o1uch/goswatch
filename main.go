package main

import (
	"fmt"
	"time"

	"github.com/o1uch/goswatch/internal/service"
)

func RunSession(sw service.Stopwatch) {
	sw.Start()
	time.Sleep(1 * time.Second)
	sw.SaveSplit()
	sw.GetResults()
	time.Sleep(1 * time.Second)
	sw.SaveSplit()
	res := sw.GetResults()

	sw.Reset()

	fmt.Println(res)
}

func main() {
	sw := service.Stopwatch{}

	RunSession(sw)
	//os.Exit(cli.Run(os.Args))

}
