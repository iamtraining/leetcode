package main

func main() {

}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m := len(matrix)
	n := len(matrix[0])
	return helper(0, 0, m-1, n-1, target, matrix)
}

func helper(m1, n1, m2, n2, target int, matrix [][]int) bool {
	if m1 > m2 || n1 > n2 {
		return false
	}
	m3 := (m1 + m2) / 2
	n3 := (n1 + n2) / 2
	if matrix[m3][n3] == target {
		return true
	}
	if target < matrix[m3][n3] {
		return helper(m1, n1, m3-1, n3-1, target, matrix) ||
			helper(m1, n1, m3-1, n2, target, matrix) ||
			helper(m1, n1, m2, n3-1, target, matrix)
	} else {
		return helper(m3+1, n3+1, m2, n2, target, matrix) ||
			helper(m1, n3+1, m3, n2, target, matrix) ||
			helper(m3+1, n1, m2, n3, target, matrix)
	}
}
