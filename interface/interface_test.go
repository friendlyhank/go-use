package _interface

import (
	"github.com/astaxie/beego/logs"
	"os"
	"testing"
)

type Server interface {
	Init()
	Run()
	Stop()
}

type xmiss struct{
}

func (x *xmiss)Init(){
	logs.Info("Initing")
}

func (x *xmiss)Run(){
	logs.Info("Running")
}

func (x *xmiss)Stop(){
	logs.Info("Stoping")
}

func DoServer(server Server){
	server.Init()
	server.Run()

}

//1.TestInterface -接口实现
func TestInterface(t *testing.T){
	xs := &xmiss{}
	//xs必须要实现server的所有方法才行Init(),Run,Stop
	DoServer(xs)
}


type Flag struct{
	Name string
	DoValue Value
}

type Value interface {
	String() string
}

type Getter interface {
	Value
	Get() interface{}
}


type boolValue bool

func newBoolValue(val bool,p *bool)*boolValue{
	*p = val
	return (*boolValue)(p)
}



func (c *boolValue)String()string{
	return "ss"
}

func (c *boolValue)Get()interface{}{
	return bool(*c)
}

func NewFlag(name string,value Value)*Flag{
	return &Flag{Name:name,DoValue:value}
}


/**
2.如果两个接口的方法相同，那么可以说这两个接口是相等的
(1)interface不能写成函数，如func (v *Value)Get()string
(2)boolValue实现了Value，Getter的方法，所以可以传值和强制转换
	//设置值
	flag := &Flag{Name:"version",DoValue:newBoolValue(false,new(bool))}
	flag := NewFlag("version",newBoolValue(false,new(bool)))
	//强制转换.(interface) .(struct)
	flag.DoValue.(Getter)
(3)p bool 类型的要转换成 type boolValue bool
	(*boolValue)(p)
(4)(c *boolValue)Get()interface{} 的实现要仔细研究newBoolValue传true或false
 */

func TestInterfaceGet(t *testing.T){
	flag := &Flag{Name:"version",DoValue:newBoolValue(true,new(bool))}
	//flag := NewFlag("version",newBoolValue(false,new(bool)))
	if flag.DoValue.(Getter).Get().(bool){
		logs.Info("%v","没进来这里吗")
		os.Exit(0)
	}
}

