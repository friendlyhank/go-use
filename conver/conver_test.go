package conver

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

//TODO HANK
type pb_PutResponse struct{
	Key string
}

type PutResponse pb_PutResponse

func TestConverStruct(t *testing.T){
	pbPutResponse := &pb_PutResponse{Key:"foo"}

	//注意分先后顺序
	putResponse := (*PutResponse)(pbPutResponse)

	t.Logf("%v",putResponse)
}

//时间转化
func TestTimeDuration(t *testing.T){
	time.Sleep(time.Duration(rand.Intn(15)) * time.Second)
	fmt.Println("时间转化成功")
}




