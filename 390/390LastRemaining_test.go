package main

import (
	"fmt"
	"testing"
)

func TestLastRemaining(t *testing.T) {
	t1 := 9
	res1 := lastRemaining(t1)
	if res1 == 4 {
		t.Log("true")
		fmt.Println("fmt-true")
	} else {
		t.Log("false")
		fmt.Println("fmt-flse")
		// t.Error("err")
	}
}

func BenchmarkLastRemaining(b *testing.B) {

	t1 := 9
	res1 := lastRemaining(t1)
	b.Log(res1)
	// b.Error("e")
}
