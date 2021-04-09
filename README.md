# pub6
pub6源码，示例网站：https://pub6.top

代码基于gfast框架，https://github.com/tiger1103/gfast 项目为前后端分离框架

0、项目数据库文件 ：
```
\pub6\data\db.sql 
\pub6\data\pub6.sql
```

创建数据库并导入；

1、配置文件为：\pub6\config\config.toml
当前server配置信息如下，端口为8200, 服务器使用的资源文件根路径为./public/resource

```
[server]
    Address          = ":8200"
    ServerRoot       = "./public/resource"
    AccessLogEnabled = true
    ErrorLogEnabled  = true
    PProfEnabled     = true
    LogPath          = "./data/log/server_log"
    SessionIdName    = "sysSessionId"
    SessionPath      = "./data/session"
    SessionMaxAge    = "24h"
    DumpRouterMap    = true
    NameToUriType = 3

```

其中jwt配置


```
[gToken]

CacheMode = 2 此处若使用了redis配置为2 若没使用redis配置1
CacheKey = "GToken:"
Timeout = 0
MaxRefresh = 0
TokenDelimiter="_"
EncryptKey = "koi29a83idakguqjq29asd9asd8a7jhq"
AuthFailMsg = "登录超时，请重新登录"
MultiLogin = true
##运行 go run main.go 直接访问http://localhost:8200  现在直接访问的是./public/resource/index.html
如果要访问gfast的管理后台，可以使用http://localhost:8200/admin.html

```


2、前端使用的layui的fly版本，路径在：\pub6\public\resource， layui的模块代码主要在\pub6\public\resource\res\mods\ ；

3、该代码仅用于学习，如有其它用途责任请自担！

