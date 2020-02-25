package main

import (
	"fmt"
	"log"
	"net/http"
	"routes"

	"github.com/gin-gonic/gin"
	"github.com/luankefei/golang-canvas/src/config"
	"github.com/luankefei/golang-canvas/src/libs"

	"github.com/tdewolff/canvas"
)

var fontFamily *canvas.FontFamily

func init() {
	config.Setup()
	libs.Setup()
}

func main() {
	init()

	libs.LoadFont("../static/HanYiLeMiao_Regular.ttf", "LeMiao", canvas.FontRegular)

	serve()
}

func serve() {
	gin.SetMode(config.ServerSetting.RunMode)

	// middlewares.Wechat()

	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	routersInit := routes.Init(r)
	readTimeout := config.ServerSetting.ReadTimeout
	writeTimeout := config.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", 8000)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()

	// If you want Graceful Restart, you need a Unix system and download github.com/fvbock/endless
	//endless.DefaultReadTimeOut = readTimeout
	//endless.DefaultWriteTimeOut = writeTimeout
	//endless.DefaultMaxHeaderBytes = maxHeaderBytes
	//server := endless.NewServer(endPoint, routersInit)
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err: %v", err)
	//}
}
