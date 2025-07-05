package strconv

import (
	"fmt"
	"strconv"
	"testing"
)

// ParseBool - 将字符串解析成bool
func Test_ParseBool(t *testing.T) {
	v, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
}
