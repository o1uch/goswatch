package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStopwatch_Start(t *testing.T) {
	sw := new(Stopwatch)

	// первичный запуск
	err := sw.Start()
	assert.NoError(t, err, "Start() не должна возвращать ошибку при первом запуске")
	assert.False(t, sw.StartTime.IsZero(), "Stopwatch.startTime не должен быть нулевым при запуске")
	assert.True(t, sw.IsWorking, "Stopwatch.isWorking должен быть true после запуска")

	// повторный запуск
	err = sw.Start()

	assert.Error(t, err, "sw.Start() должна возвращать ошибку ErrTimerAlreadyRunning при повторном запуске")
}

func TestStopwatch_Reset(t *testing.T) {
	s := Stopwatch{}
	s.Start()
	time.Sleep(10 * time.Second)
	s.SaveSplit()
	time.Sleep(5 * time.Second)
	s.SaveSplit()
	s.Pause()
	s.Resume()

	s.Reset()

	assert.False(t, s.IsWorking, "Stopwatch.isWorking должен быть false после Reset()")
	assert.True(t, s.StartTime.IsZero(), "Время начала работы таймера должно быть = 0 после Reset()")

	assert.False(t, s.IsPaused, "Stopwatch.isPaused должен быть false после Reset()")
	assert.True(t, s.PauseTime.IsZero(), "Не должно быть пауз при сброшенном таймере")
	assert.Zero(t, s.PausedDuration, "Продолжительность паузы должна быть = 0 после Reset()")
	assert.Nil(t, s.Split, "После Reset() не должно быть никаких сохранённых промежутков времени")
}
