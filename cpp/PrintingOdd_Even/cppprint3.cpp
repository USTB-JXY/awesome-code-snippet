#include <iostream>
#include <thread>
#include <mutex>
#include <condition_variable>

std::mutex mu;
std::condition_variable cond;
int count = 1;
using namespace std;

class Test
{
public:
    Test()
    {   
        cout<<m_s<< " Defalut Constructor executed\n ";
    }

    Test(string s)
    {   
        m_s=s;
        cout<<s<< " Constructor executed\n ";
    }

    ~Test()
    {
        cout<<m_s << "Destructor ----------- executed\n ";
    }
    string m_s="Defalut a ";
};

void PrintOdd()
{
    for(; count < 100;)
    {
        Test A;
        Test* B =new Test("B");
        std::unique_lock<std::mutex> locker(mu);
        cond.wait(locker,[](){ return (count%2 == 1); });
        std::cout << "From Odd:    " << count << std::endl;
        count++;
        //locker.unlock();
        cond.notify_all();
    }

}

void PrintEven()
{
    for(; count < 100;)
    {
        std::unique_lock<std::mutex> locker(mu);
        cond.wait(locker,[](){ return (count%2 == 0); });
        std::cout << "From Even: " << count << std::endl;
        count++;
        //locker.unlock();
        cond.notify_all();
    }
}

int main()
{
    std::thread t1(PrintOdd);
    std::thread t2(PrintEven);
    t1.join();
    t2.join();
    return 0;
}
// Wouldn't this lead to a deadlock if the t2 thread executes first? PrintEven() acquires the lock then waits for count to become even. But count is 1, and it will become even only when PrintOdd() executes count++. PrintOdd() is waiting for the lock, so cannot execute.
// 如果首先执行 t2 线程，这不会导致死锁吗？ PrintEven() 获取锁，然后等待计数变得偶数。但是 count 就是 1，只有当 PrintOdd() 执行 count++ 时，它才会变得偶数。 PrintOdd() 正在等待锁定，因此无法执行。
// – 
//  –
// Masked Man
//  Jun 30, 2018 at 6:46
// @MaskedMan: No it'll not. That is what condition variable is doing. It releases the lock if the condition is not met and goes to sleep kind of mode. So if t2 acquires the lock first, then it will go into waiting state since coun%2 == 0 condition will fail. When t2 releases the lock, t1 thread can acquire the lock and it can execute. Once it finishes executing, it calls notify_all(), which will notify the t2 thread saying now you can check if the condition is met or not, If the condition is met, it re-acquires the lock and print the even number and the sequence continues.
// @MaskedMan：不，不会。这就是条件变量正在做的事情。如果不满足条件，它会释放锁并进入睡眠模式。因此，如果 t2 首先获取锁，那么它将进入等待状态，因为 coun%2 == 0 条件将失败。当 t2 释放锁时，t1 线程可以获取锁并可以执行。完成执行后，它会调用 notify_all（），这将通知 t2 线程，说现在您可以检查是否满足条件，如果满足条件，它会重新获取锁并打印偶数，序列继续。


// does locker.unlock(); required for unique_lock? isnt unlock automatically?
// locker.unlock（）;需要unique_lock吗？不是自动解锁吗？
// – 
//  –
// Pravej Khan
//  Feb 10, 2022 at 4:54
// @PravejKhan, unlock will happen only if the locker goes out of scope. In this case, since we are still in the loop, the lock will still be applied. That is why we need to unlock manually.
// @PravejKhan，只有当储物柜超出范围时，才会解锁。在这种情况下，由于我们仍在循环中，因此仍将应用锁定。这就是为什么我们需要手动解锁。



// std::unique_lock use the RAII pattern.
// std::unique_lock 使用 RAII 模式。

// When you want to lock a mutex, you create a local variable of type std::unique_lock passing the mutex as parameter. When the unique_lock is constructed it will lock the mutex, and it gets destructed it will unlock the mutex. More importantly: If a exceptions is thrown, the std::unique_lock destructer will be called and so the mutex will be unlocked.
// 如果要锁定互斥锁，请创建一个类型 std::unique_lock 的局部变量，将互斥锁作为参数传递。当构造unique_lock时，它将锁定互斥锁，它被销毁后，它将解锁互斥锁。更重要的是：如果抛出异常，将调用 std::unique_lock 析构函数，因此互斥锁将被解锁。

// Example: 例：

// #include<mutex>
// int some_shared_var=0;

// int func() {
//     int a = 3;
//     { //Critical section
//         std::unique_lock<std::mutex> lock(my_mutex);
//         some_shared_var += a;
//     } //End of critical section
// }        

// Share
// Edit
// Follow
// edited Jun 24, 2018 at 14:52
// Saral Garg's user avatar
// Saral Garg
// 13277 bronze badges
// answered Feb 5, 2013 at 14:06
// André Puel's user avatar
// André Puel
// 8,66188 gold badges5151 silver badges8383 bronze badges
// 1
// Why don't you need a new keyword in order to create the lock?
// 为什么不需要 new 关键字来创建锁？
// – 
//  –
// Pro Q
//  Dec 6, 2018 at 1:39
// 16
// @ProQ It's created on the stack. If you used the new keyword, it wouldn't be destructed when you leave the scope.
// @ProQ 它是在堆栈上创建的。如果使用了 new 关键字，则在离开范围时不会销毁该关键字。
