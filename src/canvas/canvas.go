package canvas

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"image/color"

	"io/ioutil"
	"os"

	"github.com/luankefei/golang-canvas/src/libs"
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/rasterizer"
)

// Canvas is alias of *canvas.Canvas
type Canvas struct {
	*canvas.Canvas
}

// Draw canvas.draw
func (c *Canvas) Draw(d []Drawer) {
	fmt.Println("draw_data", len(d))

	ctx := canvas.NewContext(c)

	// 对画布进行一次位移调整，拉回到左上角
	// 注意绘制反向仍然是反向
	matrix := canvas.Identity.Translate(0, c.H)

	// 注意这里的对ctx.SetView不会影响到外侧CreateImage的ctx
	ctx.SetView(matrix)

	for _, v := range d {
		v.Draw(ctx)
	}

	// TODO: test draw line
	// drawline不受transform的影响
	p := &canvas.Path{}
	p.LineTo(405, 0)
	p.Close()

	// ctx.SetFillColor(color.White)
	ctx.SetStrokeColor(color.White)
	ctx.SetStrokeWidth(2)
	ctx.DrawPath(45, -669, p)
}

func prepareImage(d []Drawer) []*Image {
	testArr := make([]interface{}, len(d))
	for i, v := range d {
		testArr[i] = v
	}

	images := LoadImageFilter(testArr)
	for _, v := range images {
		// TODO: traceId在http header内，image_key用来做兜底方案
		v.fetch()
	}
	return images
}

// CreateImage is api entry
func CreateImage(d []Drawer, g GlobalConfig) {
	// Setup()

	// 先进行图片加载
	// TODO: 需要观察图片是否被正确加载
	prepareImage(d)

	fmt.Println("create image", len(d), g)

	c := Canvas{canvas.New(g.Width, g.Height)}

	ctx := canvas.NewContext(c)

	// to draw a red background
	ctx.SetFillColor(color.RGBA{0xff, 0x00, 0x00, 0xff})
	ctx.DrawPath(0, 0, canvas.Rectangle(g.Width, g.Height))

	// 对*canvas.Draw函数传入绘图数据
	c.Draw(d)

	// SavePNG的第二个参数是canvas导出时放大的倍数
	// 尽量导出2x或者3x的尺寸，但坐标是1x的，需要更多测试
	c.WriteFile(g.FileName, rasterizer.PNGWriter(1.0))

	fmt.Println("draw_data_finish", len(d))
}

// uploadImg todo...
func uploadImg() {}

// 文件转base64
func toBase64(filepath string) string {
	f, _ := os.Open(filepath)

	// Read entire JPG into byte slice.
	reader := bufio.NewReader(f)

	content, _ := ioutil.ReadAll(reader)

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)

	return encoded
}

// Setup 整个绘图模块的初始化
func Setup() {
	libs.Setup()

	InitFont()
}
