package sync

import (
	"fmt"
	"testing"
	"time"
	"sync"
)


/**
1、读锁的时候无需等待读锁的结束
2、读锁的时候要等待写锁的结束
3、写锁的时候要等待读锁的结束
4、写锁的时候要等待写锁的结束

特点：读共享，写独占，写优先
 */
var m *sync.RWMutex

//读操作(读锁一般和写锁配合使用不然将没有意义)
func read(i int) {
	//读锁
	m.RLock()
	println(i, "read start")

	defer m.RUnlock()

	println(i, "reading")
	time.Sleep(1 * time.Second)
	println(i, "read end")
}

//write
func write(i int){
	//写锁
	m.Lock()
	println(i,"write start")

	defer m.Unlock()

	println(i,"writeing")
	time.Sleep(1*time.Second)
	println(i,"write end")
}

//读写互斥锁的运用
func TestRwMutexForR(t *testing.T){
	m = &sync.RWMutex{}

	//可以多个同时读
	go write(1)
	go write(2)
	go read(3)
	go write(4)
	go read(5)
	go read(6)
	go write(7)
	go read(8)
	go write(9)
	go read(10)
	time.Sleep(20 * time.Second)
}

//sync WaitGroup使用
func TestSyncWaitGroup(t *testing.T){
	w := &sync.WaitGroup{}

	go func(){
		defer w.Done()
		//w.Add新增如果，没有被消耗掉w.Done会一直进入等待状态w.Wait
		w.Add(1)
		fmt.Println("新增一个任务等待执行")
	}()

	w.Wait()
}
