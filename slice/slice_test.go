package slice

import (
	"testing"
)

//Slice复制切片
func TestSliceCopy(t *testing.T){
	enport :=[]string{"2379","2380"}
	eps := make([]string, len(enport))
	copy(eps,enport)

	t.Logf("%v",eps)
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
