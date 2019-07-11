package main

import "fmt"

func main() {

	// Program that dynamically computes for the longest increasing sequence
	// Bottom-up, tabulation

	// Slice that will have the LIS calculated
	arr := []int{10, 22, 9, 33, 21, 50, 41, 60}

	// Table for storing values (1D slice)
	lis := make([]int, len(arr))

	// Initialize all values to 1 (every element is LIS)
	for i := range lis {
		lis[i] = 1
	}

	// LIS Proper

	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {

			if arr[i] > arr[j] && lis[j]+1 > lis[i] {
				lis[i] = lis[j] + 1
			}

		}
	}

	// Get Max from lis
	max := 0
	for i := range lis {
		if max < lis[i] {
			max = lis[i]
		}
	}

	fmt.Println(lis)
	fmt.Println("The longest increasing subsequence has length: ", max)

}
