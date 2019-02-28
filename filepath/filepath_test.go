package filepath

import (
	"fmt"
	"path/filepath"
	"testing"
)

//返回目录路径,不包含文件名
func TestFilePathDir(t *testing.T){
	dir := filepath.Dir("D:/logtest/xmiss-dailaimi.log")
	fmt.Println(dir)
}