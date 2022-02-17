package main

import (
	"fmt"
	"testing"
)

func TestStringAdd(t *testing.T) {

	x0 := "63864"
	x1 := "92653863"
	fmt.Println(stringAdd(x0, x1))
	// if stringAdd(x0, x1) != "92717727" {
	// 	t.Error("err")
	// }
	fmt.Println("88888888888")
	fmt.Println(stringAdd("1", "9"))
	fmt.Println(stringAdd("99", "100"))
	// t.Error("fff")
}
