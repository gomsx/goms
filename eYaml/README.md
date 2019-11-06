# eYaml

这个模块加入数据的序列化和反序列,使用 YAML 格式和 github.com/go-yaml/yaml 包.

> YAML: YAML Ain't Markup Language
> What It Is: YAML is a human friendly data serialization
  standard for all programming languages.

## 数据序列化

- 序列化是指把结构化的数据用特定文本格式表示.
- 反序列化是序列化的逆过程.
- 数据序列化后可以用于传输和保存.

结构化数据和序列化数据之间操作序列化/反序列化,不同的语言有各自的名称:
- go: marshalling/unmarshalling 
- python: pickling/unpickling
- Java: serialization/serialization

如果一个文件里存着格式化的数据,那么结构化的数据和这个文件间的操作一般称为:
dump 和 load

## 其他序列化格式

- TOML
- JSON
- CSON
- XML