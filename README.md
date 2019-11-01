# goms  [![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/labstack/echo/master/LICENSE) [![Language](https://img.shields.io/badge/language-go-blue.svg)](https://golang.org/) [![Build Status](http://img.shields.io/travis/fuwensun/goms.svg?style=flat-square)](https://travis-ci.org/fuwensun/goms) [![Go Report Card](https://goreportcard.com/badge/github.com/fuwensun/goms)](https://goreportcard.com/report/github.com/fuwensun/goms)

## Introduction

目前微服务架构有两种: 一种以库的形式提供中间件的传统模式,如 Kratos,Micro;另一种以进程的形式提供中间件的sidecar模式,也叫service mesh,如 K8s,Istio.两者各有优势:
- 传统模式以函数调用的方式使用中间件;
- sidecar模式远程调用的方式使用中间件;

因此传统模式的性能优于sidecar模式,也因此sidecar模式的解耦性优于传统模式.


我们先通过手动的方式一步一步的构建一个微服务,感受的开发过程中的痛点.然后,再用不同的微服务架构开发一遍,发现这些痛点是怎样被架构解决,这样有助于我们理解架构为何这样设计,达到知其所以然的目的.


## Repositories
- [eModule](https://github.com/fuwensun/goms/tree/master/eModule)
- [eDocker](https://github.com/fuwensun/goms/tree/master/eDocker)
- [eYaml](https://github.com/fuwensun/goms/tree/master/eYaml)
- [eFlag](https://github.com/fuwensun/goms/tree/master/eFlag)
- [eHttp](https://github.com/fuwensun/goms/tree/master/eHttp)
- [eGrpc](https://github.com/fuwensun/goms/tree/master/eGrpc)
- [eConf](https://github.com/fuwensun/goms/tree/master/eConf)
- [eMysql](https://github.com/fuwensun/goms/tree/master/eMysql)
- eRedis
- eKratos
- eMicro
- eK8s
- eIstio

## Issue management

TODO

