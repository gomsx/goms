# sed

## Linux 文本处理三剑客，指的是 grep、awk 和 sed。掌握常见用法，就能大大提高开发效率

## sed

1. sed 大写改成小写

```
sed -i 's/[A-Z]/\l&/g'
```

2. sed 搜索复杂的字符串

```
//复杂的字符串 ("Int64(\"user_id\", user.Uid).")，
//先用 grep 测试是否可以搜索到(检测特殊字符串是否转义)，
//再用 sed 操作
grep -r "Int64(\"user_id\", user.Uid)." 
sed -i 's/Int64(\"user_id\", user.Uid).//g' $(grep -rl "user_id")
```

3. 搜索行，包含字符串 "Msgf(.* = %v"

```
sed '/Msgf(.* = %v/p' goms/internal/server/http/user.go
```

4. 搜索行，包含字符串Msgf，替换" = "为": " 

```
sed -i "/Msgf/ s/ = /: /g"  $(grep -rl "Msgf" --exclude-dir=.git)
```
