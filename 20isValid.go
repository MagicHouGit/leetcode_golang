package main

import "fmt"

func main() {
	// test := "()"
	test := "([}])"
	// test := "({[{]})"
	// test := "{()()}"
	// test := "{()[()}]"
	// test := "]"
	fmt.Println(isValid(test))
}

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2.1 MB, 在所有 Go 提交中击败了19.90%的用户
func isValid(s string) bool {
	var g []int
	for _, v := range s {

		switch v {
		case '(':
			g = append(g, 1)
			continue
		case '[':
			g = append(g, 2)
			continue
		case '{':
			g = append(g, 3)
			continue
		}
		if len(g) == 0 {
			// break
			// continue
			return false
		}
		if v == ')' && g[len(g)-1] == 1 {
			g[len(g)-1] = 0
			g = g[:len(g)-1]
			continue
		}
		if v == ']' && g[len(g)-1] == 2 {
			g[len(g)-1] = 0
			g = g[:len(g)-1]
			continue
		}

		if v == '}' && g[len(g)-1] == 3 {
			g[len(g)-1] = 0
			g = g[:len(g)-1]
			continue
		}

		return false
	}
	if len(g) == 0 {
		return true
	}
	return false

}
