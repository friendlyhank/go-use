package slice

import (
	"testing"
)

type Student struct{
	Name string
	ID int64
}

//Slice复制切片(浅拷贝)
func TestSliceCopy(t *testing.T){
	enport :=[]*Student{&Student{Name:"张三",ID:1}}
	eps := make([]*Student, len(enport))
	copy(eps,enport)

	for _,student := range eps{
		student.Name = "小李"
	}

	for _,stu := range enport{
		t.Logf("%+v",stu)
	}
}

func DeleteSlice(a []int,num int) []int{
	for i := 0; i < len(a); i++ {
		if a[i] == num {
			a = append(a[:i], a[i+1:]...)
		}
	}
	return a
}
//删除指定的元素
func TestSliceDel(t *testing.T){
	eps := DeleteSlice([]int{2379,2380},2380)
	t.Logf("%v",eps)
}
