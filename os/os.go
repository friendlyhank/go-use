package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	OsFlag()
}

//输入参数 例如:go run os.go 1 2
func OsFlag(){
	for idx, args := range os.Args {
		fmt.Println("参数" + strconv.Itoa(idx) + ":", args)
	}
}


