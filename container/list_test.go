package test

import (
	"container/list"
	"fmt"
	"testing"
)

//TestFrontAndBack- 获取头部和尾部元素
func TestFrontAndBack(t *testing.T){
	l := list.New()
	l.PushFront(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	fmt.Println(l.Front().Value)
	fmt.Println(l.Back().Value)
}

//TestPushFront-双链表头部插入
func TestPushFront(t *testing.T){
	l := list.New()
	_ = l.PushFront(1)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

//TestPushBack-双链表尾部插入
func TestPushBack(t *testing.T){
	l := list.New()
	_ = l.PushBack(4)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

//TestInsertBefore-在某个元素之前插入
func TestInsertBefore(t *testing.T){
	l := list.New()
	e4 := l.PushBack(4)
	l.InsertBefore(3,e4)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

//TestInsertAfter-在某个元素之后插入
func TestInsertAfter(t *testing.T){
	l := list.New()
	e1 := l.PushFront(1)
	l.InsertAfter(2,e1)
	for e := l.Front();e != nil;e = e.Next(){
		fmt.Println(e.Value)
	}
}


