// https://www.hackerrank.com/challenges/missing-numbers/problem

package main

import "fmt"

func missingNumbers(arr []int32, brr []int32) []int32 {
	var num1Count [200000]int
	var num2Count [200000]int
	var result = make([]int32, 0)

	for _, num := range arr {
		num1Count[num-1]++
	}

	for _, num := range brr {
		num2Count[num-1]++
	}

	for i := 0; i < 200000; i++ {
		var count1 = num1Count[i]
		var count2 = num2Count[i]

		if count1 != count2 {
			result = append(result, int32(i+1))
		}
	}

	return result
}

func main() {
	fmt.Println(missingNumbers([]int32{203, 204, 205, 206, 207, 208, 203, 204, 205, 206}, []int32{203, 204, 204, 205, 206, 207, 205, 208, 203, 206, 205, 206, 204}))
}
