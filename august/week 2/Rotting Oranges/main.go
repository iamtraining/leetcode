package main

import "fmt"

func main() {
	grid := [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}
	orangesRotting(grid)
}

func orangesRotting(grid [][]int) {
	m := make(map[int][]map[int]bool)
	for i := 0; i < 3; i++ {
		for _, arg := range grid[i] {
			m[i] = map[int]bool{}
		}

	}
	fmt.Println(m)

}
