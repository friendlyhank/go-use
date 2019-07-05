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

var strChan =make(chan string,3)

func Receive(strChan <- chan string){
	//go func(){
	//	for{
	//		if relem,ok :=<- strChan ;ok{
	//			fmt.Println("Receice:",relem,"Receicer")
	//		}
	//	}
	//}()

	go func(){
		//注意是range strChan,而不是range <- strChan
		//range 也会阻塞
		for relem :=range strChan{
			fmt.Println("Receice:",relem,"Receicer")
		}
	}()
}

func Send(strChan chan <- string){
	go func(){
		for _,elem := range []string{"a","b","c","d"}{
			strChan <- elem
			fmt.Println("Sent:",elem,"Sender")
		}
		close(strChan)
	}()
}

//单通道接收和发送得例子
func TestReceiveAndSend(t *testing.T){
	Send(strChan)
	Receive(strChan)
	time.Sleep(2 * time.Second)
}

//单向通道  <- chan int 接收
//单向通道 chan <-  int 发送
//单接收通道和单发送通道和双通道不能互相转化
//双通道可以作为函数传入到单向通道(如上个例子)
func TestChannelConv(t *testing.T){
	var ok bool
	ch := make(chan int, 1)
	_, ok = interface{}(ch).(<-chan int)
	fmt.Println("chan int => <-chan int:", ok)
	_, ok = interface{}(ch).(chan<- int)
	fmt.Println("chan int => chan<- int:", ok)

	sch := make(chan<- int, 1)
	_, ok = interface{}(sch).(chan int)
	fmt.Println("chan<- int => chan int:", ok)

	rch := make(<-chan int, 1)
	_, ok = interface{}(rch).(chan int)
	fmt.Println("<-chan int => chan int:", ok)
}