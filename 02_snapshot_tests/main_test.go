package main

import (
	"strconv"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
)

type student struct {
	id   string
	name string
	age  int
}

func newStudent(name string, age int) student {
	return student{
		id:   name + "-" + strconv.Itoa(age),
		name: name,
		age:  age,
	}
}

func TestNewStudent(t *testing.T) {
	actual := newStudent("erica", 22)
	cupaloy.SnapshotT(t, actual)
}
