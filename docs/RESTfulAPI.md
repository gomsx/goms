# RESTful 风格接口

- 是什么
- 解决的问题
- 应用场景

### 是什么

RESTful 风格的接口叫作 RESTful API

- 结构清晰，符合标准，易于理解，扩展方便
- 面向资源类型的
- 是一种风格

误区

- 不要认为接口传递参数使用斜杠(/)分割而不用问号(?)传参
- 不要认为没有查询字符串就是 RESTful API
- 不要认为用了查询字符串就不是 RESTful API
- 不要认为用了 JSON 传输的 API 就是 RESTful API

### 解决的问题

- 互联网初期，页面请求和并发量也不高，对接口的要求没那么高，动态页面(jsp)就能满足需求。
- 随着互联网和移动设备的发展，传统的动态页面由于低效率而被 HTML+JavaScript(Ajax) 的前后端分离所取代，安卓、IOS、小程序等形式客户端的出现，客户端的种类出现多元化，
- 客户端和服务端就需要接口进行通信，接口的规范性就成了问题.

### 应用场景

RESTful 是目前最流行的接口设计规范，Github 的 API 设计就是很标准的 RESTful API。

## 一、REST介绍

### REST相关的知识

- REST (Representational State Transfer)，表现层状态转换
- 是一种软件架构风格、设计风格，不是标准，只提供了一组设计原则和约束条件
- 主要用于客户端和服务器交互类的软件
- 基于这个风格设计的软件可以更简洁，更有层次，更易于实现缓存等机制

2000 年 Roy Thomas Fielding 的博士论文中

- 定义表述性状态转移(Representational State Transfer，REST)的架构风格，
- 描述了如何使用 REST 来指导现代 Web 架构的设计和开发。
- 在符合架构原理前提下，理解和评估基于网络的应用软件的架构设计，得到一个功能强、性能好、适宜通信的架构。
需要注意的是REST并没有一个明确的标准，而更像是一种设计的风格，满足这种设计风格的程序或接口我们称之为 RESTful (从单词字面来看就是一个形容词)。所以 RESTful API 就是满足 REST 架构风格的接口。

### REST架构特征

>RESTful 是一种风格而不是标准

- **以资源为基础** 资源可以是一个图片、音乐、一个 XML 格式、HTML 格式或者JSON格式等网络上的一个实体，除了一些二进制的资源外普通的文本资源更多以 JSON 为载体、面向用户的一组数据(通常从数据库中查询而得到)。
- **统一接口** 对资源的操作包括获取、创建、修改和删除，正好对应 HTTP 协议提供的 GET、POST、PUT 和 DELETE 方法。从接口只能定位其资源，操作动作要从其 HTTP 请求方法类型得知。具体的 HTTP 方法和方法含义如下：
  - GET(SELECT)：从服务器取出资源(一项或多项)。
  - POST(CREATE)：在服务器新建一个资源。
  - PUT(UPDATE)：在服务器更新资源(客户端提供完整资源数据)。
  - PATCH(UPDATE)：在服务器更新资源(客户端提供需要修改的资源数据)。
  - DELETE(DELETE)：从服务器删除资源。
- **URI指向资源** URI = Universal Resource Identifier 统一资源标志符，用来标识抽象或物理资源的一个紧凑字符串。URI 包括 URL 和 URN，在这里更多时候可能代指URL(统一资源定位符)。RESTful 是面向资源的，每种资源可能由一个或多个URI对应，但一个URI只指向一种资源。
- **无状态** 服务器不能保存客户端的信息，每一次从客户端发送的请求中，要包含所有必须的状态信息，会话信息由客户端保存，服务器端根据这些状态信息来处理请求。当客户端可以切换到一个新状态的时候发送请求信息， 当一个或者多个请求被发送之后, 客户端就处于一个状态变迁过程中。每一个应用的状态描述可以被客户端用来初始化下一次的状态变迁。

### REST架构限制条件

>REST 架构的6个限制条件，也可称为 RESTful 6大原则

- **客户端-服务端(Client-Server)** 这个更专注客户端和服务端的分离，服务端独立可更好服务于前端、安卓、IOS 等客户端设备。
- **无状态(Stateless)** 服务端不保存客户端状态，客户端保存状态信息每次请求携带状态信息。
- **可缓存性(Cacheability)** 服务端需回复是否可以缓存以让客户端甄别是否缓存提高效率。
- **统一接口(Uniform Interface)** 通过一定原则设计接口降低耦合，简化系统架构，这是 RESTful 设计的基本出发点。当然这个内容除了上述特点提到部分具体内容比较多详细了解可以参考这篇 REST 论文内容。
- **分层系统(Layered System)** 客户端无法直接知道连接的到终端还是中间设备，分层允许你灵活的部署服务端项目。
- **按需代码(Code-On-Demand，可选)** 按需代码允许我们灵活的发送一些看似特殊的代码给客户端例如 JavaScript 代码。

## 二、RESTful API设计规范

设计一个 RESTful API
要从 URL 路径、HTTP 请求动词、状态码和返回结果等
其他的方面例如错误处理、过滤信息等规范

### URL设计规范

URL 为统一资源定位器,接口属于服务端资源,通过 URL 定位到资源才能去访问.
一个完整的 URL 由以下几个部分构成：  
URI = scheme "://" host  ":"  port "/" path [ "?" query ][ "#" fragment ]  

