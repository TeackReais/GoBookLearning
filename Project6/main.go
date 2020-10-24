package main

import "fmt"

func main() {
	var number [32]int = [32]int{3, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 11, 12, 13, 14, 15, 16, 17, 18}
	fmt.Println(number)
	//	zero(&number)
	fmt.Println(number)
	fmt.Println(cap(number), len(number))
	fmt.Println(cap(number[1:5]), len(number[1:5]))
	fmt.Println(cap(number[2:5]), number[2:32])
	fmt.Println(reverse(number[:]))
}
func zero(ptr *[32]byte) {
	*ptr = [32]byte{}
}
func reverse(s []int) []int { //要求输入的是切片，只有确切的数字填入才是数组
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
