package unicode

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

// TestRuneError - 遍历输出非法字符
func TestRuneError(t *testing.T) {
	s := "x1f�"
	for width := 0; len(s) > 0; s = s[width:] {
		r := rune(s[0])
		width = 1
		// 表示多字节utf-8,小于的表示ASCII（单字节）
		if r >= utf8.RuneSelf {
			r, width = utf8.DecodeRuneInString(s)
		}
		if r == utf8.RuneError {
			fmt.Println(s[:width])
		}
	}
}
