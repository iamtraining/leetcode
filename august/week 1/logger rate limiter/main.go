package main

import "fmt"

func main() {
	obj := Constructor()
	fmt.Println(obj.ShouldPrintMessage(1, "foo"))  // t
	fmt.Println(obj.ShouldPrintMessage(2, "bar"))  // t
	fmt.Println(obj.ShouldPrintMessage(3, "foo"))  // f
	fmt.Println(obj.ShouldPrintMessage(8, "bar"))  // f
	fmt.Println(obj.ShouldPrintMessage(10, "foo")) // f
	fmt.Println(obj.ShouldPrintMessage(11, "foo")) // t
}

type Logger struct {
	Map map[string]int
}

/** Initialize your data structure here. */
func Constructor() Logger {
	return Logger{Map: make(map[string]int)}
}

/** Returns true if the message should be printed in the given timestamp, otherwise returns false.
  If this method returns false, the message will not be printed.
  The timestamp is in seconds granularity. */
func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if _, ok := this.Map[message]; !ok {
		this.Map[message] = timestamp
		return true
	}
	t := this.Map[message]
	if timestamp-t < 10 {
		return false
	}
	this.Map[message] = timestamp
	return true
}

/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */
