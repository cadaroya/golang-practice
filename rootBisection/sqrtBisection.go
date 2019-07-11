package main

import (
	"fmt"
	"math"
)

func main() {

	// Given "x" positive number and a root "n"
	// This program will compute for the "n"th root of a "x"
	// Using the bisection method

	var x float64
	var n int

	fmt.Println("This app will 'n'th root of 'x'")
	fmt.Print("Input value for x: ")
	fmt.Scanf("%f%s", &x)
	fmt.Print("\nInput value for n: ")
	fmt.Scanf("%d%s", &n)

	res := 1.0
	err := 1.00
	high := x
	low := 0.0
	mid := 0.0

	for err > 0.00001 {
		mid = (low + high) / 2
		res = 1

		fmt.Println(low, " ", mid, " ", high)

		for i := 0; i < n; i++ {
			res = res * mid
		}

		if res > x {
			high = mid

		} else if res < x {
			low = mid
		}

		fmt.Println(low, " ", mid, " ", high)

		err = math.Abs((res - x) / x)
		fmt.Println(err)
	}

	fmt.Println("The answer is: ", mid)

}
