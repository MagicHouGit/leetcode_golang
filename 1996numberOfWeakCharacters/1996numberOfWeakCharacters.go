package main

import (
	"fmt"
	"sort"
)

func main() {
	// p0 := [][]int{{5, 5}, {6, 3}, {3, 6}}
	// p1 := [][]int{{2, 2}, {3, 3}}
	// p2 := [][]int{{1, 5}, {10, 4}, {4, 3}}
	// p3 := [][]int{{7, 7}, {1, 2}, {9, 7}, {7, 3}, {3, 10}, {9, 8}, {8, 10}, {4, 3}, {1, 5}, {1, 5}}
	p4 := [][]int{{7, 9}, {10, 7}, {6, 9}, {10, 4}, {7, 5}, {7, 10}}
	// fmt.Println(numberOfWeakCharacters(p0))
	// fmt.Println(numberOfWeakCharacters(p1))
	// fmt.Println(numberOfWeakCharacters(p2))
	// fmt.Println(numberOfWeakCharacters(p3))
	fmt.Println(numberOfWeakCharacters(p4))
	// fmt.Println(numberOfWeakCharactersOther(p4))
	// fmt.Println(numberOfWeakCharactersOther(p3))
}

///other
// performance over 100%
func numberOfWeakCharacters(properties [][]int) int {
	maxAttack := 0
	for _, x := range properties {
		if maxAttack < x[0] {
			maxAttack = x[0]
		}
	}

    attacks := make([]int,maxAttack + 1)
	for _, p := range properties {
        if  attacks[p[0]] < p[1]{
            attacks[p[0]] = p[1]
        }
	}

	maxDefence := 0
	for i := maxAttack; i >= 0; i-- {
		temp := maxDefence
		if attacks[i] > maxDefence {
			maxDefence = attacks[i]
		}
		attacks[i] = temp
	}
	ans := 0
	for _, p := range properties {
		if p[0] < maxAttack && p[1] < attacks[p[0]] {
			ans++
		}
	}
	return  ans
}


// //===============================
// // 执行用时：400 ms, 在所有 Go 提交中击败了10.61%的用户
// // 内存消耗：20.7 MB, 在所有 Go 提交中击败了53.18%的用户
// func numberOfWeakCharacters(properties [][]int) int {
// 	sort.Slice(properties, func(i, j int) bool {
// 		a, b := properties[i], properties[j]
// 		return a[0] > b[0] || (a[0] == b[0] && a[1] < b[1])
// 	})
// 	fmt.Println(properties)
// 	res := 0
// 	max := 0
// 	for i := 0; i < len(properties); i++ {
// 		if properties[i][1] >= max {
// 			max = properties[i][1]
// 		} else {
// 			res++
// 		}
// 	}
// 	return res
}

//=================================
// other
// func numberOfWeakCharactersOther(properties [][]int) int {
// 	var ans int
// 	sort.Slice(properties, func(i, j int) bool {
// 		p, q := properties[i], properties[j]
// 		return p[0] > q[0] || p[0] == q[0] && p[1] < q[1]
// 	})
// 	fmt.Println("less", properties)
// 	maxDef := 0
// 	for _, p := range properties {
// 		if p[1] < maxDef {
// 			ans++
// 		} else {
// 			maxDef = p[1]
// 		}
// 	}
// 	return ans
// }

//-===========================
// 超时警告
// func numberOfWeakCharacters(properties [][]int) int {
// 	res := 0
// 	for i := 0; i < len(properties); i++ {
// 		for ii := 0; ii < len(properties); ii++ {
// 			if i == ii {
// 				continue
// 			}
// 			if properties[i][0] < properties[ii][0] && properties[i][1] < properties[ii][1] {
// 				res++
// 				properties = append(properties[:i], properties[i+1:]...)
// 				i--
// 				break
// 			}
// 		}
// 	}
// 	return res
// }
