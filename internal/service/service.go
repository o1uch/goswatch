package service

import (
	"errors"
	"time"
)

var (
	DefaultTimeLayout      = "2006-01-02 15:04"
	ErrTimerNotStarted     = errors.New("the timer is not started")
	ErrTimerAlreadyPaused  = errors.New("the timer is already paused")
	ErrTimerAlreadyRunning = errors.New("the timer is already running")
	ErrTimerNotPaused      = errors.New("the timer is not paused")
	ErrCannotSaveSplit     = errors.New("unable to save time - timer paused")
	ErrListValues          = errors.New("the list of values is empty")
	ErrCountValnFunc       = errors.New("count of values is incorrect")
)

type Stopwatch struct {
	// default
	IsWorking bool      `json:"IsWorking" yaml:"IsWorking"` // старт работы таймера
	StartTime time.Time `json:"StartTime" yaml:"StartTime"` // время начала работы таймера

	// paused
	IsPaused       bool          `json:"IsPaused" yaml:"IsPaused"`             // флаг постановки таймера на паузу
	PauseTime      time.Time     `json:"PauseTime" yaml:"PauseTime"`           // время начала паузы
	PausedDuration time.Duration `json:"PausedDuration" yaml:"PausedDuration"` // время, проведённое на паузе

	Split []Split `json:"Split" yaml:"Split"` // структура для фиксации отрезков времени
}

type Split struct {
	CheckTime    time.Time     `json:"CheckTime" yaml:"CheckTime"`       // сохраняет текущее время на момент сохранения сплита
	PausedBefore time.Duration `json:"PausedBefore" yaml:"PausedBefore"` // время проведённое на паузе, до сохранения сплита
}

type Stat struct {
	IsWorking     bool
	IsPaused      bool
	StartTime     time.Time
	PausedTime    time.Time
	Elapsed       time.Duration
	SplitsCount   int
	Splits        []time.Duration
	AllpausedTime time.Duration
}

// запустить и сбросить таймер
func (s *Stopwatch) Start() error {
	if s.IsWorking {
		return ErrTimerAlreadyRunning
	}

	s.IsWorking = true
	s.StartTime = time.Now()
	s.Split = nil
	return nil
}

func (s *Stopwatch) Reset() {
	s.IsWorking = false
	s.StartTime = time.Time{}

	s.IsPaused = false
	s.PauseTime = time.Time{}
	s.PausedDuration = 0
	s.Split = nil
}

// поставить таймер на паузу
func (s *Stopwatch) Pause() error {
	if !s.IsWorking {
		return ErrTimerNotStarted
	}

	if s.IsPaused {
		return ErrTimerAlreadyPaused
	}

	s.IsPaused = true
	s.PauseTime = time.Now()
	return nil
}

func (s *Stopwatch) Resume() error {
	if !s.IsPaused {
		return ErrTimerNotPaused
	}

	s.PausedDuration += time.Since(s.PauseTime)

	s.IsPaused = false
	return nil
}

// сохранить промежуток времени работы таймера (создать метку времени) (без учета паузы)
func (s *Stopwatch) SaveSplit() error {

	if !s.IsWorking {
		return ErrTimerNotStarted
	} else if s.IsPaused {
		return ErrCannotSaveSplit
	}

	s.Split = append(s.Split, Split{CheckTime: time.Now(), PausedBefore: s.PausedDuration})
	return nil
}

// возвращает время от StartTime до текущего момента за вычетом PausedDuration.
// Считает, сколько времени потрачено.
func (s *Stopwatch) Elapsed() time.Duration {

	// если таймер на паузе — считает до момента PauseTime
	if s.IsPaused {
		passed := s.PauseTime.Sub(s.StartTime).Round(time.Second)
		return passed
	}

	passed := time.Since(s.StartTime).Round(time.Second)
	passed = passed - s.PausedDuration
	return passed
}

// вернуть текущий результат из Stopwatch.Split
// на данный момент времени возврадает в секундах. Как вариант можно доработать для
// cli чтобы можно было указывать флаг с желаемым форматом времени (s, m, h) в output'e
func (s *Stopwatch) GetResults() []time.Duration {

	if len(s.Split) == 0 {
		return []time.Duration{}
	}

	var result []time.Duration
	for _, v := range s.Split {
		currValue := v.CheckTime.Sub(s.StartTime) - v.PausedBefore
		result = append(result, currValue.Round(time.Second))
	}
	return result
}

// тут для cli можно сообразить тоже, что и для GetResults() выше.
func (s *Stopwatch) GetSpentOnPause() time.Duration {
	return s.PausedDuration.Round(time.Second)
}

// тут, как вариант, можно рассмотреть сохранение в разных форматах: yml, json, просто вывод в консоль
func (s Stopwatch) GetStatistics() Stat {

	getSplits := s.GetResults()

	return Stat{
		IsWorking:     s.IsWorking,
		IsPaused:      s.IsPaused,
		StartTime:     s.StartTime,
		PausedTime:    s.PauseTime,
		Elapsed:       s.Elapsed(),
		SplitsCount:   len(s.Split),
		Splits:        getSplits,
		AllpausedTime: s.PausedDuration,
	}

}

// получить время время запуска таймера
// для cli также можно организовать возможность указывать формат времени при вызове функции
func (s *Stopwatch) GetTime(layout ...string) (string, error) {

	// если задано больше одного аругмента
	if len(layout) > 1 {
		return "", ErrCountValnFunc
	}

	// если макет не задан используем дефолтный
	timeLayout := DefaultTimeLayout

	if len(layout) == 1 {
		timeLayout = layout[0]
	}

	// если таймер не запущен
	if s.StartTime.Equal(time.Time{}) {
		return "", ErrTimerNotStarted
	}

	return s.StartTime.Format(timeLayout), nil
}
