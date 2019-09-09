package _interface

import (
	"fmt"
	//"github.com/astaxie/beego/logs"
	//"os"
	"testing"
)

//type Server interface {
//	Init()
//	Run()
//	Stop()
//}
//
//type xmiss struct{
//}
//
//func (x *xmiss)Init(){
//	logs.Info("Initing")
//}
//
//func (x *xmiss)Run(){
//	logs.Info("Running")
//}
//
//func (x *xmiss)Stop(){
//	logs.Info("Stoping")
//}
//
//func DoServer(server Server){
//	server.Init()
//	server.Run()
//}
//
////1.TestInterface -接口实现
//func TestInterface(t *testing.T){
//	xs := &xmiss{}
//	//xs必须要实现server的所有方法才行Init(),Run,Stop
//	DoServer(xs)
//}
//
//
//type Flag struct{
//	Name string
//	DoValue Value
//}
//
//type Value interface {
//	String() string
//}
//
//type Getter interface {
//	Value
//	Get() interface{}
//}
//
//
//type boolValue bool
//
//func newBoolValue(val bool,p *bool)*boolValue{
//	*p = val
//	return (*boolValue)(p)
//}
//
//
//
//func (c *boolValue)String()string{
//	return "ss"
//}
//
//func (c *boolValue)Get()interface{}{
//	return bool(*c)
//}
//
//func NewFlag(name string,value Value)*Flag{
//	return &Flag{Name:name,DoValue:value}
//}
//
//
///**
//2.如果两个接口的方法相同，那么可以说这两个接口是相等的
//(1)interface不能写成函数，如func (v *Value)Get()string
//(2)boolValue实现了Value，Getter的方法，所以可以传值和强制转换
//	//设置值
//	flag := &Flag{Name:"version",DoValue:newBoolValue(false,new(bool))}
//	flag := NewFlag("version",newBoolValue(false,new(bool)))
//	//强制转换.(interface) .(struct)
//	flag.DoValue.(Getter)
//(3)p bool 类型的要转换成 type boolValue bool
//	(*boolValue)(p)
//(4)(c *boolValue)Get()interface{} 的实现要仔细研究newBoolValue传true或false
// */
//
//func TestInterfaceGet(t *testing.T){
//	flag := &Flag{Name:"version",DoValue:newBoolValue(true,new(bool))}
//	//flag := NewFlag("version",newBoolValue(false,new(bool)))
//	if flag.DoValue.(Getter).Get().(bool){
//		logs.Info("%v","没进来这里吗")
//		os.Exit(0)
//	}
//}

//3
//(1)
type KvClient interface{
	Put(id int64)int64
}

type kvClient struct{
	KvClient
	Name string
}

type key struct{
}

func NewKvClient()KvClient{
	return &key{}
}

func (k *key)Put(id int64)int64{
	fmt.Println(id)
	return id
}

//TestMoreInterface -
func TestMoreInterface1(t *testing.T){
	kv :=&kvClient{NewKvClient(),"kv"}

	id := kv.Put(100)

	t.Logf("Result：%v",id)
}

//(2)这种方式只会输出一个,KvClient2 Put
type KvClient2 interface{
	Put(id int64)int64
}

type kvClient2 struct{
	KvClient2
	Name string
}

func (kv *kvClient2)Put(id int64)int64{
	//加10做测试
	id += 10
	fmt.Println(id)
	return id
}

type key2 struct{
}

func NewKvClient2()KvClient2{
	return &key2{}
}

func (k *key2)Put(id int64)int64{
	fmt.Println(id)
	return id
}

//TestMoreInterface -
func TestMoreInterface2(t *testing.T){
	kv :=&kvClient2{NewKvClient2(),"kv"}

	id := kv.Put(100)

	t.Logf("Result：%v",id)
}

//(3)这种方式只会输出一个Put
type KvClient3 interface{
	Put(id int64)int64
}

type kvClient3 struct{
	KvClient3
	Name string
}

func (kv *kvClient3)Put(id int64)int64{
	//加10做测试
	fmt.Println(id)
	return kv.KvClient3.Put(id)
}

type key3 struct{
}

func NewKvClient3()KvClient3{
	return &key3{}
}

func (k *key3)Put(id int64)int64{
	fmt.Println(id)
	return id
}

//TestMoreInterface -
func TestMoreInterface3(t *testing.T){
	kv :=&kvClient3{NewKvClient3(),"kv"}

	id := kv.Put(100)

	t.Logf("Result：%v",id)
}



