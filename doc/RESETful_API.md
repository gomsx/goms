
## HTTP response status codes

1xx - informational  
2xx - success  
3xx - redirection  
4xx - client error  
5xx - server error  

200 ok  
300 multiple choices  
301 moved permanently  
302 found  
304 not modified  
307 temporary redirect  
400 bad request  
401 unauthorized  
403 forbidden  
404 not found  
410 gone  
500 internal server error  
501 not implemented  
503 service unavailable  
550 permission denied  

200 表示删除请求被成功执行，返回被删除的资源  
202 表示删除请求被接受，但还没有被执行  
204 表示删除请求被执行，但没有返回被删除的资源  


## HTTP request methods

### POST

- 新建成功  
    - 有返回 200 Ok  
    - 无返回 201 Created

>https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Status/200
https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Status/201

### PUT  

- 更新成功  
    - 有返回 200 (通常不用)  
    - 无返回 204 No Content  
- 新建成功   
    - 有返回 200 (通常不用)  
    - 无返回 201 Created  

>https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Status/200
https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Status/201
https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Status/204

### GET

- 读取成功
    - 有返回 200 Ok

>https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Status/200

### DELETE

- 删除成功  
    - 无返回 204 No Content  

>https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Status/200
https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Status/204


总结:  
post - 200 201  
put - 204 201  
get - 200  
delete - 204  

## go 状态码

go\src\net\http\status.go
```
package http

// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
const (
	StatusContinue           = 100 // RFC 7231, 6.2.1
	StatusSwitchingProtocols = 101 // RFC 7231, 6.2.2
	StatusProcessing         = 102 // RFC 2518, 10.1
	StatusEarlyHints         = 103 // RFC 8297

	StatusOK                   = 200 // RFC 7231, 6.3.1
	StatusCreated              = 201 // RFC 7231, 6.3.2
	StatusAccepted             = 202 // RFC 7231, 6.3.3
	StatusNonAuthoritativeInfo = 203 // RFC 7231, 6.3.4
	StatusNoContent            = 204 // RFC 7231, 6.3.5
	StatusResetContent         = 205 // RFC 7231, 6.3.6
	StatusPartialContent       = 206 // RFC 7233, 4.1
	StatusMultiStatus          = 207 // RFC 4918, 11.1
	StatusAlreadyReported      = 208 // RFC 5842, 7.1
	StatusIMUsed               = 226 // RFC 3229, 10.4.1

	StatusMultipleChoices   = 300 // RFC 7231, 6.4.1
	StatusMovedPermanently  = 301 // RFC 7231, 6.4.2
	StatusFound             = 302 // RFC 7231, 6.4.3
	StatusSeeOther          = 303 // RFC 7231, 6.4.4
	StatusNotModified       = 304 // RFC 7232, 4.1
	StatusUseProxy          = 305 // RFC 7231, 6.4.5
	_                       = 306 // RFC 7231, 6.4.6 (Unused)
	StatusTemporaryRedirect = 307 // RFC 7231, 6.4.7
	StatusPermanentRedirect = 308 // RFC 7538, 3

	StatusBadRequest                   = 400 // RFC 7231, 6.5.1
	StatusUnauthorized                 = 401 // RFC 7235, 3.1
	StatusPaymentRequired              = 402 // RFC 7231, 6.5.2
	StatusForbidden                    = 403 // RFC 7231, 6.5.3
	StatusNotFound                     = 404 // RFC 7231, 6.5.4
	StatusMethodNotAllowed             = 405 // RFC 7231, 6.5.5
	StatusNotAcceptable                = 406 // RFC 7231, 6.5.6
	StatusProxyAuthRequired            = 407 // RFC 7235, 3.2
	StatusRequestTimeout               = 408 // RFC 7231, 6.5.7
	StatusConflict                     = 409 // RFC 7231, 6.5.8
	StatusGone                         = 410 // RFC 7231, 6.5.9
	StatusLengthRequired               = 411 // RFC 7231, 6.5.10
	StatusPreconditionFailed           = 412 // RFC 7232, 4.2
	StatusRequestEntityTooLarge        = 413 // RFC 7231, 6.5.11
	StatusRequestURITooLong            = 414 // RFC 7231, 6.5.12
	StatusUnsupportedMediaType         = 415 // RFC 7231, 6.5.13
	StatusRequestedRangeNotSatisfiable = 416 // RFC 7233, 4.4
	StatusExpectationFailed            = 417 // RFC 7231, 6.5.14
	StatusTeapot                       = 418 // RFC 7168, 2.3.3
	StatusMisdirectedRequest           = 421 // RFC 7540, 9.1.2
	StatusUnprocessableEntity          = 422 // RFC 4918, 11.2
	StatusLocked                       = 423 // RFC 4918, 11.3
	StatusFailedDependency             = 424 // RFC 4918, 11.4
	StatusTooEarly                     = 425 // RFC 8470, 5.2.
	StatusUpgradeRequired              = 426 // RFC 7231, 6.5.15
	StatusPreconditionRequired         = 428 // RFC 6585, 3
	StatusTooManyRequests              = 429 // RFC 6585, 4
	StatusRequestHeaderFieldsTooLarge  = 431 // RFC 6585, 5
	StatusUnavailableForLegalReasons   = 451 // RFC 7725, 3

	StatusInternalServerError           = 500 // RFC 7231, 6.6.1
	StatusNotImplemented                = 501 // RFC 7231, 6.6.2
	StatusBadGateway                    = 502 // RFC 7231, 6.6.3
	StatusServiceUnavailable            = 503 // RFC 7231, 6.6.4
	StatusGatewayTimeout                = 504 // RFC 7231, 6.6.5
	StatusHTTPVersionNotSupported       = 505 // RFC 7231, 6.6.6
	StatusVariantAlsoNegotiates         = 506 // RFC 2295, 8.1
	StatusInsufficientStorage           = 507 // RFC 4918, 11.5
	StatusLoopDetected                  = 508 // RFC 5842, 7.2
	StatusNotExtended                   = 510 // RFC 2774, 7
	StatusNetworkAuthenticationRequired = 511 // RFC 6585, 6
)
```

<!-- return StatusCode(HttpStatusCode.NoContent);//成功
return NotFound(); //失败
HTTP状态码400 Bad Request请求PUT失败，在响应正文中使用自然语言文本（例如英文）解释PUT失败的原因。(RFC 2616 Section 10.4) -->

## api 风格
- fluent style 风格 api  
    ruby on rails 风格的 rest 路由映射  
    /people/{person_id}/grounps/{group_id}  