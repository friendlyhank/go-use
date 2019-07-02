package _go

import (
	"fmt"
	"testing"
	"time"
)

//Go传参案例
func TestForGo(t *testing.T){
		names := []string{"Eric","Harry","Robert","Jim","Mark"}

		/**
			for语句执行完了才执行go,go函数得执行时机是没办法保证的
		 */

		//for _,name := range names{
		//	go func(){
		//		fmt.Printf("Hello,%s !\n",name)
		//	}()
		//}
		//time.Sleep(time.Microsecond)

		/**
		结果:
			Hello,Mark !
			Hello,Mark !
			Hello,Mark !
			Hello,Mark !
			Hello,Mark !
		 */

		/**
		go 函数也可以参数声明，把他作为副本就不会被轻易得改变了
		 */
		for _,name := range names{
			go func(who string){
				fmt.Printf("Hello,%s !\n",who)
			}(name)
		}
		time.Sleep(time.Microsecond)
}


