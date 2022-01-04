
package cpn

func checkPerfectNumber(num int) bool {
	if num<=1{
		return false
	}
	var count []int
	count = append(count,1)
	for i:=2;i*i<=num;i++{
		if num%i == 0{
			count = append(count,i)
			v := num/i
			count = append(count,v)
		}
	}
	for _,v:= range count{
		num -= v
	}
	if num == 0{
		return true
	}
	return false
}