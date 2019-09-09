package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func add(a, b int) int {
	return a + b
}

func TestAddPositiveNumbers(t *testing.T) {

	// .. maybe other setup

	actual := add(1, 2)
	assert.Equal(t, 3, actual)
}

func TestAddNegativeNumbers(t *testing.T) {

	// .. maybe other setup

	actual := add(-1, -3)
	assert.Equal(t, -4, actual)
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

		// .. maybe other setup

		actual := add(tc.a, tc.b)
		assert.Equal(t, tc.expected, actual, testName)
	}
}
