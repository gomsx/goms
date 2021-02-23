# Bash

## 变量 vxx

定义
```
vxx="string"
vxx=("hello" "word" "haha")
```

引用
```
${vxx}
"$vxx"
```

## 函数 fxx

定义
```
function fxx{}
fxx{}
```

传参，参数的获取
```
fxx(){
    ver="$1"
    xxx="$2"
}

fxx "argx"
```

返回
```
fxx(){return}
fxx(){return 255}
fxx(){echo "xxx"}
```

调用，返回值的获取
```
fxx "argx"
rv=$? 

rv="$(fxx "argx")"
```

## 控制语句

```
if true ;then 
    echo "ok"
fi

for var in vars;do
    echo "ok"
done
```

## 文件的调用

source filex.sh
. filex.sh

bash filex.sh
sh filex.sh

exec filex.sh
