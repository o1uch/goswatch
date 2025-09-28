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
	ErrCountValnFunc       = errors.New("count of values in the function is incorrect")
)

// для тестирования
type TimeProvider func() time.Time

type Stopwatch struct {
	// default
	isWorking bool      // старт работы таймера
	startTime time.Time // время начала работы таймера

	// paused
	isPaused       bool          // флаг постановки таймера на паузу
	pauseTime      time.Time     // время начала паузы
	pausedDuration time.Duration // время, проведённое на паузе

	split []Split // структура для фиксации отрезков времени

	//элемент тестирования
	now TimeProvider
}

type Split struct {
	checkTime    time.Time     // сохраняет промежуточное время
	pausedBefore time.Duration // время проведённое на паузе, до сохранения сплита
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

func NewStopwatch(tp TimeProvider) *Stopwatch {
	return &Stopwatch{now: tp}
}

// запустить и сбросить таймер
func (s *Stopwatch) Start() error {
	if s.isWorking {
		return ErrTimerAlreadyRunning
	}

	s.isWorking = true
	s.startTime = NewStopwatch(time.Now())
	s.split = nil
	return nil
}

func (s *Stopwatch) Reset() {
	s.isWorking = false
	s.startTime = time.Time{}

	s.isPaused = false
	s.pauseTime = time.Time{}
	s.pausedDuration = 0
	s.split = nil
}

func (s *Stopwatch) Pause() error {
	if !s.isWorking {
		return ErrTimerNotStarted
	}

	if s.isPaused {
		return ErrTimerAlreadyPaused
	}

	s.isPaused = true
	s.pauseTime = time.Now()
	return nil
}

func (s *Stopwatch) Resume() error {
	if !s.isPaused {
		return ErrTimerNotPaused
	}

	s.pausedDuration += time.Since(s.pauseTime)

	s.isPaused = false
	return nil
}

// зафиксировать время работы таймера в данный момент (без учета паузы)
func (s *Stopwatch) SaveSplit() error {

	if !s.isWorking {
		return ErrTimerNotStarted
	} else if s.isPaused {
		return ErrCannotSaveSplit
	}

	s.split = append(s.split, Split{checkTime: time.Now(), pausedBefore: s.pausedDuration})
	return nil
}

// возвращает время от startTime до текущего момента за вычетом pausedDuration
func (s *Stopwatch) Elapsed() time.Duration {

	// если таймер на паузе — считает до момента pauseTime
	if s.isPaused {

		passed := s.pauseTime.Sub(s.startTime).Round(time.Second)
		return passed
	}

	passed := time.Since(s.startTime).Round(time.Second)
	passed = passed - s.pausedDuration
	return passed
}

// вернуть текущий результат из SaveSplit()
func (s *Stopwatch) GetResults() []time.Duration {

	if len(s.split) == 0 {
		return []time.Duration{}
	}

	var result []time.Duration
	for _, v := range s.split {
		currValue := v.checkTime.Sub(s.startTime) - v.pausedBefore
		result = append(result, currValue.Round(time.Second))
	}
	return result
}

func (s *Stopwatch) GetSpentOnPause() time.Duration {
	return s.pausedDuration.Round(time.Second)
}

func (s Stopwatch) GetStatistics() Stat {

	getSplits := s.GetResults()

	return Stat{
		IsWorking:     s.isWorking,
		IsPaused:      s.isPaused,
		StartTime:     s.startTime,
		PausedTime:    s.pauseTime,
		Elapsed:       s.Elapsed(),
		SplitsCount:   len(s.split),
		Splits:        getSplits,
		AllpausedTime: s.pausedDuration,
	}

}

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
	if s.startTime.Equal(time.Time{}) {
		return "", ErrTimerNotStarted
	}

	return s.startTime.Format(timeLayout), nil
}
