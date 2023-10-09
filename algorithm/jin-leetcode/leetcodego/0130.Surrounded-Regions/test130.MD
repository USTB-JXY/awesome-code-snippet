// package main

// import "fmt"

// func main() {
// 	arr := [5]int{0, 1, 2, 3, 4}

// 	slice1 := arr[1:4]
// 	slice2 := arr[2:5]

// 	fmt.Printf("arr %v, slice1 %v, slice2 %v   arr addr: %p, slice1 addr: %p, slice2 addr: %p\n", arr, slice1, slice2, &arr, &slice1, &slice2)

// 	fmt.Printf("arr[2] addr: %p, slice1[1] addr: %p, slice2[0] addr: %p\n", &arr[2], &slice1[1], &slice2[0])

// 	arr[2] = 2222
// 	fmt.Printf("arr: %v, slice1: %v, slice2: %v\n", arr, slice1, slice2)
// }

// // 通过结果我们可以看到，切片作为引用类型，确实是我们在函数内对切片值进行修改时，main函数中的切片值也发生了变化，这不就是引用传递的表现吗？
// // 但是我们发现原切片的地址和函数中切片的地址是不一样的，如果函数切片参数传的是引用，那么main打印的切片地址应该是和test函数切片的地址是一样的，因此我们可以确定的说，go函数中切片参数是传引用的说法是错误的
// // 从上面的分析来看，切片参数传递的方式并非是传引用。反而极有可能是传指针，而在传指针那小节的分析中我们可以知道，传指针其实就是指针的拷贝，形参和实参是两个不同的指针，但是它们的值是一样的。本质上可以说还是传值
