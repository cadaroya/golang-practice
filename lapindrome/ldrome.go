package main

import "fmt"

func main() {

	// palindrome: racecar
	// lapindrome: accac, abccab (when cut in half, same characters and same frequencies)
	// if ASCII values are the same left and right.

	str := "abccab"
	fmt.Println(str)
	left := 0
	right := 0
	l := len(str)

	for i := 0; i < l/2; i++ {
		left += int(str[i])
		right += int(str[l-i-1])
	}

	// Lapindrome checking
	if left == right {
		fmt.Println("YES it is a lapindrome")
	} else {
		fmt.Println("NO it is NOT a lapindrome")
	}

}
