package service

import "time"

type Stopwatcher interface {
	Start() error
	Reset()
	Pause() error
	Resume() error
	SaveSplit() error
	Elapsed() time.Duration
	GetResults() []time.Duration
}

var _ Stopwatcher = (*Stopwatch)(nil)

/*
type Stopwatcher interface {
	Start() bool
	Reset()
	Pause() error
	Resume() error
	SaveSplit() error
	Elapsed() time.Duration
	GetResults() []time.Duration
	GetSpentOnPause() time.Duration
	GetStatistics() Stat
	GetTime(...string) (string, error)
}
*/
