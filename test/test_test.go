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

//300000	      4759 ns/op
func BenchmarkCreateOrder(b *testing.B){
	for i:=0;i<b.N;i++{
		fmt.Println("TestCreateOrder")
	}
}

//并发度去跑 200000	      6581 ns/op
func BenchmarkCreateOrderParellel(b *testing.B){
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next(){
			fmt.Println("TestCreateOrder")
		}
	})
}


