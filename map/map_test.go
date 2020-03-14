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
/*

/*
 *
 */
func TestMapBool(t *testing.T){
	mapScore :=map[string]int{
		"Tom":78,
		"Mary":60,
		"Kevin":90,
	}

	//
	score,ok := mapScore["Tom"]
	if ok{
		t.Logf("Name:Tom,Score:%d",score)
	}

	_,exist := mapScore["Tom"]
	if exist{
		t.Logf("这个值是存在滴!")
	}
}

/**
 *Map的遍历以及无序性
 */
func TestMapRange(t *testing.T){
	mapScore :=map[string]int{
		"Tom":78,
		"Mary":60,
		"Kevin":90,
	}

	for name,score := range mapScore{
		t.Logf("Name:%s,Score:%d",name,score)
	}
	//多次运行结果可以看出map的输出是无序
}

/*
 *delete删除
 */
func TestMapDelete(t *testing.T){
	mapScore := make(map[string]int)

	mapScore["Mary"] = 60
	t.Logf("Mary score:%d",mapScore["Mary"])

	delete(mapScore,"Mary")
	t.Logf("delete Mary:%d",mapScore["Mary"])
}

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




