import threading
lock = threading.Lock()
owner = "A"
i = 0

def test1():
    global owner, i
    while i <= 5:
        lock.acquire()  # Here
        if owner == "A" and i<=5:
            print(owner + ":" + str(i))
            owner = "B"
            i += 1
        lock.release()

def test2():
    global owner, i
    while i <= 5:
        lock.acquire()  # Here
        if owner == "B" and i<=5:
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
