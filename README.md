# goms  

[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/labstack/echo/master/LICENSE) [![Language](https://img.shields.io/badge/language-go-blue.svg)](https://golang.org/) [![Build Status](http://img.shields.io/travis/fuwensun/goms.svg?style=flat-square)](https://travis-ci.org/fuwensun/goms) [![Go Report Card](https://goreportcard.com/badge/github.com/fuwensun/goms)](https://goreportcard.com/report/github.com/fuwensun/goms) [![make all](https://github.com/fuwensun/goms/workflows/make_all/badge.svg)](https://github.com/fuwensun/goms/actions?query=workflow%3Amake_all)

## Introduction

本库的目的是学习如何用 go 语言开发微服务.

我们先通过手动的方式一步一步的构建一个微服务,感受的开发过程中的痛点.然后,再用不同的微服务架构开发一遍,发现这些痛点是怎样被架构解决,这样有助于我们理解架构为何这样设计,达到知其所以然的目的.

目前微服务架构有两种:
- 一种以库的形式提供中间件的 sdk 模式,如 [Kratos][15] 和 [Micro][16];
- 一种以进程的形式提供中间件的 sidecar 模式,也叫 service mesh,如 [K8s][17] 和 [Istio][18].

两者各有优势:
- sdk 模式以函数调用的方式使用中间件;
- sidecar 模式以进程间通讯的方式使用中间件;

因此:
- sdk 模式的性能优于 sidecar 模式;
- sidecar 模式的解耦性优于 sdk 模式;

以上介绍的是微服务的实现,微服务的另一个重要方面是设计,这部分的内容参考领域驱动设计(DDD).

## Repositories

- [eDocker][22].  主题应用容器化, docker 是应用容器引擎,打包应用及依赖到一个可移植的容器.

- [eModule][21].  主题依赖管理, go module 是 go 标准的依赖管理工具.

- [eYaml][23].  主题配置/编排文件格式. yaml 一种常用的数据序列化格式,k8s/istio 等编排文件的格式.

- [eFlag][24].  主题命令行参数处理, 使用准库 flag 包.

- [eHttp][25].  主题 http 服务, 使用 gin 框架实现.

- [eGrpc][26].  主题 rpc 服务, 使用 grpc 包.

- [eConf][27].  主题服务的配置, 一个简单包 conf.

- [eMysql][28].  主题关系型数据库, 最常用的关系型数据库 MySQL.

- [eRedis][29].  主题缓存, 最常用缓存数据库 Redis. 

- [eTest][30].  主题测试, 覆盖各种类型的测试. 

- eKratos.  主题微服务框架, 使用 Kratos 开发一个微服务.

- eMicro.  主题微服务框架, 使用 Micro 开发一个微服务.

- [eK8s][33].  主题微服务部署, 使用 K8s 部署一个微服务.

- eIstio.  主题微服务部署, 使用 Istio 部署一个微服务.

## Issue management

欢迎提交 bugs 和 feature 报告.



[15]:https://github.com/bilibili/kratos
[16]:https://github.com/micro/micro
[17]:https://github.com/kubernetes/kubernetes
[18]:https://github.com/istio/istio

[21]:https://github.com/fuwensun/goms/tree/master/eModule
[22]:https://github.com/fuwensun/goms/tree/master/eDocker
[23]:https://github.com/fuwensun/goms/tree/master/eYaml
[24]:https://github.com/fuwensun/goms/tree/master/eFlag
[25]:https://github.com/fuwensun/goms/tree/master/eHttp
[26]:https://github.com/fuwensun/goms/tree/master/eGrpc
[27]:https://github.com/fuwensun/goms/tree/master/eConf
[28]:https://github.com/fuwensun/goms/tree/master/eMysql
[29]:https://github.com/fuwensun/goms/tree/master/eRedis
[30]:https://github.com/fuwensun/goms/tree/master/eTest
[31]:https://github.com/fuwensun/goms/tree/master
[32]:https://github.com/fuwensun/goms/tree/master
[33]:https://github.com/fuwensun/goms/tree/master/eK8s
[34]:https://github.com/fuwensun/goms/tree/master

