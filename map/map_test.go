package _map

import(
	"fmt"
	"testing"
)

/*
 *map是引用类型
 */

type mapOids map[int64]bool

func (mo mapOids)FindMapOids(){
	mo[123]=false
	fmt.Println(mo)
}

//Map用方法的表示
func TestMapOids(t *testing.T){
	mo := make(mapOids,0)
	mo[123]=true
	mo.FindMapOids()
	t.Logf("%v",mo)
}
