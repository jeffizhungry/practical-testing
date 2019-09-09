package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateID() string {
	var id string
	for i := 0; i < 8; i++ {
		n := rand.Intn(10)
		id += strconv.Itoa(n)
	}
	return id
}

func main() {
	id := GenerateID()
	fmt.Println("New ID:", id)
}
