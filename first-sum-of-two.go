package main

import (
	"fmt"
)

func sumFinder(nums []int, target int) []int {
	numChecker := make(map[int]int)
	answer := []int{}
	for i, x := range nums {
		numChecker[i] = x
		for j := 0; j < len(numChecker); j++ {
			if nums[i] + numChecker[j] == target {
				answer = append(answer, numChecker[j], nums[i])
				return answer
			}
		}
	}
	return answer // if target sum not found, return empty arr
}


func main() {
	s := []int{1, 5, 8, 19, 4, 12, 6, 7, 14, 42}
	fmt.Println(sumFinder(s, 17)) // will return 5, 12
}
