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
	var string1 [4]string = [4]string{"1232131", "22222", "33333", "$444444"}
	fmt.Println(string1[0:4])
	var string2 [3]string = [...]string{"!23123123123", "a", "b"}
	a := []int{1, 2, 3, 4}
	fmt.Println(string2[0:3])
	b := string2[:] //如果是字符串的切片，则无法修改，如果是字符串组的切片，则可以修改
	b[2] = "c"
	fmt.Println(string2[0:3])
	fmt.Println(cap(a))
	a[3] = 5
	fmt.Println(cap(a))
	fmt.Println(reverse(a))
	//	var shuzu1 [2]int=[2]int{1,2} //无法比较
	//	var shuzu2 [3]int=[2]int{3,4,5}
	//	fmt.Println(shuzu1 == shuzu2)
}
func zero(ptr *[32]byte) { //使用函数令数组为0
	*ptr = [32]byte{}
}
func reverse(s []int) []int { //要求输入的是slice，只有确切的数字填入才是数组
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

//a := []int{1, 2, 3, 4} //slice的定义
//  a := [...]int{1, 2, 3, 4} //数组的定义
