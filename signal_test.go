package util_test

import (
	"testing"
	"time"

	"github.com/free-ran-ue/util"
	"github.com/go-playground/assert/v2"
)

func TestSignalComplete_DeliversToWaiter(t *testing.T) {
	ch := make(chan struct{}, 1)

	util.SignalComplete(ch)

	err := util.WaitComplete(ch, 50*time.Millisecond)
	assert.Equal(t, nil, err)
}

func TestSignalComplete_RepeatedSignalDoesNotBlockOrPanic(t *testing.T) {
	ch := make(chan struct{}, 1)

	util.SignalComplete(ch)
	util.SignalComplete(ch) // must not block and must not panic

	err := util.WaitComplete(ch, 50*time.Millisecond)
	assert.Equal(t, nil, err)

	// only one signal should have been buffered
	select {
	case <-ch:
		t.Fatalf("expected no second signal buffered")
	default:
	}
}

func TestWaitComplete_TimesOutWhenNeverSignalled(t *testing.T) {
	ch := make(chan struct{}, 1)

	start := time.Now()
	err := util.WaitComplete(ch, 20*time.Millisecond)
	elapsed := time.Since(start)

	assert.NotEqual(t, nil, err)
	assert.Equal(t, true, elapsed >= 20*time.Millisecond)
}

func TestWaitComplete_UnblocksWhenSignalledConcurrently(t *testing.T) {
	ch := make(chan struct{}, 1)

	go func() {
		time.Sleep(10 * time.Millisecond)
		util.SignalComplete(ch)
	}()

	err := util.WaitComplete(ch, 200*time.Millisecond)
	assert.Equal(t, nil, err)
}
