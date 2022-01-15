package main

import "fmt"

func main() {
	s0 := "(1+(2*3)+((8)/4))+1"
	s1 := "(1)+((2))+(((3)))"
	s2 := "1+(2*3)/(2-1)"
	s3 := "0"
	fmt.Println(maxDepth(s0))
	fmt.Println(maxDepth(s1))
	fmt.Println(maxDepth(s2))
	fmt.Println(maxDepth(s3))
}
func maxDepth(s string) int {
	temp := 0
	heap := []bool{}
	sb := []byte(s)
	for len(heap) > 0 || len(sb) > 0 {
		if sb[0] == '(' {
			heap = append(heap, true)
			if len(heap) > temp {
				temp = len(heap)
			}
		} else if sb[0] == ')' {
			heap = heap[:len(heap)-1]
		}

		if len(sb) > 1 {
			sb = sb[1:]
		} else {
			sb = nil
		}
	}
	return temp
}
