package main

import (
	"time"
)

type Stopwatch struct {
	isWorking   bool        // фиксирует старт работы таймера
	startTime   time.Time   // фиксирует время начала работы таймера
	currentTime time.Time   // получает текущее время
	checkTime   []time.Time // сохраняет промежуточное время
}

func (s *Stopwatch) Start() {

	if s.isWorking {
		// в случае, если нужно сбросить секундомер
		s.startTime = time.Time{} // не понимаю, как будет выглядеть этот элемент структуры после ресета таймера и как сделать лучше...
		s.isWorking = false       // выставляет статус работы таймера как "не запущен"
		s.checkTime = nil         // очищает сохранённые промежуточные значения
		return
	}

	s.isWorking = true

	s.startTime = time.Now()
}

func (s *Stopwatch) SaveSplit() {

}

// func (s *Stopwatch) GetResults() []time.Duration {}

func main() {

	var testTime Stopwatch

	testTime.Start()

}
