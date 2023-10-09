package main

import (
	"fmt"
	"sync"
)

var ch1, ch2 = make(chan struct{}), make(chan struct{})
var wg sync.WaitGroup

func go1() {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		<-ch1
		fmt.Println("goo 111 print ", i)
		ch2 <- struct{}{}
	}
	<-ch1
}
func go2() {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		<-ch2
		fmt.Println("goo 222 print ", i)
		ch1 <- struct{}{}
	}
}
func main() {
	wg.Add(2)
	go go1()
	go go2()
	ch1 <- struct{}{}
	wg.Wait()
}
