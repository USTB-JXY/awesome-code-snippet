package main

import (
	"fmt"
	"sync"
)

var ch1, ch2, ch3 = make(chan struct{}), make(chan struct{}), make(chan struct{})
var wg sync.WaitGroup

func go1() {
	defer wg.Done()
	for i := 1; i <= 100; i += 3 {
		<-ch1
		fmt.Println("goo 111 print ", i)
		if i < 100 {
			ch2 <- struct{}{}
		}

	}

}
func go2() {
	defer wg.Done()
	for i := 2; i <= 100; i += 3 {
		<-ch2
		fmt.Println("goo 222 print ", i)
		ch3 <- struct{}{}
	}

}
func go3() {
	defer wg.Done()
	for i := 3; i <= 100; i += 3 {
		<-ch3
		fmt.Println("goo 333 print ", i)
		ch1 <- struct{}{}
	}
}
func main() {
	wg.Add(3)
	go go1()
	go go2()
	go go3()
	ch1 <- struct{}{}
	wg.Wait()
}
