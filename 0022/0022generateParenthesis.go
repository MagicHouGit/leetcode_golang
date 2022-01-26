package main

import "fmt"

func main() {
	// fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(3))
}

// ===================================
//回溯,递归,看了题解后,重写的
// 感觉这严格的说,我感觉不算回溯,
func generateParenthesis(n int) []string {
	var res []string
	var dfs func(l int, r int, path string)
	dfs = func(l int, r int, path string) {
		if l+r == 0 {
			res = append(res, path)
		}
		if l > 0 && r >= l {
			dfs(l-1, r, path+"(")
		}
		if r > 0 {
			dfs(l, r-1, path+")")
		}
	}
	dfs(n, n, "")
	return res
}
