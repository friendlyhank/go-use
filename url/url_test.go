package main

import (
	"github.com/astaxie/beego/logs"
	"net/url"
	"testing"
)


//TestURLParse -解析URL
func TestURLParse(t *testing.T){
	//解析url
	apurl,_ := url.Parse("http://localhost:2380")
	logs.Info("%v",apurl.Host)
}

//TestURLToString -将URL转成stirng
func TestURLToString(t *testing.T){
	apurl,_ := url.Parse("http://localhost:2380")
	logs.Info("%v",apurl.String())
}
