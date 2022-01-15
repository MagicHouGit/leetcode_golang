package main

import "fmt"

func main() {
	releaseTimes := []int{9, 29, 49, 50}
	keysPressed := "cbcd"
	// fmt.Println(slowestKey(releaseTimes, keysPressed))
	fmt.Printf("%q\n", slowestKey(releaseTimes, keysPressed))
	releaseTimes1 := []int{12, 23, 36, 46, 62}
	keysPressed1 := "spuda"
	fmt.Printf("%q\n", slowestKey(releaseTimes1, keysPressed1))
}

func slowestKey(releaseTimes []int, keysPressed string) byte {
	key := []byte(keysPressed)
	t := releaseTimes[0]
	res := key[0]
	for i := 1; i < len(releaseTimes); i++ {

		if releaseTimes[i]-releaseTimes[i-1] > t {
			res = key[i]
			t = releaseTimes[i] - releaseTimes[i-1]
		} else if releaseTimes[i]-releaseTimes[i-1] == t {
			if key[i] > res {
				res = key[i]
			}
		}
	}
	return res
}
