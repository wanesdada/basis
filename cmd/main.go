package cmd

import (
	redis "basis/cache"
	"basis/config"
	"basis/db"
	"basis/httpserver"
	"basis/util"
	"context"
	"fmt"
	_ "github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	config.InitConfig("/../config/config.yml")
	err := redis.InitRedis(config.GetRedis())
	if err != nil {
		log.Fatalf("redis init err %v", err)
		return
	}
	initDB, err := db.InitDB(config.GetDb())
	if err != nil {
		log.Fatalf("db init err %v", err)
		return
	}
	defer func() {
		err := initDB.Close()
		if err != nil {
			log.Fatalf("db close err %v", err)
		}
	}()

	router := httpserver.SetupRouter()

	server := &http.Server{
		Addr:         config.GetService().Port,
		Handler:      router,
		ReadTimeout:  config.GetService().AppReadTimeout * time.Second,
		WriteTimeout: config.GetService().AppWriteTimeout * time.Second,
	}

	util.Logger.Info("|-----------------------------------|")
	util.Logger.Info("|            qp-web-server          |")
	util.Logger.Info("|-----------------------------------|")
	util.Logger.Info("|  Go Http Server Start Successful  |")
	util.Logger.Info("|    HttpPort" + config.GetService().Port + "  Pid:" + fmt.Sprintf("%d", os.Getpid()) + "       |")
	util.Logger.Info("|-----------------------------------|")

	log.Println("|-----------------------------------|")
	log.Println("|            qp-web-server          |")
	log.Println("|-----------------------------------|")
	log.Println("|  Go Http Server Start Successful  |")
	log.Println("|    HttpPort" + config.GetService().Port + "  Pid:" + fmt.Sprintf("%d", os.Getpid()) + "       |")
	log.Println("|    TcpPort" + config.GetService().TCPPort + "   Pid:" + fmt.Sprintf("%d", os.Getpid()) + "       |")
	log.Println("|-----------------------------------|")

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	sig := <-signalChan
	log.Println("Get Signal:", sig)
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
