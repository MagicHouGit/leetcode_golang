package main

import "fmt"

func main (){
	test_nums := []int{1,2,3,1}
    fmt.Println(containsDuplicate(test_nums))
}
func containsDuplicate(nums []int) bool {
	if len(nums) == 1{
		return false
	}
	for i:= 0;i < len(nums)-1;i++{
		for j:= i+1; j<len(nums);j++ {
			if nums[i]== nums[j]{
				return true
			}

		}
	}
	return false
}