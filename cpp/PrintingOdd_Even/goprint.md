Golang面试题: 两个goroutine交替打印1-100之间的奇数和偶数


题目介绍
       使用两个goroutine交替打印1-100之间的奇数和偶数, 输出时按照从小到大输出.

做法一
package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 101; i++ {
			ch <- struct{}{}
			//奇数
			if i%2 == 1 {
				fmt.Println("线程1打印:",i)
			}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 1; i < 101; i++ {
			<- ch
			//偶数
			if i%2 == 0 {
				fmt.Println("线程2打印:",i)
			}
		}
	}()
	wg.Wait()
}

对做法一的理解:
       首先因为变量ch是一个无缓冲的channel, 所以只有读写同时就绪时才不会阻塞。所以两个goroutine会同时进入各自的 if 语句（此时 i 是相同的），但是此时只能有一个 if 是成立的，不管哪个goroutine快，都会由于读channel或写channel导致阻塞，因此程序会交替打印1-100且有顺序。

做法二
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
		fmt.Println(i)
		ch <- struct{}{}//不能与上一行交换位置
		<- ch
	}
}
func go2() {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		<- ch
		fmt.Println(i)
		ch <- struct{}{}
	}
}
func main() {
	wg.Add(2)
	go go1()
	go go2()
	wg.Wait()
}

对做法二的理解:
       做法二中两个goroutine分别打印奇数和偶数，在两个goroutine中每次for循环时fmt都能得到执行，所以为了实现交替打印1-100，需要使用两对无缓冲的channel。13行和20行构成第一对同步channel，14和22行构成第二对同步channel，第一对同步channel完成读取数据之前，奇数就已经打印；第二对同步channel完成打印之前，偶数就会打印，所以最终输出就是交替打印1-100。

       做法二虽然可以完成交替打印的工作，但是我觉得不具有可扩展性，例如三个协程如何交替打印ABC。后来在网上搜索了一下，发现有比较好的做法，并且也比较容易理解，而且还能扩展到N个协程交替打印数字。

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
		<- ch1
		fmt.Println(i)
		ch2 <- struct{}{}
	}
	<- ch1
}
func go2() {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		<- ch2
		fmt.Println(i)
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

       对于做法二的改进方法，能够做到两个协程交替打印奇偶数，而且还能保证顺序。但是要注意代码16行的“<- ch1”，如果不加的话会报错，因为不加16行的代码，go2最终会被阻塞，wg.Done()无法被执行。

做法三
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//设置可同时使用的CPU核数为1
	var wg sync.WaitGroup
	runtime.GOMAXPROCS(1)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 101; i++ {
			//奇数
			if i%2 == 1 {
				fmt.Println("线程1打印:",i)
			}
			//让出cpu
			runtime.Gosched()
		}
	}()
	go func() {
		defer wg.Done()
		for i := 1; i < 101; i++ {
			//偶数
			if i%2 == 0 {
				fmt.Println("线程2打印:",i)
			}
			//让出cpu
			runtime.Gosched()
		}
	}()
	wg.Wait()
}

       runtime.Gosched这个函数的作用是让当前goroutine让出CPU，好让其它的goroutine获得执行的机会。同时，当前的goroutine也会在未来的某个时间点继续运行。


git clone https://github.com/zeromicro/go-zero.git
git clone https://github.com/go-gorm/gorm.git