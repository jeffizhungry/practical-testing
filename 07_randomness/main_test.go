package main

import (
	"fmt"
	"math/rand"
	"testing"

	"gotest.tools/assert"
)

func TestGenerateFiveNumbers(t *testing.T) {
	// Seed in testcase
	rand.Seed(0)

	s := GenerateID()
	fmt.Println(s)
	assert.Equal(t, "44365677", s)
}
