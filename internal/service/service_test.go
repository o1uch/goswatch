package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStopwatch_Start(t *testing.T) {
	sw := new(Stopwatch)

	// первичный запуск
	err := sw.Start()
	assert.NoError(t, err, "Start() не должна возвращать ошибку при первом запуске")
	assert.False(t, sw.startTime.IsZero(), "Stopwatch.startTime не должен быть нулевым при запуске")
	assert.True(t, sw.isWorking, "Stopwatch.isWorking должен быть true при запуске")

	// повторный запуск
	err = sw.Start()

	assert.Error(t, err, "sw.Start() должна возвращать ошибку ErrTimerAlreadyRunning при повторном запуске")
}
