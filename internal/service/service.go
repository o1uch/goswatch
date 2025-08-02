package service

import (
	"fmt"
	"time"
)

type Stopwatch struct {
	// default
	isWorking bool      // старт работы таймера
	startTime time.Time // время начала работы таймера

	// paused
	isPaused       bool          // момент постановки таймера на паузу
	pauseTime      time.Time     // время начала паузы
	pausedDuration time.Duration // время, проведённое на паузе

	split []Split
}

type Split struct {
	checkTime    time.Time // сохраняет промежуточное время
	pausedBefore time.Duration
}

// запустить и сбросить секундомер
func (s *Stopwatch) Start() {
	if s.isWorking {

		// сброс
		s.startTime = time.Time{}
		s.isWorking = false // выставляет статус работы таймера как "не запущен"
		s.split = []Split{}

		s.isPaused = false
		s.pauseTime = time.Time{}
		s.pausedDuration = 0
		return
	}

	s.isWorking = true
	s.startTime = time.Now()
	s.split = []Split{}
}

func (s *Stopwatch) Pause() {
	if !s.isWorking {
		fmt.Println("Timer not started")
		return
	}

	if s.isPaused {
		fmt.Println("the timer is already paused")
		return
	}

	s.isPaused = true
	s.pauseTime = time.Now()
}

func (s *Stopwatch) Resume() {
	if !s.isPaused {
		fmt.Println("the timer is not paused")
		return
	}

	s.pausedDuration += time.Since(s.pauseTime)

	fmt.Println("Timer started. Time spent on pause:", s.pausedDuration.Round(time.Second).String())

	s.isPaused = false
}

// сохранить промежуточное время
func (s *Stopwatch) SaveSplit() {

	if !s.isWorking {
		fmt.Println("the timer is not started")
		return
	} else if s.isPaused {
		fmt.Println("unable to save time - timer paused")
		return
	}

	s.split = append(s.split, Split{checkTime: time.Now(), pausedBefore: s.pausedDuration})

}

// Возвращает время от startTime до текущего момента за вычетом pausedDuration
func (s *Stopwatch) Elapsed() time.Duration {

	// Если таймер на паузе — считает до момента pauseTime
	if s.isPaused {

		passed := s.pauseTime.Sub(s.startTime)
		return passed
	}

	passed := time.Since(s.startTime)
	passed = passed - s.pausedDuration
	return passed
}

// вернуть текущий результат
func (s *Stopwatch) GetResults() []time.Duration {

	if len(s.split) == 0 {
		fmt.Println("Список значений пуст")
		return []time.Duration{0}
	}

	var result []time.Duration
	for _, v := range s.split {
		currValue := v.checkTime.Sub(s.startTime) - v.pausedBefore
		result = append(result, currValue.Round(time.Second))
	}
	return result
}
