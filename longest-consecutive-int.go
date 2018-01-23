package main

import (
	"fmt"
)

//function returns integer in a slice that appears the most consecutive times,
//followed the number of consecutive times it appeared
//In the event of a tie it returns the first one of the equivalent streak
//Returns -1, -1 if there are no consecutive integers
func longestConsecutiveInt(arr []int) (int, int) {
	consec := 0
	ctr := 0
	maxCtr := 0
	maxNum := 0
	for _, n := range arr {
		if consec != n {
			temp := consec
			consec = n
			if maxCtr < ctr {
				maxCtr = ctr
				maxNum = temp
			}
			ctr = 1
		} else {
			ctr++
		}
	}
	if maxCtr <= 1 {
		return -1, -1
	} else {
		return maxNum, maxCtr
	}
}

func main() {
	s := []int{1, 5, 9, 9, 8, 8, 8, 8, 7, 5, 10, 12}
	s2 := []int{5, 5, 5, 5, 5, 5, 5, 6, 7, 8, 9, 10, 10, 10, 10}
	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(longestConsecutiveInt(s))
	fmt.Println(longestConsecutiveInt(s2))
	fmt.Println(longestConsecutiveInt(s3))
}
