package main

import "fmt"

func main() {
	t0 := "??yw?ipkj?"
	t1 := "i??"
	t2 := "b?b"

	fmt.Println(modifyString(t0))
	fmt.Println(modifyString(t1))
	fmt.Println(modifyString(t2))

}

func modifyString(s string) string {
	l := len(s)
	var r byte
	sb := []byte(s)
	for i := 0; i < l; i++ {
		r = 'a'
		if sb[i] == '?' {
			for (i > 0 && sb[i-1] == r) || (i < l-1 && sb[i+1] == r) {
				r++
			}
			sb[i] = r
		}
	}
	return string(sb)
}

// func modifyString(s string) string {
// 	l := len(s)
// 	res := ""
// 	if l == 0 {
// 		return res
// 	}
// 	sb := []byte(s)
// 	if l == 1 && sb[0] == '?' {
// 		sb[0] = 'a'
// 	}

// 	if sb[0] == '?' {
// 		if sb[1] == 'a' {
// 			sb[0] = 'b'
// 		} else {
// 			sb[0] = 'a'
// 		}
// 	}
// 	if sb[l-1] == '?' {
// 		if sb[l-2] == 'a' {
// 			sb[l-1] = 'b'
// 		} else {
// 			sb[l-1] = 'a'
// 		}
// 	}

// 	for i := 1; i < len(sb)-1; i++ {
// 		temp := 97
// 		if sb[i] == '?' {
// 			if sb[i-1] != 'a' && sb[i+1] != 'a' {
// 				sb[i] = byte(temp)
// 				continue
// 			} else {
// 				temp++
// 			}
// 			if sb[i+1] == 'b' || sb[i-1] == 'b' {
// 				temp++
// 			}
// 			sb[i] = byte(temp)
// 		}
// 	}
// 	return string(sb)
// }
