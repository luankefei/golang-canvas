package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/luankefei/golang-canvas/routes"

	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/gin-gonic/gin"

	"image/color"
	"image/jpeg"
	"os"

	"github.com/tdewolff/canvas"
)

var fontFamily *canvas.FontFamily

func main() {
	// 加载本地字体
	leMiaoSrc := "../static/HanYiLeMiao_Regular.ttf"

	fontFamily = canvas.NewFontFamily("LeMiao")
	fontFamily.Use(canvas.CommonLigatures)
	if err := fontFamily.LoadFontFile(leMiaoSrc, canvas.FontRegular); err != nil {
		panic(err)
	}

	serve()
	testFont()
}

func serve() {
	gin.SetMode(setting.ServerSetting.RunMode)

	// middlewares.Wechat()

	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// r.Use(router.init())

	routersInit := routes.Init(r)
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
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

func testFont() {
	c := canvas.New(265, 90)
	ctx := canvas.NewContext(c)
	draw(ctx)
	c.SavePNG("out.png", 5.0)
}

func drawText(c *canvas.Context, x, y float64, halign, valign canvas.TextAlign, indent float64) {
	face := fontFamily.Face(6.0, color.Black, canvas.FontRegular, canvas.FontNormal)
	phrase := "测试的文本一，phase"

	text := canvas.NewTextBox(face, phrase, 60.0, 35.0, halign, valign, indent, 0.0)
	rect := text.Bounds()
	rect.Y = 0.0
	rect.H = -35.0
	c.SetFillColor(canvas.Whitesmoke)
	c.DrawPath(x, y, rect.ToPath())
	c.SetFillColor(canvas.Black)
	c.DrawText(x, y, text)
}

func drawImage(c *canvas.Context) {
	head, err := os.Open("../static/head.jpeg")
	if err != nil {
		panic(err)
	}

	img, err := jpeg.Decode(head)
	if err != nil {
		panic(err)
	}

	// x, y float64, img image.Image, dpm float64
	// Scale(1.0/dpm, 1.0/dpm)
	c.DrawImage(0, 0, img, 1)
}

func draw(c *canvas.Context) {
	face := fontFamily.Face(14.0, color.Black, canvas.FontRegular, canvas.FontNormal)
	c.SetFillColor(canvas.Black)
	c.DrawText(132.5, 88.0, canvas.NewTextBox(face, "测试的文本二", 0.0, 0.0, canvas.Center, canvas.Top, 0.0, 0.0))

	drawText(c, 5.0, 80.0, canvas.Left, canvas.Top, 10.0)
	drawText(c, 70.0, 80.0, canvas.Center, canvas.Top, 10.0)
	drawText(c, 135.0, 80.0, canvas.Right, canvas.Top, 10.0)
	drawText(c, 200.0, 80.0, canvas.Justify, canvas.Top, 10.0)
	drawText(c, 5.0, 40.0, canvas.Left, canvas.Top, 10.0)
	drawText(c, 70.0, 40.0, canvas.Left, canvas.Center, 10.0)
	drawText(c, 135.0, 40.0, canvas.Left, canvas.Bottom, 10.0)
	drawText(c, 200.0, 40.0, canvas.Left, canvas.Justify, 10.0)

	// test drawImage
	drawImage(c)
}
