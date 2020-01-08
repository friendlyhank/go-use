package test

import (
	"fmt"
	"testing"
)

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

