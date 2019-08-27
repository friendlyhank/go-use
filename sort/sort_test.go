package sort

import (
	"fmt"
	"sort"
	"testing"
)

//SortString
func TestSortString(t *testing.T){
	var s= []string{"b","d","g","a","h"}

	sort.Strings(s)

	fmt.Println(s)
}
