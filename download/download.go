package main

import (
	"deepgo/download/pkg/setting"
	"deepgo/download/routers"
	"fmt"
	_ "github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := routers.InitRouter()
	router.Static("static", "./assets")

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
