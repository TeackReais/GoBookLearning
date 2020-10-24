package main

import "fmt"

func main() {
	var string1 [4]string = [...]string{"abcd", "", "ghi", "aaeae"}
	var string3 []string
	string3 = string1[:4]
	fmt.Println(string1)
	fmt.Println(string3)
	string3 = nonempty(string3)
	fmt.Println(string1)
	fmt.Println(string3)
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
