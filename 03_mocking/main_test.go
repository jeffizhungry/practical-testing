package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStudentServiceCreate(t *testing.T) {
	mockDB := NewMockStudentStorage()
	service := NewStudentService(mockDB)
	_, err := service.Create()
	assert.NoError(t, err)
}
