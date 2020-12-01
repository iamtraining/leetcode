package main

func main() {

}

func totalNQueens(n int) int {
	count := 0
	cols := make([]bool, n)
	d1, d2 := make([]bool, n*2), make([]bool, n*2)
	helper(n, 0, cols, d1, d2, &count)
	return count
}

func helper(n, currRow int, cols, d1, d2 []bool, count *int) {
	if currRow == n {
		*count++
		return
	}

	for i := 0; i < n; i++ {
		idx1, idx2 := n-i+currRow, currRow+i
		if d1[idx1] || d2[idx2] || cols[i] {
			continue
		}
		d1[idx1], d2[idx2], cols[i] = true, true, true
		helper(n, currRow+1, cols, d1, d2, count)
		d1[idx1], d2[idx2], cols[i] = false, false, false
	}
}
