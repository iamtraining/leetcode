package main

func main() {

}

var solution1 = func(read4 func([]byte) int) func([]byte, int) int {
	// implement read below.
	return func(buf []byte, n int) int {
		copiedChars, readChars := 0, 4
		buf4 := make([]byte, 4)
		for copiedChars < n && readChars == 4 {
			readChars = read4(buf4)
			for i := 0; i < readChars; i++ {
				if copiedChars == n {
					return copiedChars
				}
				buf[copiedChars] = buf4[i]
				copiedChars++
			}
		}
		return copiedChars
	}
}

var solution2 = func(read4 func([]byte) int) func([]byte, int) int {
	// implement read below.
	return func(buf []byte, n int) int {
		var data []byte
		b := make([]byte, 4)
		totalChars := 0
		readChars := 4
		for readChars == 4 {
			readChars = read4(b)
			data = append(data, b[:readChars]...)
			totalChars += readChars
		}
		return copy(buf, data[:min(totalChars, n)])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
