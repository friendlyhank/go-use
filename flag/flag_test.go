package main

import (
	"flag"
	"os"
	"testing"
)

// TestNewFlagSet - 初始化新的指令解析
func TestNewFlagSet(t *testing.T) {
	NewFlagSet()
}

// TestVar - 自定义的复杂的命令规则
func TestVar(t *testing.T) {
	os.Args = []string{"myprogram", "--list", "a,b,c"}
	FlagVar()

	flag.PrintDefaults()
}
