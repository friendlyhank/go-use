package reflect

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"reflect"
	"testing"
)


//v.Kind的运用
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

//TestReflect -
func TestReflect(t *testing.T){
	maps := map[string]string{
		"a":"1111111",
		"b":"2222222",
	}

	structs := &struct {
		a string
		b string
	}{
		a:"1111111",
		b:"2222222",
	}

	sa := "11111111"
	ia := 151515115


	mapv := reflect.ValueOf(maps)
	structv := reflect.ValueOf(structs)
	sav := reflect.ValueOf(sa)
	iav := reflect.ValueOf(ia)

	//(1)ValueOf用来获取输入参数接口中的数据的值，如果接口为空则返回0
	fmt.Println("Print  =====valueof=====")
	logs.Info("%v",mapv)
	logs.Info("%v",structv)
	logs.Info("%v",sav)
	logs.Info("%v",iav)

	//(2)TypeOf用来动态获取输入参数接口中的值的类型，如果接口为空则返回nil
	mapt := reflect.TypeOf(maps)
	structt := reflect.TypeOf(structs)
	sat := reflect.TypeOf(sa)
	iat := reflect.TypeOf(ia)

	fmt.Println("Print  =====typeof=====")
	logs.Info("%v",mapt)
	logs.Info("%v",structt)
	logs.Info("%v",sat)
	logs.Info("%v",iat)

	//(3)
	fmt.Println("Print  =====valueof Kind=====")
	logs.Info("%v",mapv.Kind())
	logs.Info("%v",structv.Kind())
	logs.Info("%v",sav.Kind())
	logs.Info("%v",iav.Kind())

	//(4)
	fmt.Println("Print  =====typeof Kind=====")
	logs.Info("%v",mapt.Kind())
	logs.Info("%v",structt.Kind())
	logs.Info("%v",sat.Kind())
	logs.Info("%v",iat.Kind())

	//(5)
	//valueof isnil
	//如mapv.IsNil

	//(6)
	fmt.Println("Print  =====valueof Elem=====")
	//logs.Info("%v",mapv.Elem())
	logs.Info("%v",structv.Elem())
	//logs.Info("%v",sav.Elem())   //panic
	//logs.Info("%v",iav.Elem())	//panic

	//(7)
	fmt.Println("Print  =====typeof Elem=====")
	logs.Info("%v",mapt.Elem())
	logs.Info("%v",structt.Elem())
	//logs.Info("%v",sat.Elem())  //panic
	//logs.Info("%v",iat.Elem())  //panic

	//已知类型强制转换
	fmt.Println("Print  =====valueof interface=====")
	//转换的时候，如果转换的类型不完全符合，则直接panic，类型要求非常严格！
	//转换的时候，要区分是指针还是指
	convenmap := mapv.Interface().(map[string]string)
	logs.Info("%v",convenmap)

	fmt.Println("Print  =====valueof NumField=====")
	//(8)=====typeof interface===== NumField()
	//logs.Info("%v",sat.NumField()) painc 这样获取元素是错误的

	//logs.Info("%v",mapv.Elem().NumField())  //panic map获取不了下级元素
	//hank-import
	logs.Info("%+v",mapv.MapKeys())

	logs.Info("%v",structt.Elem().NumField())
	//logs.Info("%v",sat.Elem().NumField())  //panic string没有下级元素，会异常
	//logs.Info("%v",iat.Elem().NumField())  //panic int没有下级元素


	fmt.Println("Print  =====valueof Filed=====")
	//(9)key和value获取
	//field := getType.Field(i)
	//value := getValue.Field(i).Interface()
	logs.Info("%v",structt.Elem().Field(0))


	fmt.Println("Print  =====valueof MapIndex=====")
	//hank-import如果是map
	//mapv.MapIndex(key Value)
	logs.Info("%v",mapv.MapIndex(mapv.MapKeys()[0]).String())


	//(10)方法，方法个数
	// getType.NumMethod()
	//getType.Method(i)

	//(11)方法的调用
	// 一定要指定参数为正确的方法名
	// 2. 先看看带有参数的调用方法
	//methodValue := getValue.MethodByName("ReflectCallFuncHasArgs")
	//args := []reflect.Value{reflect.ValueOf("wudebao"), reflect.ValueOf(30)}
	//methodValue.Call(args)

	// 一定要指定参数为正确的方法名
	// 3. 再看看无参数的调用方法
	//methodValue = getValue.MethodByName("ReflectCallFuncNoArgs")
	//args = make([]reflect.Value, 0)
	//methodValue.Call(args)

	//(12)
	logs.Info("%v",structt.Elem().NumField())
	//logs.Info("%v",structt.NumField()) //painc 应该是下级元素的数量
}

func TestReflectAAA(t *testing.T) {
	//maps := map[string]string{
	//	"a": "1111111",
	//	"b": "2222222",
	//}

	structs := &struct {
		a string
		b string
	}{
		a:"1111111",
		b:"2222222",
	}

	//mapv := reflect.ValueOf(maps)
	//structv := reflect.ValueOf(structs)
	//
	//mapt := reflect.TypeOf(maps)
	structt := reflect.TypeOf(structs)

	logs.Info("%v",structt.Elem().NumField())
	logs.Info("%v",structt.NumField())




}