# eFlag

这个模块加入命令行参数的处理,使用标准库 flag 包.

## 运行
```
cd eFlag/cmd

go run .  -s=string -i=1 -bool=true

go run .  -s string --i 1 -bool=true

go run .  -s=string -i 1 --bool true

go run .  -s=string -i=1 -bool=true xx yy zz

go run .  -s string -i 1 -bool=true xx yy zz
```

## 其他更好的库

