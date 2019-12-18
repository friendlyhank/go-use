package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

/**
Invalid Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Ptr
	Slice
	String
	Struct
	UnsafePointer
 */
//TestReflectKind- 通过kind处理不同的分支
func TestReflectKind(t *testing.T){
	for _,v := range []interface{}{"hi",42,func(){}}{
		//shitch v.Kind
		switch v := reflect.ValueOf(v);v.Kind() {
		case reflect.String:
				fmt.Println(v.String())
		case reflect.Int,reflect.Int8,reflect.Int16,reflect.Int32,reflect.Int64:
			fmt.Println(v.Int())
		default:
			fmt.Printf("unhandled kind %s",v.Kind())
		}
	}
}

type HelloDemo struct{}
func (h *HelloDemo)Do(){
	fmt.Println("hello")
}
//TestMethodCallNoParam-无参数调用函数
func TestMethodCallNoParam(t *testing.T){
	hd := &HelloDemo{}
	value := reflect.ValueOf(hd).MethodByName("Do")
	fmt.Println(value)
}

type SayHello struct{}
func (s *SayHello)Do(a int,b string){
	fmt.Println("hello"+b,a)
}

//TestMethodCallWithParam- 带参数的函数调用
func TestMethodCallWithParam(t *testing.T){

	a := reflect.ValueOf(33)
	b := reflect.ValueOf("word")

	in := []reflect.Value{a,b}

	sh := &SayHello{}
	 reflect.ValueOf(sh).MethodByName("Do").Call(in)
}

type TagTest struct {
	A int    `json:"aaa" test:"testaaa"`
	B string `json:"bbb" test:"testbbb"`
}

//TestTagSetAndGet -struct tag 解析
func TestTagGet(t *testing.T){
	ts := TagTest{
		A: 123,
		B: "hello",
	}
	tt := reflect.TypeOf(ts)
	for i := 0; i < tt.NumField(); i++ {
		field := tt.Field(i)
		fmt.Println("field",field)
		if json, ok := field.Tag.Lookup("json"); ok {
			fmt.Println(json)
		}
		//设置用Set
		test := field.Tag.Get("test")
		fmt.Println(test)
	}
}

//TODO 深度复制