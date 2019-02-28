package log

import (
	"testing"
	"errors"
	"log"
)

//打开文件
//fmt.Sprint一样，多了os.Exit(1)错误退出
func TestLogFatal(t *testing.T){
	err := errors.New("错误了")
	log.Fatal(err)
}

