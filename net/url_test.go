package net

import (
	"net/url"
	"testing"
)

//TestNetQueryUnescape -url解码
func TestNetQueryUnescape(t *testing.T){
	host := "http://www.baidu.com/log.gif?data={%22bid%22:100022,%22sid%22:101235}"
	str,_ := url.QueryUnescape(host)
	t.Logf("%v",str)
}

//TestQueryEscape-url编码
func TestQueryEscape(t *testing.T){
	host := `http://www.baidu.com/log.gif?data={"bid":100022,"sid":101235}`
	str:= url.QueryEscape(host)
	t.Logf("%v",str)
}
