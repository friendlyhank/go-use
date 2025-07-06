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

// 返回字符串字面量和转义字符
func Test_Quote(t *testing.T) {
	// 返回一个用双引号括起的 Go 字符串字面量，并返回对应go的转义序列(\t, \n, \xFF, \u0100)
	fmt.Println(strconv.Quote(`Hello
World`))
	fmt.Println("Hello\nWorld")
}
