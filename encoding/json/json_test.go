package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"
)

//编解码时忽略指定某个字段
func TestJsonIgnore(t *testing.T){
	type UserInfo struct{
		Name string `json:"name"`
		Age int `json:"age,string"`
		Address string `json:"-"`
	}

	user := &UserInfo{Name: "123456",Age:18,Address: "广东省广州市天河区"}
	b,_ := json.Marshal(user)
	fmt.Println(string(b))
}

//编解码如果字段为空则忽略该字段
func TestJsonOmitempty(t *testing.T){
	type UserInfo struct{
		Name string `json:"name"`
		Age int `json:"age,string"`
		Address string `json:"address,omitempty"`
	}
	user := &UserInfo{Name: "123456",Age:18}
	b,_ := json.Marshal(user)
	fmt.Println(string(b))
}

//json.Number
func TestJsonNumberForInt(t *testing.T){
	type UserInfo struct{
		Name string `json:"name"`
		Age json.Number `json:"age"`
		Address string `json:"address"`
	}

	str := `{"name":"张三","age":18,"address":"广东省广州市天河区"}`
	user := &UserInfo{}
	err := json.Unmarshal([]byte(str),user)
	if err != nil{
		panic(err)
	}
	fmt.Printf("%+v",user)

	str1 := `{"name":"张三","age":"18","address":"广东省广州市天河区"}`
	user1 := &UserInfo{}
	err = json.Unmarshal([]byte(str1),user1)
	if err != nil{
		print(err)
	}
	fmt.Printf("%+v",user1)
}

//用于interface类型编解码过程将数值解析成float64类型，如果数值较大会精度丢失的问题
func TestJsonNumberForFloat64(t *testing.T){
	s := `{"a":6673221165400540161}`

	d := make(map[string]interface{})
	err := json.Unmarshal([]byte(s), &d)
	if err != nil {
		panic(err)
	}

	s2, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(s2))

	//改进方法
	decoder := json.NewDecoder(strings.NewReader(s))
	decoder.UseNumber()
	err = decoder.Decode(&d)
	if err != nil{
		panic(err)
	}

	s2,err = json.Marshal(d)
	if err != nil{
		panic(err)
	}
	fmt.Println(string(s2))
}

type doTime time.Time
func (t *doTime)UnmarshalJSON(b []byte) error {
	b = bytes.Trim(b, "\"")
	dt,err  := time.Parse("2006-01-02 15:04:05",string(b))
	if err != nil{
		panic(err)
	}
	*t = doTime(dt)
	return nil
}

func (t *doTime) MarshalJSON() ([]byte, error) {
	str :=  fmt.Sprintf("\"%s\"",time.Time(*t).Format("2006-01-02 15:04:05"))
	return []byte(str),nil
}

type UserInfo struct{
	Name string `json:"name"`
	Age int `json:"age,string"`
	Address string `json:"address"`
	LoginTime doTime `json:"logintime"`
}

/*
   type UserInfo struct{
		Name string `json:"name"`
		Age int `json:"age,string"`
		Address string `json:"address"`
		LoginTime time.Time `json:"logintime"`
	}
	user := &UserInfo{Name: "123456",Age:18,LoginTime: time.Now()}
	b,_ := json.Marshal(user)
	fmt.Println(string(b))

	user1 := &UserInfo{}
	json.Unmarshal([]byte(b),user1)
//有时候输出的并不是我们想要的数据
{"name":"123456","age":"18","address":"","logintime":"2021-01-11T15:08:32.7838158+08:00"}
 */
func TestImplementsMarshaler(t *testing.T){
	user := &UserInfo{Name: "123456",Age:18,LoginTime: doTime(time.Now())}
	b,err := json.Marshal(user)
	if err != nil{
		t.Errorf("%v",err)
	}
	fmt.Println(string(b))

	user1 := &UserInfo{}
	json.Unmarshal([]byte(b),user1)
	fmt.Println(time.Time(user1.LoginTime))
}

func (t *doTime)UnmarshalText(b []byte) error {
	dt,err  := time.Parse("2006-01-02 15:04:05",string(b))
	if err != nil{
		panic(err)
	}
	*t = doTime(dt)
	return nil
}

func (t *doTime) MarshalText() ([]byte, error) {
	str := time.Time(*t).Format("2006-01-02 15:04:05")
	return []byte(str),nil
}
func TestImplementsTextMarshaler(t *testing.T){
	user := &UserInfo{Name: "123456",Age:18,LoginTime: doTime(time.Now())}
	b,err := json.Marshal(user)
	if err != nil{
		t.Errorf("%v",err)
	}
	fmt.Println(string(b))

	user1 := &UserInfo{}
	json.Unmarshal([]byte(b),user1)
	fmt.Println(time.Time(user1.LoginTime))
}

