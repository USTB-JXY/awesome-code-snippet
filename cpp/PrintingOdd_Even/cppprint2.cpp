#include <iostream>
#include <thread>
#include <unistd.h>

using namespace std;



// jxy 不知道是否安全
// volatile

// 说道c++多线程，volatile就很容易被提起，它的作用有：
// 1. 告诉编译器，这个变量是易变的，当编译器遇到这个变量的时候，只能从变量的内存地址中读取这个变量，不可以从寄存器读取
// 2. 告诉编译器不要将变量优化掉，保证这个指令一定会被执行
// 3. 两个包含volatile变量的指令，编译后不可以乱序（编译器屏障）。但是实际执行时CPU可能打乱顺序

// volatile关键字只针对编译器，对CPU没有作用，所以关于volatile在多线程中的作用的结论是：
// 1. 不保证原子性
// 2. 不保障实际执行顺序
// 3. 不提供CPU内存屏障



const int MAXNUM = 1000;    // 输出范围：1 - MAXNUM
int number;
int sum=0;
int sumsafe=0;
// 打印奇数
void add_1() {
    while (1) {
        if (number % 2 == 0) {
            int temp_number = number;       // 读出 number
            temp_number = temp_number + 1;  // 加1
            sum++;
            cout << "mythread_1: " << (++number) << endl;    // 输出
            number = temp_number;           // 写回 number

            int tempsum=sumsafe;
            tempsum=tempsum+1;
            sumsafe=tempsum;
            if (temp_number == MAXNUM - 1) {
                break;
            }
        }
    }
    cout  << "mythread_1 finish" << endl;     // mythread_1完成
}

// 打印偶数
void add_2() {
    while (1) {
        if (number % 2 == 1) {
            int temp_number = number;       // 读出 number
            temp_number = temp_number + 1;  // 加1
            cout << "mythread_2: " << (++number) << endl;    // 输出
            sum++;
            number = temp_number;           // 写回 number
            int tempsum=sumsafe;
            tempsum=tempsum+1;
            sumsafe=tempsum;
            if (temp_number == MAXNUM) {
                break;
            }
        }
    }
    cout  << "mythread_2 finish" << endl;     // mythread_2完成
}

int main() {
    number = 0;
    cout << endl << "Create and Start!" << endl;
    thread mythread_1(add_1);	// 打印奇数
    thread mythread_2(add_2);	// 打印偶数

    mythread_1.join();
    mythread_2.join();
    cout << endl <<" sum = "<<sum<< "  Finish and Exit!" << endl;
    cout << endl <<" sumsafe = "<<sumsafe<< "  Finish and Exit!" << endl;
    return 0;
}


