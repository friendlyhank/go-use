package _map

import (
	"testing"
)

type MapGoods map[string]string
/*
*map定义和初始化

#定义map m
var m map[string]int  //nil

#初始化方式1 make
m = make(map[string]int)

#初始化方式2 花括号
m = map[string]int{}

#也可以直接赋值
m = map[string]int{
		"Tom":50,
		"Mary":60,
	}
 */

//TestQuoteMap- map是引用类型案例1
func TestQuoteMap(t *testing.T){
	mapOids := make(map[int64]bool,0)
	mapOids[10011] = true

	copyMapOids := mapOids
	copyMapOids[10011]=false

	t.Logf("%v",mapOids[10011])
}


type mapOids map[int64]bool
func (mo mapOids)FindMapOids(){
	mo[123]=false
}

//map是引用类型案例2
func TestMapOids(t *testing.T){
	mo := make(mapOids,0)
	mo[123]=true
	mo.FindMapOids()
	t.Logf("%v",mo)
}




