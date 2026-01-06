package main

import (
	"fmt"
	"time"

	"github.com/o1uch/goswatch/internal/service"
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

	testTime := service.Stopwatch{}
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

}
