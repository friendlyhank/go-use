package utf_8

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

// DecodeRuneInString 输出字符串的第一个字符和大小
func Test_DecodeRuneInString(t *testing.T) {
	// DecodeRuneInString 返回两个值：
	// 1. r rune: 解码出的Unicode字符
	// 2. size int: 该字符占用的字节数

	// 简单案例
	SingleExample()

	// 便利字符串
	iterateString()

	// 判断混合字符串类型
	mixedStringExample()

	//错误处理和边界情况
	errorHandlingExample()

	// 安全的安字符数截取字符串
	stringTruncateExample()

}

// Test_RuneCountInString - 获取字符串字符数量
func Test_RuneCountInString(t *testing.T) {
	str := "s你好世界😀"
	fmt.Printf("字符串字符数量: %d\n", utf8.RuneCountInString(str))
}
