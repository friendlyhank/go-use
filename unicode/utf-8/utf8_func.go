package utf_8

import (
	"fmt"
	"unicode/utf8"
)

func SingleExample() {
	str := "s你好世界😀"

	r, size := utf8.DecodeRuneInString(str)
	fmt.Printf("第一个字符: %c (Unicode: U+%04X)\n", r, r)
	fmt.Printf("第一个字符占用字节数: %d\n", size)
	fmt.Printf("原始字符串字节长度: %d\n", len(str))
	fmt.Printf("字符串字符数量: %d\n", utf8.RuneCountInString(str))
}

func iterateString() {
	fmt.Println("=== 遍历字符串中的所有字符 ===")

	str := "Go语言😊"
	fmt.Printf("遍历字符串: %s\n", str)

	// 方法1: 使用DecodeRuneInString手动遍历
	fmt.Println("方法1 - 使用DecodeRuneInString:")
	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("  位置 %d: %c (U+%04X) 占用 %d 字节\n", i, r, r, size)
		i += size
	}

	// 方法2: 使用range遍历（推荐）
	fmt.Println("方法2 - 使用range遍历:")
	for i, r := range str {
		fmt.Printf("  位置 %d: %c (U+%04X)\n", i, r, r)
	}
}

// 示例5: 处理混合字符串
func mixedStringExample() {
	fmt.Println("=== 示例5: 处理混合字符串 ===")

	str := "Hello世界🌍123"
	fmt.Printf("混合字符串: %s\n", str)

	position := 0
	charCount := 0

	for position < len(str) {
		r, size := utf8.DecodeRuneInString(str[position:])
		charCount++

		var charType string
		switch {
		case r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z':
			charType = "英文字母"
		case r >= '0' && r <= '9':
			charType = "数字"
		case r >= 0x4E00 && r <= 0x9FFF:
			charType = "中文汉字"
		case r >= 0x1F000:
			charType = "Emoji"
		default:
			charType = "其他"
		}

		fmt.Printf("  字符 %d: %c (%s) - 字节位置: %d, 占用: %d字节\n",
			charCount, r, charType, position, size)

		position += size
	}
	fmt.Println()
}

// 示例6: 错误处理和边界情况
func errorHandlingExample() {
	fmt.Println("=== 示例6: 错误处理和边界情况 ===")

	// 空字符串
	emptyStr := ""
	r, size := utf8.DecodeRuneInString(emptyStr)
	fmt.Printf("空字符串: r=%d (U+%04X), size=%d\n", r, r, size)

	// 无效的UTF-8序列
	invalidUTF8 := string([]byte{0xFF, 0xFE, 0xFD})
	r, size = utf8.DecodeRuneInString(invalidUTF8)
	fmt.Printf("无效UTF-8: r=%d (U+%04X), size=%d\n", r, r, size)
	fmt.Printf("是否为替换字符: %t\n", r == utf8.RuneError) // 注意这个很实用，可以用来判断是否乱码RuneError

	// 检查字符串是否为有效UTF-8
	testStrings := []string{
		"Hello",
		"你好",
		"😀",
		string([]byte{0xFF, 0xFE}),
	}

	for _, s := range testStrings {
		fmt.Printf("字符串 '%s' 是否为有效UTF-8: %t\n", s, utf8.ValidString(s))
	}
	fmt.Println()
}

// 实际应用 - 字符串截取
func stringTruncateExample() {
	fmt.Println("=== 示例7: 实际应用 - 字符串截取 ===")

	str := "这是一个很长的中文字符串，包含emoji😊和English"
	maxChars := 100

	fmt.Printf("原字符串: %s\n", str)
	fmt.Printf("截取前 %d 个字符:\n", maxChars)

	// 安全截取字符串，避免截断UTF-8字符
	truncated := safeStringTruncate(str, maxChars)
	fmt.Printf("截取结果: %s\n", truncated)
	fmt.Printf("截取后字符数: %d\n", utf8.RuneCountInString(truncated))
	fmt.Printf("截取后字节数: %d\n", len(truncated))
}

// 安全截取字符串的辅助函数
func safeStringTruncate(s string, maxChars int) string {
	if maxChars <= 0 {
		return ""
	}

	charCount := 0
	bytePos := 0

	for bytePos < len(s) {
		if charCount >= maxChars {
			break
		}

		_, size := utf8.DecodeRuneInString(s[bytePos:])
		bytePos += size
		charCount++
	}

	return s[:bytePos]
}
