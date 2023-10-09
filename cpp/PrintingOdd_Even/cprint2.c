#include <unistd.h>
#include <stdlib.h>
#include <pthread.h>
#include <stdio.h>


int g_num = 0;
pthread_mutex_t mutex;
pthread_cond_t cond1,cond2;

void* thread1(void* arg)
{
	while(1)
	{
		pthread_mutex_lock(&mutex);
		pthread_cond_wait(&cond1,&mutex);
		printf("Thread1: %d \n",g_num);
		pthread_mutex_unlock(&mutex);
		sleep(1);
	}
	return NULL;
}

void* thread2(void* arg)
{
	while(1)
	{
		pthread_mutex_lock(&mutex);
		pthread_cond_wait(&cond2,&mutex);
		printf("Thread2: %d \n",g_num);
		pthread_mutex_unlock(&mutex);
		sleep(1);
	}
	return NULL;
}

void* thread3(void* arg)
{
	while(1)
	{
        //有可能出现线程3都运行了1次了，线程1还没开始，导致不是从1开始打印，为了避免这种情况，所以先让管理线程休眠一会。
        sleep(1);
		pthread_mutex_lock(&mutex);
		++g_num;
		pthread_mutex_unlock(&mutex);
		if((g_num % 2) == 0)
			pthread_cond_signal(&cond2);
		else if((g_num % 2) == 1){
			pthread_cond_signal(&cond1);}
	}
	return NULL;
}

int main()
{
	pthread_t p1,p2,p3;
	
	pthread_mutex_init(&mutex,NULL);
	pthread_cond_init(&cond1,NULL);
	pthread_cond_init(&cond2,NULL);	
	
	pthread_create(&p1,NULL,thread1,NULL);
	pthread_create(&p2,NULL,thread2,NULL);
	pthread_create(&p3,NULL,thread3,NULL);

	pthread_join(p1,NULL);
	pthread_join(p2,NULL);
	pthread_join(p3,NULL);

	pthread_mutex_destroy(&mutex);
	pthread_cond_destroy(&cond1);
	pthread_cond_destroy(&cond2);

	return 0;
}
