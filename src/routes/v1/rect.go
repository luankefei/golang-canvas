package v1

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"image/color"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tdewolff/canvas"
)

// GetRect for api test
func GetRect(ginc *gin.Context) {
	c := canvas.New(265, 90)
	ctx := canvas.NewContext(c)

	ctx.SetView(canvas.Identity.Translate(0.0, 0.0))
	ctx.SetFillColor(color.RGBA{0, 0, 0, 100})
	ctx.DrawPath(10, 10, canvas.Rectangle(200, 50))

	// c := canvas.New(265, 90)
	// drawText(canvas.NewContext(c), 5.0, 80.0, canvas.Left, canvas.Top, 10.0)
	// img := c.WriteImage(5.0)
	// k
	// buf := make([]byte, ctx.Length())
	c.SavePNG("./out.png", 5.0)
	f, _ := os.Open("./out.png")

	// Read entire JPG into byte slice.
	reader := bufio.NewReader(f)

	content, _ := ioutil.ReadAll(reader)

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)

	// reader := bufio.NewReader(img)
	// reader.Read()
	// b := base64.StdEncoding.EncodeToString(buf)
	fmt.Printf("enc=[%s]\n", encoded)
	html := "data:image/png;base64," + encoded
	// ginc.PureJSON(http.StatusOK, gin.H{
	// 	"html": html,
	// })
	ginc.JSON(http.StatusOK, html)
	// ginc.HTML(http.StatusOK, html, gin.H{
	// 	"title": "test go canvas",
	// })

	// appG := app.Gin{C: ginc}
	// appG.Response(http.StatusOK, e.SUCCESS, data)
}

// func drawText(c *canvas.Context, x, y float64, halign, valign canvas.TextAlign, indent float64) {
// 	face := fontFamily.Face(6.0, color.Black, canvas.FontRegular, canvas.FontNormal)
// 	phrase := "测试的文本一，phase"

// 	text := canvas.NewTextBox(face, phrase, 60.0, 35.0, halign, valign, indent, 0.0)
// 	rect := text.Bounds()
// 	rect.Y = 0.0
// 	rect.H = -35.0
// 	c.SetFillColor(canvas.Whitesmoke)
// 	c.DrawPath(x, y, rect.ToPath())
// 	c.SetFillColor(canvas.Black)
// 	c.DrawText(x, y, text)
// }
