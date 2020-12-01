package main

func main() {

}

func getRow(rowIndex int) []int {
	switch rowIndex {
	case 0:
		return []int{1}
	case 1:
		return []int{1, 1}
	default:
		return helper([]int{1, 1}, 2, rowIndex)
	}
}

func helper(prev []int, curr int, target int) []int {
	row := make([]int, curr+1)
	row[0], row[len(row)-1] = 1, 1
	for i := 1; i < curr; i++ {
		row[i] = prev[i-1] + prev[i]
	}
	if curr == target {
		return row
	} else {
		return helper(row, curr+1, target)
	}
}
