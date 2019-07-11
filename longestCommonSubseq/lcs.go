package main

import "fmt"

func main() {
	// Program that will compute for the longest common subsequence of strings "str1" and "str2"
	// Change and str1 and str2 to test the program
	str1 := "XGXTAB"
	str2 := "GXTXAYB"

	x := len(str1) + 1
	y := len(str2) + 1

	arr := make([][]int, x)
	for i := range arr {
		arr[i] = make([]int, y)
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {

			// Init condition
			if i == 0 || j == 0 {
				arr[i][j] = 0
				// not necessary, just need to continue

				// Equal condition, diagonal recall
			} else if str1[i-1] == str2[j-1] {
				arr[i][j] = arr[i-1][j-1] + 1

				// Not equal condition, left and top recall (max(left,top))
			} else {
				if arr[i][j-1] > arr[i-1][j] {
					arr[i][j] = arr[i][j-1]
				} else {
					arr[i][j] = arr[i-1][j]
				}
			}

		}
	}
	fmt.Println("str1: ", str1)
	fmt.Println("str2: ", str2)
	fmt.Print("     ")
	for i := range str2 {
		fmt.Print(string(str2[i]), " ")
	}
	fmt.Print("\n  ")
	for i := range arr {
		if i < 1 {
			fmt.Println(arr[i])
			continue
		}
		fmt.Print(string(str1[i-1]), " ")
		fmt.Println(arr[i])
	}

	fmt.Println("The length of the longest common subsequence between str1 and str2 is: ", arr[x-1][y-1])

}
