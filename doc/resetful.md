
## 返回码

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

## api 风格
- fluent style 风格 api  
    ruby on rails 风格的 rest 路由映射  
    /people/{person_id}/grounps/{group_id}  
