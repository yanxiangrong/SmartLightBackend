package main

import (
	_ "SmartLightBackend/control"
	"SmartLightBackend/models"
	"SmartLightBackend/network"
	"SmartLightBackend/pkg/logging"
	"SmartLightBackend/router"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	logging.Info("Server starting")
	gin.SetMode(gin.DebugMode) // Debug 模式

	r := router.InitRouter()
	server := &http.Server{
		Addr:    ":9527", // 端口9527
		Handler: r,
	}
	// 监听信号，实现软退出
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		logging.Info("Receive interrupt signal")
		if err := server.Close(); err != nil {
			logging.Fatal("Server Close:", err)
		}
	}()

	go network.ListenAndServe()

	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			logging.Info("Server closed under request")
		} else {
			logging.Fatal("Server closed unexpect")
		}
		network.Close()
		models.CloseDB()
	}
	logging.Info("Server exiting")
}
