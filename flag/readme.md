

方法大全：

## 初始化命令解析:
NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet

## 设置指令的错误/用法信息输出目标
func (f *FlagSet) SetOutput(output io.Writer)

## 命令行标志解析器的错误/用法信息输出目标
func (f *FlagSet) Output() io.Writer

## var 自定义复杂的指令
flag.Var(&list, "list", "Comma-separated list of values")

## 解析命令参数
func Parse() 

## 遍历所有符合条件的命行行参数
func (f *FlagSet) VisitAll(fn func(*Flag)) 