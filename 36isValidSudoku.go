package main

import "fmt"

func main() {
	Tboard := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	isValid := isValidSudoku(Tboard)
	fmt.Println(isValid)

}

func isValidSudoku(board [][]byte) bool {
	// fmt.Printf("%c\n", board[0][0])
	// var gh string = "5"
	// tempM := make(map[int]bool)

	//row col
	for i := 0; i < 9; i++ {
		tempR := make(map[byte]int, 10)
		tempC := make(map[byte]int, 10)
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				tempR[board[i][j]]++
			}
		}
		for j := 0; j < 9; j++ {
			if board[j][i] != '.' {
				tempC[board[j][i]]++
			}
		}
		for _, v := range tempR {
			if v > 1 {
				return false
			}
		}
		for _, v := range tempC {
			if v > 1 {
				return false
			}
		}

	}

	//九块
	var n int = 3
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			tempM := make(map[byte]int, 10)
			for row := i * n; row < (i+1)*n; row++ {
				for col := j * n; col < (j+1)*n; col++ {
					if board[row][col] != '.' {
						tempM[board[row][col]]++
					}
				}
			}
			for _, v := range tempM {
				if v > 1 {
					return false
				}
			}
		}
	}

	return true
}
