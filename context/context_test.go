package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func doStuff(ctx context.Context){
	for{
		time.Sleep(1 *time.Second)
		select{
		case <- ctx.Done():
			fmt.Println("cancel")
		default:
			fmt.Println("working")
		}
	}
}

//context 超时取消
func TestContextWithTimeOut(t *testing.T){
	ctx,cancel := context.WithTimeout(context.TODO(),3 *time.Second)

	go doStuff(ctx)

	time.Sleep(10 *time.Second)
	cancel()
}
