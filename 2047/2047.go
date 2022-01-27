package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s0 := " cat and  dog"
	s1 := "!this  1-s b8d!"
	s2 := "alice and  bob are playing stone-game10"
	s3 := "he bought 2 pencils, 3 erasers, and 1  pencil-sharpener."
	s4 := "a-b-c"
	s5 := "!g 3 !sy  "
	s6 := " 62   nvtk0wr4f  8 qt3r! w1ph 1l ,e0d 0n 2v 7c.  n06huu2n9 s9   ui4 nsr!d7olr  q-, vqdo!btpmtmui.bb83lf g .!v9-lg 2fyoykex uy5a 8v whvu8 .y sc5 -0n4 zo pfgju 5u 4 3x,3!wl  fv4   s  aig cf j1 a i  8m5o1  !u n!.1tz87d3 .9    n a3  .xb1p9f  b1i a j8s2 cugf l494cx1! hisceovf3 8d93 sg 4r.f1z9w   4- cb r97jo hln3s h2 o .  8dx08as7l!mcmc isa49afk i1 fk,s e !1 ln rt2vhu 4ks4zq c w  o- 6  5!.n8ten0 6mk 2k2y3e335,yj  h p3 5 -0  5g1c  tr49, ,qp9 -v p  7p4v110926wwr h x wklq u zo 16. !8  u63n0c l3 yckifu 1cgz t.i   lh w xa l,jt   hpi ng-gvtk8 9 j u9qfcd!2  kyu42v dmv.cst6i5fo rxhw4wvp2 1 okc8!  z aribcam0  cp-zp,!e x  agj-gb3 !om3934 k vnuo056h g7 t-6j! 8w8fncebuj-lq    inzqhw v39,  f e 9. 50 , ru3r  mbuab  6  wz dw79.av2xp . gbmy gc s6pi pra4fo9fwq k   j-ppy -3vpf   o k4hy3 -!..5s ,2 k5 j p38dtd   !i   b!fgj,nx qgif "
	fmt.Println(countValidWords(s0))
	fmt.Println(countValidWords(s1))
	fmt.Println(countValidWords(s2))
	fmt.Println(countValidWords(s3))
	fmt.Println(countValidWords(s4))
	fmt.Println(countValidWords(s5))
	fmt.Println(countValidWords(s6))
}

// 执行用时：4 ms, 在所有 Go 提交中击败了43.72%的用户
// 内存消耗：4.2 MB, 在所有 Go 提交中击败了8.37%的用户
// func countValidWords(sentence string) int {
// 	res := 0
// 	by := []byte(sentence)
// 	var sen [][]byte
// 	temp := []byte{}
// 	for i := 0; i < len(by); i++ {
// 		if by[i] != ' ' {
// 			temp = append(temp, by[i])
// 		} else {
// 			if len(temp) == 0 {
// 				continue
// 			} else {
// 				sen = append(sen, temp)
// 				temp = []byte{}
// 			}
// 		}
// 	}
// 	if len(temp) > 0 {
// 		sen = append(sen, temp)
// 	}
// 	for _, v := range sen {
// 		tf := true
// 		//false 代表已经有一个'-'了
// 		tfGang := true
// 		for ii := 0; ii < len(v); ii++ {
// 			if v[ii] < 97 || v[ii] > 122 {
// 				if v[ii] == '-' {
// 					if tfGang == false {
// 						tf = false
// 						break
// 					}

// 					if ii == 0 || ii == len(v)-1 {
// 						tf = false
// 						break
// 					} else {
// 						if v[ii-1] < 97 || v[ii-1] > 122 || v[ii+1] < 97 || v[ii+1] > 122 {
// 							tf = false
// 						}
// 					}
// 					tfGang = false
// 				}
// 				if v[ii] == ',' && ii != len(v)-1 {
// 					tf = false
// 					break
// 				}
// 				if v[ii] == '.' && ii != len(v)-1 {
// 					tf = false
// 					break
// 				}
// 				if v[ii] == '!' && ii != len(v)-1 {
// 					tf = false
// 					break
// 				}
// 			}
// 			if v[ii] >= 48 && v[ii] <= 57 {
// 				tf = false
// 				break
// 			}
// 		}
// 		if tf == true {
// 			res++
// 		}
// 	}
// 	return res
// }

//===============================
//other
// github.com/EndlessCheng/codeforces-go
func valid(s string) bool {
	hasHyphens := false
	for i, ch := range s {
		if unicode.IsDigit(ch) || strings.ContainsRune("!.,", ch) && i < len(s)-1 {
			return false
		}
		if ch == '-' {
			if hasHyphens || i == 0 || i == len(s)-1 || !unicode.IsLower(rune(s[i-1])) || !unicode.IsLower(rune(s[i+1])) {
				return false
			}
			hasHyphens = true
		}
	}
	return true
}

func countValidWords(sentence string) (ans int) {
	for _, s := range strings.Fields(sentence) { // 按照空格分割
		if valid(s) {
			ans++
		}
	}
	return
}
