package path

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

// 返回目录路径,不包含文件名
func TestFilePathDir(t *testing.T) {
	dir := filepath.Dir("../test.txt")
	t.Logf("%v", dir)
}

// TestFilePathJoin - 拼接文件名到目录路径中，生成新的完整路径
func TestFilePathJoin(t *testing.T) {
	executablePath, err := os.Executable()
	if err != nil {
		fmt.Println(executablePath)
		return
	}
	dirPath := filepath.Dir(executablePath)
	newFilePath := filepath.Join(dirPath, "test.txt")
	fmt.Println(newFilePath)
}

// 返回的文件的扩展名
func TestFileExt(t *testing.T) {
	exr := filepath.Ext("../test.txt")

	t.Logf("%v", exr)
}

// TestFileName- 获取文件名和路径中目录
func TestFileName(t *testing.T) {
	dir := "E:\\data\\test.txt"
	_, file := filepath.Split(dir)
	t.Logf("%v", file)
}

// 获取文件的绝对路径
func TestFilepathAbs(t *testing.T) {
	abs, err := filepath.Abs("../test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(abs)
}

// Test_FilePathBase - 获取路径的文件名部分
func Test_FilePathBase(t *testing.T) {
	path := filepath.Base("/Users/hank/go/src/github.com/friendlyhank/go-use/path/filepath_test.go")
	fmt.Println(path)
}
