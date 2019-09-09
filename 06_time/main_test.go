package main

import (
	"fmt"
	"testing"
	"time"
)

// ComputeDuration compute the duration between now and a given time
func ComputeDuration(t time.Time) time.Duration {
	now := time.Now()
	return now.Sub(t)
}

func TestComputeDuration(t *testing.T) {
	p := time.Date(2019, 01, 01, 0, 0, 0, 0, time.UTC)
	d := ComputeDuration(p)
	fmt.Println(d)
}

// ComputeDurationModified compute the duration between now and a given time
func ComputeDurationModified(now, t time.Time) time.Duration {
	return now.Sub(t)
}

func TestComputeDurationModified(t *testing.T) {
	now := time.Date(2019, 9, 8, 0, 0, 0, 0, time.UTC)
	p := time.Date(2019, 01, 01, 0, 0, 0, 0, time.UTC)
	d := ComputeDurationModified(now, p)
	fmt.Println(d)
}
