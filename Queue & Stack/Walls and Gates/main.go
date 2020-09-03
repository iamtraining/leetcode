package main

import "math"

func main() {

}

type Cell struct {
	i, j int
}

func wallsAndGates(rooms [][]int) {
	q := []Cell{}
	for i := 0; i < len(rooms); i++ {
		for j := 0; j < len(rooms[i]); j++ {
			if rooms[i][j] == 0 {
				q = append(q, Cell{i, j})
			}
		}
	}

	for len(q) != 0 {
		size := len(q)
		for k := 0; k < size; k++ {
			cell := q[0]
			q = q[1:]

			i, j := cell.i, cell.j
			//down
			for i < len(rooms)-1 && rooms[i+1][j] == math.MaxInt32 {
				q = append(q, Cell{i + 1, j})
				rooms[i+1][j] = rooms[i][j] + 1
			}
			//up
			for i > 0 && rooms[i-1][j] == math.MaxInt32 {
				q = append(q, Cell{i - 1, j})
				rooms[i-1][j] = rooms[i][j] + 1
			}
			//right
			for j < len(rooms[i])-1 && rooms[i][j+1] == math.MaxInt32 {
				q = append(q, Cell{i, j + 1})
				rooms[i][j+1] = rooms[i][j] + 1
			}
			//left
			for j > 0 && rooms[i][j-1] == math.MaxInt32 {
				q = append(q, Cell{i, j - 1})
				rooms[i][j-1] = rooms[i][j] + 1
			}
		}
	}
}
