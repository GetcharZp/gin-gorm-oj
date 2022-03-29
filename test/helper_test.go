package test

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestRandGenerate(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	s := ""
	for i := 0; i < 6; i++ {
		s += strconv.Itoa(rand.Intn(10))
	}
	println(s)
}
