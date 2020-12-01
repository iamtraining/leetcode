package main

import (
	"fmt"
	"sync"
)

var (
	mu   sync.Mutex
	curr int
)

func main() {
	curr = 1
	n := 15

	var wg sync.WaitGroup
	wg.Add(4)
	go num(n, &wg)
	go fizz(n, &wg)
	go buzz(n, &wg)
	go fizzbuzz(n, &wg)
	wg.Wait()
}

// divisible by 3
func fizz(n int, wg *sync.WaitGroup) {
	for curr <= n {
		if curr%3 == 0 && curr%5 != 0 {
			mu.Lock()
			fmt.Print("fizz ")
			curr++
			mu.Unlock()
		}
	}
	wg.Done()
}

// divisible by 5
func buzz(n int, wg *sync.WaitGroup) {
	for curr <= n {
		if curr%3 != 0 && curr%5 == 0 {
			mu.Lock()
			fmt.Print("buzz ")
			curr++
			mu.Unlock()
		}
	}
	wg.Done()
}

// divisible by both 3 and 5
func fizzbuzz(n int, wg *sync.WaitGroup) {
	for curr <= n {
		if curr%3 == 0 && curr%5 == 0 {
			mu.Lock()
			fmt.Print("fuzzbuzz ")
			curr++
			mu.Unlock()
		}
	}
	wg.Done()
}

func num(n int, wg *sync.WaitGroup) {
	for curr <= n {
		if curr%3 != 0 && curr%5 != 0 {
			mu.Lock()
			fmt.Print(curr, " ")
			curr++
			mu.Unlock()
		}
	}
	wg.Done()
}

// If the number is divisible by 3, output "fizz".
// If the number is divisible by 5, output "buzz".
// If the number is divisible by both 3 and 5, output "fizzbuzz".
