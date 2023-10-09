import threading
lock = threading.Lock()
owner = "A"
i = 0

def test1():
    global owner, i
    while i <= 5:
        lock.acquire()
        if owner == "A":
            print(owner + ":" + str(i))
            owner = "B"
            i += 1
        lock.release()

def test2():
    global owner, i
    while i <= 5:
        lock.acquire()
        if owner == "B":
            print(owner + ":" + str(i))
            owner = "A"
            i += 1
        lock.release()

A = threading.Thread(target=test1)
B = threading.Thread(target=test2)

A.start()
B.start()

A.join()
B.join()



# The reason you are seeing the issue, is that the thread is still within the while loop when i moves from 5 to 6. The thread seems to obtain lock before the value is updated. i.e. The speed of release and lock is quicker than the assignment of i. So the thread goes through one more round, but the only check you have is if the owner is equal to A or B. So you have to once again have the check in place to test that i<=5.
# 您看到此问题的原因是，当 i 从 5 移动到 6 时，线程仍在 while 循环中。线程似乎在更新值之前获取了锁。即释放和锁定的速度比分配 3 快。因此，线程再经过一轮，但您唯一的检查是所有者是否等于 A 或 B 。因此，您必须再次进行检查才能测试该 i<=5 。


# If you had a single thread, then you would not have this issue, but because you are using two threads, you have to think of them as different entities, doing their own jobs. While i is being updated in one thread, the other thread is still spinning and checking, locking and unlocking.
# 如果你有一个线程，那么你就不会有这个问题，但是因为你使用两个线程，你必须将它们视为不同的实体，做自己的工作。当 i 在一个线程中更新时，另一个线程仍在旋转并检查、锁定和解锁。


# The best way to visualise the thread actions is to have print statements within the lock (before the if-statement) - to see what each thread is up to.
# 可视化线程操作的最佳方法是在锁中（在 if 语句之前）有 print 个语句 - 以查看每个线程在做什么。