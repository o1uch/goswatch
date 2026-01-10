package main

import (
	"fmt"
	"time"

	"github.com/o1uch/goswatch/internal/service"
	"github.com/o1uch/goswatch/internal/storage"
)

func RunSession(sw service.Stopwatcher) {
	sw.Start()
	time.Sleep(time.Second)
	sw.SaveSplit()
	sw.GetResults()
	sw.Reset()
}

func main() {

	fmt.Println("Go")

	var testTime service.Stopwatch
	testTime.Start()
	time.Sleep(3 * time.Second)
	testTime.SaveSplit()
	testTime.Pause()
	time.Sleep(4 * time.Second)
	testTime.Resume()
	time.Sleep(10 * time.Second)
	testTime.SaveSplit()

	r := testTime.GetResults()
	fmt.Println(r)

	testTime.GetTime()

	err := storage.SaveYAML(&testTime)

	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := storage.LoadYAML()

	fmt.Println(*data)
}
