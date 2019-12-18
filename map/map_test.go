package _map

import (
	"testing"
)

//如何定义map
type MapGoods map[string]string

//TestQuoteMap- map是引用类型案例1
func TestQuoteMap(t *testing.T){
	mapOids := make(map[int64]bool,0)
	mapOids[10011] = true

	copyMapOids := mapOids
	copyMapOids[10011]=false

	t.Logf("%v",mapOids[10011])
}


type mapOids map[int64]bool
func (mo mapOids)FindMapOids(){
	mo[123]=false
}

//map是引用类型案例2
func TestMapOids(t *testing.T){
	mo := make(mapOids,0)
	mo[123]=true
	mo.FindMapOids()
	t.Logf("%v",mo)
}




