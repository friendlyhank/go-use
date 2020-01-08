package test

import (
	"fmt"
	"testing"
)

//Error Log
func TestDivision(t *testing.T) {
	var b = 0
	if b == 0 {
		t.Error("B不能为零")
	} else {
		t.Log("B不为零")
	}
}

//test before and after
func TestQueryOrderList(t *testing.T){
	fmt.Println("TestQueryOrderList")
}

func TestMain(m *testing.M){
	setUp()
	m.Run()
	tearDown()
}


func setUp(){
	fmt.Println("connection db")
}

func tearDown(){
	fmt.Println("destory db connection")
}





