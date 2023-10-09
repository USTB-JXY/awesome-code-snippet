package main

import (
	"fmt"
	"sync"
)

var ch = make(chan struct{})
var wg sync.WaitGroup

func go1() {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		fmt.Println("go1111 print ", i)
		ch <- struct{}{} //不能与上一行交换位置
		<-ch
	}
}
func go2() {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		<-ch
		fmt.Println("go2222 print ", i)
		ch <- struct{}{}
	}
}
func main() {
	wg.Add(2)
	go go1()
	go go2()
	wg.Wait()
}
