package main

import (
	"fmt"
	"strconv"
)

func main() {

}

func largestTimeFromDigits(A []int) string {
	max := ""
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if j == i {
				continue
			}
			for k := 0; k < 4; k++ {
				if k == i || k == j {
					continue
				}
				for m := 0; m < 4; m++ {
					if m == i || m == j || m == k {
						continue
					}
					if curr := helper(A[i], A[j], A[k], A[m]); curr > max {
						max = curr
					}
				}

			}
		}
	}
	return max
}

func helper(i, j, k, m int) string {
	hour := i*10 + j
	if hour >= 24 {
		return ""
	}
	min := k*10 + m
	if min >= 60 {
		return ""
	}
	return fmt.Sprintf("%02s:%02s", strconv.Itoa(hour), strconv.Itoa(min))
}
