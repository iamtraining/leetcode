package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(numbers []int, target int) []int {
	i, j, l := 0, len(numbers)-1, len(numbers)
	for i < j {
		if numbers[i]+numbers[j] == target {
			numbers = append(numbers, i+1, j+1)
			return numbers[l:]
		} else if numbers[i]+numbers[j] < target {
			i++
		} else {
			j--
		}
	}
	numbers = append(numbers, -1, -1)
	return numbers[l:]
}
