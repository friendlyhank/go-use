package path

import (
	"path/filepath"
	"testing"
)

//返回目录路径,不包含文件名
func TestFilePathDir(t *testing.T){
	dir := filepath.Dir("../test.txt")

	t.Logf("%v",dir)
}

//返回的文件的扩展名
func TestFileExt(t *testing.T){
	exr := filepath.Ext("../test.txt")

	t.Logf("%v",exr)
}

//TestFileName- 获取文件名和路径中目录
func TestFileName(t *testing.T){
	dir := "E:\\data\\test.txt"
	_,file := filepath.Split(dir)
	t.Logf("%v",file)
}