package channel

import (
	"fmt"
	"testing"
	"time"
)

//通道属于引用类型
func TestChannelMap(t *testing.T){
	var mapChan = make(chan map[string]int, 1)

	syncChan := make(chan struct{}, 2)
	go func() { // 用于演示接收操作。
		//不多接收值需要for循环
		for {
			if elem, ok := <-mapChan; ok {
				elem["count"]++
			} else {
				break
			}
		}
		t.Logf("%v","Stopped. [receiver]")
		syncChan <- struct{}{}
	}()
	go func() { // 用于演示发送操作。
		countMap := make(map[string]int)
		for i := 0; i < 5; i++ {
			mapChan <- countMap
			time.Sleep(time.Millisecond)
			t.Logf("The count map: %v. [sender]\n", countMap)
		}
		close(mapChan)
		syncChan <- struct{}{}
	}()
	<-syncChan
	<-syncChan
}

//一个接收信号触发接收得例子
func TestChannelSyncMap(t *testing.T){
	var strChan = make(chan string, 3)

	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)
	go func() { // 用于演示接收操作。
		//没有for可以接收单个值，因为没有值是阻塞状态，可以作为触发装置控制
		<-syncChan1
		fmt.Println("Received a sync signal and wait a second... [receiver]")
		time.Sleep(time.Second)
		for {
			if elem, ok := <-strChan; ok {
				fmt.Println("Received:", elem, "[receiver]")
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		syncChan2 <- struct{}{}
	}()
	go func() { // 用于演示发送操作。
		for _, elem := range []string{"a", "b", "c", "d"} {
			strChan <- elem
			fmt.Println("Sent:", elem, "[sender]")
			if elem == "c" {
				syncChan1 <- struct{}{}
				fmt.Println("Sent a sync signal. [sender]")
			}
		}
		fmt.Println("Wait 2 seconds... [sender]")
		time.Sleep(time.Second * 2)
		close(strChan)
		syncChan2 <- struct{}{}
	}()
	<-syncChan2
	<-syncChan2
}