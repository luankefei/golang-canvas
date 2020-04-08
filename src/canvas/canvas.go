package canvas

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"image/color"

	// "image/color"
	"io/ioutil"
	"os"

	"github.com/luankefei/golang-canvas/src/libs"
	"github.com/tdewolff/canvas"
)

// Canvas is alias of *canvas.Canvas
type Canvas struct {
	*canvas.Canvas
}

// Draw canvas.draw
func (c *Canvas) Draw(d []Drawer) {
	fmt.Println("draw_data", len(d))

	// c.Fit(1.0)
	ctx := canvas.NewContext(c)
	// ctx.SetFillColor(color.RGBA{0xff, 0x00, 0x00, 0xff})
	// ctx.SetFillColor(color.RGBA{0, 0, 0, 100})
	// ctx.DrawPath(0, 0, canvas.Rectangle(1125, 1125))

	// 对画布进行一次位移调整，拉回到左上角
	// 注意绘制反向仍然是反向
	matrix := canvas.Identity.Translate(0, c.H)
	ctx.SetView(matrix)

	for _, v := range d {
		v.Draw(ctx)
	}
	fmt.Println("draw_data_finish", len(d))
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
	c.SavePNG(g.FileName, 1)
}

// func compatibleData(d Drawer) {
// 	switch (d.type) {

// 	}

// }

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

	// fmt.Printf("enc=[%s]\n", encoded)

	return encoded
}

// Setup 整个绘图模块的初始化
func Setup() {
	libs.Setup()

	InitFont()
}

// async createImage(params: any, globalConfig?: any) {
// 	try {
// 		// step 0 先整理数据，按照各绘图数据类型完善字段
// 		const data = params.map((item: any) => compatibleDataV2(item.type, item))

// 		// step 2: 设置基础设置
// 		const ctx = this.context
// 		ctx.textBaseline = defaultConfig.textBaseline

// 		// step 3-1: clear canvas
// 		this.clear()

// 		// step 3-2: 加载所有的图片资源
// 		const imageList = await prepareImage(
// 			data.filter((item: any) => item.type === EElementType.image)
// 		)

// 		// 绘制所有的元素
// 		data.forEach((item: any) => {
// 			const { shadow } = globalConfig
// 			if (shadow && shadow.rangeList.indexOf(item.type) !== -1) {
// 				ctx.shadowOffsetX = shadow.offsetX || 0 // 阴影Y轴偏移
// 				ctx.shadowOffsetY = shadow.offsetY || 0 // 阴影X轴偏移
// 				ctx.shadowBlur = shadow.blur || 0 // 模糊尺寸
// 				ctx.shadowColor = shadow.color || '' // 颜色
// 			}

// 			switch (item.type) {
// 				case EElementType.image:
// 					drawImage.call(this, [item], [imageList.shift()])
// 					break

// 				case EElementType.text:
// 					this.drawText([item])
// 					break

// 				case EElementType.line:
// 					this.drawLine([item])
// 					break

// 				case EElementType.rect:
// 					this.drawRect([item])
// 					break

// 				case EElementType.arc:
// 					this.drawArc([item])
// 					break
// 			}

// 			// 重置阴影属性
// 			if (shadow) {
// 				ctx.shadowOffsetX = 0 // 阴影Y轴偏移
// 				ctx.shadowOffsetY = 0 // 阴影X轴偏移
// 				ctx.shadowBlur = 0 // 模糊尺寸
// 				ctx.shadowColor = '' // 颜色
// 			}
// 		})

// 		return this

// 		// 尝试对外侧抛出异常
// 	} catch (e) {
// 		logger.error('draw_error_1', { params, e: JSON.stringify(e) })
// 		return Promise.reject(e)
// 	}
// }
