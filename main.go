package main

import (
	"os"
	"time"

	"github.com/o1uch/goswatch/internal/cli"
	"github.com/o1uch/goswatch/internal/service"
)

func RunSession(sw service.Stopwatcher) {
	sw.Start()
	time.Sleep(1 * time.Second)
	sw.SaveSplit()
	sw.GetResults()
	sw.Reset()
}

func main() {

	os.Exit(cli.Run(os.Args))
}
