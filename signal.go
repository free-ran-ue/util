package util

import (
	"fmt"
	"time"
)

func SignalComplete(ch chan struct{}) {
	select {
	case ch <- struct{}{}:
	default:
	}
}

func WaitComplete(ch chan struct{}, timeout time.Duration) error {
	select {
	case <-ch:
		return nil
	case <-time.After(timeout):
		return fmt.Errorf("timed out waiting for completion")
	}
}
