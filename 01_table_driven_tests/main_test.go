package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	// Define test cases
	testcases := map[string]struct {
		a        int
		b        int
		expected int
	}{
		"adding positive numbers": {a: 1, b: 2, expected: 3},
		"adding negative numbers": {a: -1, b: -3, expected: -4},
	}

	// Perform tests
	for testName, tc := range testcases {
		actual := add(tc.a, tc.b)
		assert.Equal(t, tc.expected, actual, testName)
	}
}
