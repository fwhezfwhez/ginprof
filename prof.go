package ginprof

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/pprof"
)

func ProfRouter(r gin.IRoutes) {

	// 主页面
	r.GET("/debug/pprof/", adapt(pprof.Index))

	r.GET("/debug/pprof/Profile", Profile())

	r.GET("/debug/pprof/heap", Heap())

	r.GET("/debug/pprof/block", Block())

	r.GET("/debug/pprof/goroutine", Goroutine())

	r.GET("/debug/pprof/allocs", Allocs())
	r.GET("/debug/pprof/cmdline", Cmdline())
	r.GET("/debug/pprof/threadcreate", Threadcreate())
	r.GET("/debug/pprof/mutex", Mutex())
	r.GET("/debug/pprof/trace", Trace())

}

func adapt(f func(w http.ResponseWriter, r *http.Request)) func(c *gin.Context) {
	return func(c *gin.Context) {
		f(c.Writer, c.Request)
	}
}

func Profile() gin.HandlerFunc {
	return adapt(pprof.Profile)
	//return adapt(pprof.Handler("Profile").ServeHTTP)
}

func Heap() gin.HandlerFunc {
	return adapt(pprof.Handler("heap").ServeHTTP)
}

func Block() gin.HandlerFunc {
	return adapt(pprof.Handler("block").ServeHTTP)
}

func Goroutine() gin.HandlerFunc {
	return adapt(pprof.Handler("goroutine").ServeHTTP)
}

func Allocs() gin.HandlerFunc {
	return adapt(pprof.Handler("allocs").ServeHTTP)
}

func Cmdline() gin.HandlerFunc {
	return adapt(pprof.Cmdline)
}

func Threadcreate() gin.HandlerFunc {
	return adapt(pprof.Handler("threadcreate").ServeHTTP)
}

func Mutex() gin.HandlerFunc {
	return adapt(pprof.Handler("mutex").ServeHTTP)
}

func Trace() gin.HandlerFunc {
	return adapt(pprof.Trace)
}
