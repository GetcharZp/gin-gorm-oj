package test

import (
	"fmt"
	"getcharzp.cn/helper"
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

func TestCheckGoCodeValid(t *testing.T) {
	valid, err := helper.CheckGoCodeValid("../code/code-user/main.go")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(valid)
}