- scheme: 指底层用的协议，如 http、https、ftp
- host: 服务器的 IP 地址或者域名
- port: 端口，http 默认 80 端口
- path: 访问资源的路径，就是各种 web 框架中定义的 route 路由
- query: 查询字符串，为发送给服务器的参数，如数据`分页`、`排序`等
- fragment: 锚点，定位到页面的资源

RESTful 对 path 做了一些规范，

- 通常 RESTful API 的 path 组成如下：
/{version}/{resources}/{resource_id}

  - version：API 版本号，些版本号也可放置在头信息中，版本号有利于应用迭代
  - resources：资源，推荐用小写英文单词的复数形式
  - resource_id：资源的 id，访问或操作该资源
- 资源级别可能较大，其下还可细分很多子资源，那么 URL 的 path，如：
/{version}/{resources}/{resource_id}/{subresources}/{subresource_id}
- 有时增删改查无法满足业务要求，可以在 URL 末尾加上 action，表示是对资源的操作，如
/{version}/{resources}/{resource_id}/action

对 RESTful API 的 URL 其他规范如下：

- 使用英文且小写，不用大写字母
- 使用连字符用中杠"-"，而不用下杠"_"
- 使用 "/"表示层级关系,层级不要过深，靠前的层级相对稳定
- 结尾不要包含正斜杠分隔符"/"
- 不出现动词，用请求方式表示动作
- 资源表示用复数，不要用单数
- 不要使用文件扩展名

### HTTP动词

在 RESTful API 中，不同的 HTTP 请求方法有各自的含义

- GET /collection：从服务器查询资源的列表（数组）
- GET /collection/resource：从服务器查询单个资源
- POST /collection：在服务器创建新的资源
- PUT /collection/resource：更新服务器资源
- DELETE /collection/resource：删除服务器资源

在非 RESTful 风格的 API 中，我们通常使用 GET 请求和 POST 请求完成增删改查以及其他操作

- 查询和删除一般使用 GET 方式请求
- 更新和插入一般使用 POST 请求
- 从请求方式上无法知道 API 具体是干什么
- URL 上有操作的动词来表示 API 进行的动作，如 query，add，update，delete 等

四个 HTTP 请求方法的安全性和幂等性

- 安全性是指方法不会修改资源状态，即读的为安全的，写的操作为非安全的
- 幂等性的意思是操作多次效果相同，客户端重复调用也只返回同一个结果

| HTTP Method | 安全性 | 幂等性 | 解释 |
|:-|:-:|-:|-:|
| GET    | 安全   | 幂等   | 读安全，查询多次结果一致 |
| POST   | 非安全 | 非幂等 | 写非安全，插入多次结果不同 |
| PUT    | 非安全 | 幂等   | 写非安全，更新多次结果一致 |
| DELETE | 非安全 | 幂等   | 写非安全，删除多次结果一致 |

### 状态码和返回数据

服务器响应时，包含状态码和返回数据两个部分。

状态码

用各类状态码来表示请求的结果。状态码分为五大类：

- 1xx：相关信息
- 2xx：操作成功
- 3xx：重定向
- 4xx：客户端错误
- 5xx：服务器错误

每一大类有小类，而主要常用状态码：

- 200 OK - [GET]：用户请求的数据成功返回，操作是幂等的(Idempotent)
- 201 CREATED - [POST/PUT/PATCH]：用户新建或修改数据成功
- 202 Accepted - [*]：用户请求已经进入后台排队(异步任务)
- 204 NO CONTENT - [DELETE]：用户删除数据成功
- 400 INVALID REQUEST - [POST/PUT/PATCH]：用户发出的请求有错误，服务器没有新建或修改数据的操作，该操作是幂等的
- 401 Unauthorized - [*]：表示用户没有权限(令牌、用户名、密码错误)
- 403 Forbidden - [*] 表示用户得到授权(与401错误相对)，但是访问是被禁止的
- 404 NOT FOUND - [*]：用户发出的请求不存在的记录，服务器没有进行操作，操作是幂等的
- 406 Not Acceptable - [GET]：用户请求的格式不可得(比如请求JSON格式，但只有XML格式)
- 410 Gone -[GET]：用户请求的资源被永久删除，且不会再得到的
- 422 Unprocesable entity - [POST/PUT/PATCH] 当创建对象时，发生验证错误
- 500 INTERNAL SERVER ERROR - [*]：服务器发生错误，用户无法判断发出的请求是否成功

返回结果

返回JSON格式数据给客户端。

## 三、RESTful API 缺点

- RESTful API 根据 HTTP 的 GET、POST、PUT、DELETE 来区分操作资源的动作，而 HTTP Method 本身不可直接见，如果将动作放到 URL 的 path 上反而清晰可见
- 有些浏览器对 GET,POST 之外的请求支持不太友好，还需要特殊额外的处理
- 过分强调资源，而实际业务 API 需求可能比较复杂，只使用资源的增删改查不能满足所有的需求，强行使用 RESTful 风格 API 只会增加开发难度

如果使用场景和 RESTful 风格很匹配，那么可以采用 RESTful 风格 API。如果使用场景和 RESTful 风格不匹配，那也可以不用 RESTful 风格 API 或者可以借鉴,不能墨守成规。

