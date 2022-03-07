package main

import (
	"fmt"
	"testing"
)

func TestWo(t *testing.T) {
	s0 := "2552"
	fmt.Println(isValid(s0))
	s1 := "255"
	fmt.Println(isValid(s1))
}
