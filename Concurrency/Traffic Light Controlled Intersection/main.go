package main

import "fmt"

// channels
var oneDone = make(chan (bool))
var twoDone = make(chan (bool))

func first() {
	fmt.Println("first")
	oneDone <- true
}

func second() {
	<-oneDone
	fmt.Println("second")
	twoDone <- true
}

func third() {
	<-twoDone
	fmt.Println("third")
}

func main() {
	go third()
	go second()
	go first()

	fmt.Scanln()
}

/*
func main() {
	var f, s, t sync.Mutex
	f.Lock()
	s.Lock()
	t.Lock()

	firstDone := sync.NewCond(&f)
	secondDone := sync.NewCond(&s)
	thirdDone := sync.NewCond(&t)

	go third(secondDone, thirdDone)
	go second(firstDone, secondDone)
	go first(firstDone)

	thirdDone.Wait()
}

func first(firstDone *sync.Cond) {
	time.Sleep(time.Second)
	fmt.Println("first")
	firstDone.Broadcast()
}

func second(first, secondDone *sync.Cond) {
	time.Sleep(time.Millisecond)
	first.Wait()
	fmt.Println("second")
	secondDone.Broadcast()
}

func third(second, thirdDone *sync.Cond) {
	time.Sleep(time.Millisecond)
	second.Wait()
	fmt.Println("third")
	thirdDone.Broadcast()
}
*/

/*
func main() {
	var f, s, t sync.Mutex
	f.Lock()
	s.Lock()
	t.Lock()

	firstDone := sync.NewCond(&f)
	secondDone := sync.NewCond(&s)
	thirdDone := sync.NewCond(&t)

	fC := make(chan struct{})
	sC := make(chan struct{})
	tC := make(chan struct{})

	go third(secondDone, thirdDone)
	go second(firstDone, secondDone)
	go first(firstDone)

	thirdDone.Wait()

	go thirdC(sC, tC)
	go secondC(fC, sC)
	go firstC(fC)
	<-tC

}

func first(firstDone *sync.Cond) {
	time.Sleep(time.Second)
	fmt.Println("first")
	firstDone.Broadcast()
}

func second(firstDone, secondDone *sync.Cond) {
	time.Sleep(time.Millisecond)
	firstDone.Wait()
	fmt.Println("second")
	secondDone.Broadcast()
}

func third(secondDone, thirdDone *sync.Cond) {
	time.Sleep(time.Millisecond)
	secondDone.Wait()
	fmt.Println("third")
	thirdDone.Broadcast()
}

func firstC(fC chan struct{}) {
	time.Sleep(time.Second)
	fmt.Println("firstC")
	fC <- struct{}{}
}

func secondC(fC, sC chan struct{}) {
	<-fC
	time.Sleep(time.Millisecond)
	fmt.Println("secondC")
	sC <- struct{}{}
}

func thirdC(sC, tC chan struct{}) {
	<-sC
	time.Sleep(time.Millisecond)
	fmt.Println("thirdC")
	tC <- struct{}{}
}
*/
