# dockerfile


## 目标
- 构建快(层数少，利用缓存)
- 体积小(层数少，文件少)
- 可读性好

## 方法
- 编写 .dockerignore 文件
- 单个容器只运行单个应用
- 指定基础镜像的 TAG
- 使用 COPY 优于 ADD
- 使用 WORKDIR 而不是 cd 命令
- 使用绝对路径
- 明确需要暴露的接口
- 多个 RUN 合并
- RUN 之后可以删除多余的文件
- 使用 exec 方式启动应用程序
- 使用 LABEL 设置元数据
- 使用 HEALTHCHECK