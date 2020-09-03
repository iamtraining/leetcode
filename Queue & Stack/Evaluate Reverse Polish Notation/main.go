package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))                                             // 9
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))                                            // 6
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})) // 22
}

type stack []int

func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}
	s := stack{}
	for _, c := range tokens {
		switch c {
		case "+":
			temp := s.pop()
			s.push(temp + s.pop())
		case "-":
			temp := s.pop()
			s.push(s.pop() - temp)
		case "/":
			temp := s.pop()
			s.push(s.pop() / temp)
		case "*":
			temp := s.pop()
			s.push(temp * s.pop())
		default:
			temp, _ := strconv.Atoi(c)
			s.push(temp)
		}
	}
	return s[0]
}

func (this *stack) push(num int) {
	*this = append(*this, num)
}

func (this *stack) pop() int {
	length := len(*this)
	num := (*this)[length-1]
	*this = (*this)[:length-1]
	return num
}
