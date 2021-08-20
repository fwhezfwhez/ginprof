## ginprof
ginprof包是面向以下场景开发的:

- 意图将golang net/http/pprof的pprof路由，接入业务项目里，使得pprof端口和业务端口保持一致，防止维护业务端口和pprof两套端口过于复杂。
- 意图将http/pprof里的func(w Writer, r *http.Request)转换为gin的func(c *gin.Context)接入gin的项目。

## 使用
```go
func main() {
    r:= gin.New()
    ginprof.ProfRouter(r)

    // 业务路由追加
    // xxx.Router(r)
    r.Run("8080")
}
```