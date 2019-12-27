package container

import (
	"container/list"
	"fmt"
	"testing"
)

func TestPushFront(t *testing.T){
	//创建一个新的链表
	l := list.New()
	l.PushFront(1)

	fmt.Println(l.Back())

	//遍历链表并打印它包含的元素
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func TestPushBack(t *testing.T){
	//创建一个新的链表
	l := list.New()
	l.PushBack(4)
	//遍历链表并打印它包含的元素
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
