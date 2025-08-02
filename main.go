package main

import (
	"fmt"
	"time"

	"github.com/o1uch/goswatch/internal/service"
)

func main() {

	fmt.Println("Go")
	var testTime service.Stopwatch
	testTime.Start()
	time.Sleep(3 * time.Second)
	testTime.SaveSplit()
	testTime.Pause()
	time.Sleep(4 * time.Second)
	testTime.Resume()
	testTime.SaveSplit()

	r := testTime.GetResults()

	fmt.Println(r)

}
