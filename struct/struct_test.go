package _struct

import "testing"

//struct保存格式
type Student struct{
	ID int64
	Name string
	Age int64
}

var student2 = &struct {
	ID int64
	Name string
	Age int64
}{
	10023,
	"张三",
	18,
}

func TestMakeStudent(t *testing.T){
	student1 := &Student{
		ID:10023,
		Name:"张三",
		Age:18,
	}

	t.Logf("%v",student1)
	t.Logf("%v",student2)
}