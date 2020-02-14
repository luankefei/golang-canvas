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

func toBase64(filepath string) {
	f, _ := os.Open(filepath)

	// Read entire JPG into byte slice.
	reader := bufio.NewReader(f)

	content, _ := ioutil.ReadAll(reader)

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)

	fmt.Printf("enc=[%s]\n", encoded)

	return encoded
}
