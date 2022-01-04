package cpn

import(
	"testing"
	"fmt"
)

func TestCheckPerfectNumber(t *testing.T){
	num0 := 1
	num1 := 2
	num2 := 3
	num3 := 6
	num4 := 28
	num5 := 496
	num6 := 497
	fmt.Println(checkPerfectNumber(num0))
	fmt.Println(checkPerfectNumber(num1))
	fmt.Println(checkPerfectNumber(num2))
	fmt.Println(checkPerfectNumber(num3))
	fmt.Println(checkPerfectNumber(num4))
	fmt.Println(checkPerfectNumber(num5))
	fmt.Println(checkPerfectNumber(num6))
}