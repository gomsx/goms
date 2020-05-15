# goms  

[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/labstack/echo/master/LICENSE) [![Language](https://img.shields.io/badge/language-go-blue.svg)](https://golang.org/) [![Build Status](http://img.shields.io/travis/fuwensun/goms.svg?style=flat-square)](https://travis-ci.org/fuwensun/goms) [![Go Report Card](https://goreportcard.com/badge/github.com/fuwensun/goms)](https://goreportcard.com/report/github.com/fuwensun/goms) [![make_v1.2.x](https://github.com/fuwensun/goms/workflows/make_v1.2.x/badge.svg)](https://github.com/fuwensun/goms/actions?query=workflow%3Amake_v1.2.x)



## Introduction

本库的目的是学习如何用 go 语言开发微服务.

我们先通过手动的方式一步一步地构建一个微服务,感受开发过程中的痛点.然后,再用微服务框架开发,发现这些痛点是怎样被框架解决,这样有助于我们理解框架为何这样设计,达到知其所以然的目的.

目前微服务框架有两种:
- 一种以库的形式提供中间件的 sdk 模式,如 [Kratos][15] 和 [Micro][16];
- 一种以进程的形式提供中间件的 sidecar 模式,也叫 service mesh,如 [K8s][17] 和 [Istio][18].

两者各有优势:
- sdk 模式以函数调用的方式使用中间件;
- sidecar 模式以进程间通讯的方式使用中间件;

因此:
- sdk 模式的性能优于 sidecar 模式;
- sidecar 模式的解耦性优于 sdk 模式;

以上介绍的是微服务的实现,微服务的另一个重要方面是设计,这部分内容参考领域驱动设计(DDD).

## Repositories

- [eDocker][22].  主题应用容器化, docker 是广泛使用的打包应用和依赖的容器.

- [eModule][21].  主题依赖管理, go module 是 go 标准的依赖管理工具.

- [eYaml][23].  主题数据序列化, yaml 格式数据序列化,常用作配置/编排文件.

- [eFlag][24].  主题命令行参数处理, 使用准库 flag 包.

- [eHttp][25].  主题 http 服务, 使用 Gin 框架实现.

- [eGrpc][26].  主题 rpc 服务, 使用 grpc 包.

- [eConf][27].  主题服务的配置, 一个简单包 conf.

- [eMysql][28].  主题关系型数据库, 常用的 MySQL.

- [eRedis][29].  主题缓存数据库, 常用的 Redis.

- [eTest][30].  主题测试, 覆盖各种类型的测试.

- [eK8s][31].  主题微服务部署, 使用 K8s 部署一组微服务.

    - [namespace][311].
    - [configmap][312].
    - [volume][313].
    - [deployment][314].
    - [service][315].

- eIstio.  主题微服务部署, 使用 Istio 部署一组微服务.

- eKratos.  主题微服务框架, 使用 Kratos 开发一个微服务.

- eMicro.  主题微服务框架, 使用 Micro 开发一个微服务.

## Issue management

欢迎提交 bugs 和 feature 报告.

[15]:https://github.com/bilibili/kratos
[16]:https://github.com/micro/micro
[17]:https://github.com/kubernetes/kubernetes
[18]:https://github.com/istio/istio

[21]:https://github.com/fuwensun/goms/tree/release-v1.2.0/eModule
[22]:https://github.com/fuwensun/goms/tree/release-v1.2.0/eDocker
[23]:https://github.com/fuwensun/goms/tree/release-v1.2.0/eYaml
[24]:https://github.com/fuwensun/goms/tree/release-v1.2.0/eFlag
[25]:https://github.com/fuwensun/goms/tree/release-v1.2.0/eHttp
[26]:https://github.com/fuwensun/goms/tree/release-v1.2.0/eGrpc
[27]:https://github.com/fuwensun/goms/tree/release-v1.2.0/eConf
[28]:https://github.com/fuwensun/goms/tree/release-v1.2.0/eMysql
[29]:https://github.com/fuwensun/goms/tree/release-v1.2.0/eRedis
[30]:https://github.com/fuwensun/goms/tree/release-v1.2.0/eTest
[31]:https://github.com/fuwensun/goms/tree/release-v1.2.0/eK8s


[311]:https://github.com/fuwensun/goms/tree/release-v1.2.0/eK8s/namespace
[312]:https://github.com/fuwensun/goms/tree/release-v1.2.0/eK8s/configmap
[313]:https://github.com/fuwensun/goms/tree/release-v1.2.0/eK8s/volume
[314]:https://github.com/fuwensun/goms/tree/release-v1.2.0/eK8s/deployment
[315]:https://github.com/fuwensun/goms/tree/release-v1.2.0/eK8s/service