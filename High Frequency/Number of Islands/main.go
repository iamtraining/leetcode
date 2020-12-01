package main

func main() {

}

func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}

	nr, nc, count := len(grid), len(grid[0]), 0
	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if grid[r][c] == '1' {
				count++
				helper(&grid, r, c)
			}
		}
	}
	return count
}

func helper(grid *[][]byte, r, c int) {
	if r < 0 || c < 0 || r >= len(*grid) || c >= len((*grid)[0]) || (*grid)[r][c] == '0' {
		return
	}
	(*grid)[r][c] = '0'
	helper(grid, r, c+1)
	helper(grid, r+1, c)
	helper(grid, r-1, c)
	helper(grid, r, c-1)
}
