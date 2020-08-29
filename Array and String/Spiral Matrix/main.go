package main

func main() {

}

func spiralOrder(matrix [][]int) []int {
	result := []int{}
	if len(matrix) == 0 {
		return result
	}

	numR := (len(matrix)) / 2
	if numR2 := (len(matrix[0])) / 2; numR2 < numR {
		numR = numR2
	}

	i := 0
	for ; i < numR; i++ {
		for x := i; x < len(matrix[0])-i; x++ {
			result = append(result, matrix[i][x])
		}
		for x := i + 1; x < len(matrix)-i; x++ {
			result = append(result, matrix[x][len(matrix[0])-1-i])
		}
		for x := i + 1; x < len(matrix[0])-i; x++ {
			result = append(result, matrix[len(matrix)-1-i][len(matrix[0])-1-x])
		}
		for x := i + 1; x < len(matrix)-i-1; x++ {
			result = append(result, matrix[len(matrix)-1-x][i])
		}
	}

	oddRows := len(matrix)%2 == 1
	oddCols := len(matrix[0])%2 == 1
	if oddRows && (!oddCols || len(matrix) <= len(matrix[0])) {
		for x := i; x < len(matrix[0])-i; x++ {
			result = append(result, matrix[i][x])
		}
	} else if oddCols {
		for x := i; x < len(matrix)-i; x++ {
			result = append(result, matrix[x][i])
		}
	}

	return result
}
