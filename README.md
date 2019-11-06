# goms  

[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/labstack/echo/master/LICENSE) [![Language](https://img.shields.io/badge/language-go-blue.svg)](https://golang.org/) [![Build Status](http://img.shields.io/travis/fuwensun/goms.svg?style=flat-square)](https://travis-ci.org/fuwensun/goms) [![Go Report Card](https://goreportcard.com/badge/github.com/fuwensun/goms)](https://goreportcard.com/report/github.com/fuwensun/goms)

## Introduction

本库的目的是学习如何用 go 语言构建微服务.

我们先通过手动的方式一步一步的构建一个微服务,感受的开发过程中的痛点.然后,再用不同的微服务架构开发一遍,发现这些痛点是怎样被架构解决,这样有助于我们理解架构为何这样设计,达到知其所以然的目的.

目前微服务架构有两种: 一种以库的形式提供中间件的传统模式,如 Kratos,Micro;另一种以进程的形式提供中间件的sidecar模式,也叫service mesh,如 K8s,Istio.两者各有优势:
- 传统模式以函数调用的方式使用中间件;
- sidecar模式进程间通讯的方式使用中间件;

因此:
- 传统模式的性能优于sidecar模式;
- sidecar模式的解耦性优于传统模式;

## Repositories

- [eModule][21]
- [eDocker][22]
- [eYaml][23]
- [eFlag][24]
- [eHttp][25]
- [eGrpc][26]
- [eConf][27]
- [eMysql][28]
- [eRedis][29] TODO
- [eKratos][30] TODO
- [eMicro][31] TODO
- [eK8s][32] TODO
- [eIstio][33] TODO

[21]:https://github.com/fuwensun/goms/tree/master/eModule
[22]:https://github.com/fuwensun/goms/tree/master/eDocker
[23]:https://github.com/fuwensun/goms/tree/master/eYaml
[24]:https://github.com/fuwensun/goms/tree/master/eFlag
[25]:https://github.com/fuwensun/goms/tree/master/eHttp
[26]:https://github.com/fuwensun/goms/tree/master/eGrpc
[27]:https://github.com/fuwensun/goms/tree/master/eConf
[28]:https://github.com/fuwensun/goms/tree/master/eMysql
[29]:https://github.com/fuwensun/goms/tree/master
[30]:https://github.com/fuwensun/goms/tree/master
[31]:https://github.com/fuwensun/goms/tree/master
[32]:https://github.com/fuwensun/goms/tree/master
[33]:https://github.com/fuwensun/goms/tree/master

## Issue management

TODO