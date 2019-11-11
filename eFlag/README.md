# eFlag

命令行参数的处理,使用标准库 flag 包.

## 运行
```
cd eFlag/cmd

go run .  -s=string -i=1 -bool=true

go run .  -s string --i 1 -bool=true

go run .  -s=string -i 1 --bool true

go run .  -s=string -i=1 -bool=true xx yy zz

go run .  -s string -i 1 -bool=true xx yy zz

go run .  -s string -i 1 -bool=true - xx yy zz

go run .  -s string -i 1 -bool=true -- xx yy zz

```

## 标准库 flag 包

### 命令行参数:

类型:
- 标志型flag, 如: cmd -flag=arg
- 非标型non-flag,如: cmd arg

传递:
- 标志型命令参数和变量绑定.
- 解析命令参数.
- 标志型通过绑定的变量来获得,非标志型通过 flag.Args() 获得.

标志性命令参数和变量绑定:

```go
//方式一:
    var flagvar = flag.Int("flagname", "default value", "help message for flagname")
```

```go
//方式二:
    var flagvar string
    func init() {
        flag.IntVar(&flagvar, "flagname", "default value", "help message for flagname")
    }
```
解析:
```
func main() {
    flag.Parse()
    //...
}
```

使用:

```go
xxx := flagvar //标志型
yyy := flag.Args() //非标志型
```

## 命令格式:

标志型参数的使用格式:

- -flag
- -flag=x
- -flag x  // 非bool型标志参数

为什么 `-flag x` 格式不适用于bool型标志参数?

- cmd -flagbool file1 file2 file3 , 表示 cmd 命令带flagbool标志操作 file1,file2,file3 文件.
- cmd -flagbool false  file2 file3, 有歧义:一 表示 cmd 命令带flagbool标志操作 false,file2,file3 三文件;二 表示 cmd 命令带flagbool false 标志操作 file2,file3 两文件.

命令行参数的使用格式:

- cmd -flag=farg nfarg
最常用格式,解析到第一个不带"-"的参数前.
- cmd -flag=farg - nfarg
标志型参数解析到"-"之前,"-"属于第一个非标志型参数,被解析到 flag.arg(0) 中.
- cmd -flag=farg -- nfarg
标志型参数解析到"--"之后,"--"是标志型参数终止符.


## 其他更好的库

