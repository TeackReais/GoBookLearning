package main

import "fmt"

func main() {
	var runes1 []rune
	var runes3 [4]rune = [...]rune{'a', 'b', '我', '你'}
	var runes2 []rune = runes3[0:4]
	var times int
	fmt.Println(runes1 == nil)
	fmt.Println(cap(runes1), len(runes1))
	fmt.Println(cap(runes2), len(runes2))
	for _, b := range "Hello世界" {
		runes1 = append(runes1, b)
		runes2 = runes1
		fmt.Println(cap(runes1), len(runes1))
		times++
	}
	fmt.Println(runes1)
	fmt.Println(times)
	var numbers [4]int = [...]int{1, 2, 3, 4}
	b := numbers[:3]
	fmt.Println(numbers)
	fmt.Println(b)
	numbers[0] = 3
	fmt.Println(numbers)
	fmt.Println(b)
}
