#include <unistd.h>
#include <stdlib.h>
#include <pthread.h>
#include <stdio.h>

int g_num = 1;
pthread_mutex_t mutex;
pthread_cond_t cond1,cond2;

void* thread1(void* arg)
{
	while(1)
	{
		pthread_mutex_lock(&mutex);
        //如果需要交替打印一定范围(例如1-10)内的数字，那么可以加上下面两行代码
        // exit()
        // exit并不算一个线程结束函数，但调用它可以关闭所有文件，终止正在执行的进程，进而导致当前进程中所有线程结束。一般exit(0)表示正常退出，exit(x)都表示异常退出，这个x是返回给操作系统以供其他程序使用。在任何一个线程中调用exit函数都会导致进程结束。进程一旦结束，那么进程中的所有线程都将结束。

        // pthread_exit()
        // 只会终止当前线程，不会影响进程中其它线程的执行，因此常用于终止子线程，主线程中调用也只会终止主线程。

        // void pthread_exit(void *retval);
        if(g_num > 20)
        {
           printf("Thread1: pthread_exit \n");
           pthread_exit(NULL);

           //printf("Thread1: return \n");
           //return NULL;
        }

        printf("Thread1: %d \n",g_num);
        g_num ++;
        pthread_cond_signal(&cond2);
        pthread_cond_wait(&cond1,&mutex);		
		pthread_mutex_unlock(&mutex);
		//sleep(1);
	}
	return NULL;
}

void* thread2(void* arg)
{
	while(1)
	{
        //这个sleep(1)加在前面是因为开启线程时有可能是线程2先打印，
        //导致变成thread2输出奇数，threa1输出偶数。为了避免这种情况，可以在延迟下线程2的打印。
        //sleep(1);
		pthread_mutex_lock(&mutex);
        printf("Thread2: %d \n",g_num);
		g_num++;
        pthread_cond_signal(&cond1);
        pthread_cond_wait(&cond2,&mutex);
		pthread_mutex_unlock(&mutex);
		
	}
	return NULL;
}
int main()
{
	pthread_t p1,p2;
	
	pthread_mutex_init(&mutex,NULL);
	pthread_cond_init(&cond1,NULL);
	pthread_cond_init(&cond2,NULL);	
	
	pthread_create(&p1,NULL,thread1,NULL);
	pthread_create(&p2,NULL,thread2,NULL);
    pthread_join(p1,NULL);
	pthread_join(p2,NULL);
	
	pthread_mutex_destroy(&mutex);
	pthread_cond_destroy(&cond1);
	pthread_cond_destroy(&cond2);

	return 0;
}
