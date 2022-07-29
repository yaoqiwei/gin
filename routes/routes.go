package routes

import (
	"net/http"
	"time"

	"github.com/gin-contrib/pprof"

	"github.com/gin-gonic/gin"
)

// InitRouter ...
func InitRouter() {
	r := gin.New()
	apiGroupV1 := r.Group("/api/v1")
	Test(apiGroupV1)

	s := &http.Server{
		Addr:           ":9000",
		Handler:        r,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
	pprof.Register(r)
}
