package canvas

import (
	"bufio"
	"encoding/base64"
	"fmt"
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
	fmt.Println("draw_data", d)

	ctx := canvas.NewContext(c)

	// 对画布进行一次位移调整，拉回到左上角
	// 注意绘制反向仍然是反向
	matrix := canvas.Identity.Translate(0, c.H)
	ctx.SetView(matrix)

	// // 数据转换
	// text := Text{
	// 	X:          100,
	// 	Y:          100,
	// 	Size:       90,
	// 	LineHeight: 40,
	// 	Color:      "#000000",
	// 	Content:    "测试的文字，我爱北京天安门",
	// 	FontStyle:  400,
	// 	FontFamily: "PingFang",
	// }

	// text.Draw(ctx)

	fmt.Println("canvas_draw", c)

	text := Text{
		X: 45,
		Y: 277.5,
		// Y:          185,
		Size:       39,
		LineHeight: 39,
		Color:      "#000000",
		Content:    "连续早起",
		FontStyle:  400,
		FontFamily: "PingFang",
	}

	text.Draw(ctx)

	text = Text{
		X: 45,
		Y: 336,
		// Y:          224,
		Size:       99,
		LineHeight: 99,
		Color:      "#000000",
		Content:    "154",
		FontStyle:  700,
		FontFamily: "PingFang",
	}

	text.Draw(ctx)

	text = Text{
		X: 45,
		Y: 472.5,
		// Y:          315,
		Size:       39,
		LineHeight: 39,
		Color:      "#000000",
		Content:    "今日早起",
		FontStyle:  700,
		FontFamily: "PingFang",
	}

	text.Draw(ctx)

	text = Text{
		X: 45,
		Y: 528,
		// Y:          352,
		Size:       99,
		LineHeight: 99,
		Color:      "#000000",
		Content:    "10:22",
		FontStyle:  700,
		FontFamily: "PingFang",
	}

	text.Draw(ctx)

	text = Text{
		X: 45,
		Y: 694.5,
		// Y:          463,
		Size:       36,
		LineHeight: 36,
		Color:      "#000000",
		Content:    "10906993人正在参与 比185万人起的早",
		FontStyle:  700,
		FontFamily: "PingFang",
		Limit:      250,
	}

	text.Draw(ctx)

	text = Text{
		X: 967.5,
		// X: 645,
		Y: 25.5,
		// Y:          17,
		Size:       75,
		LineHeight: 75,
		Color:      "#000000",
		Content:    "13",
		FontStyle:  700,
		FontFamily: "PingFang",
		// Limit:      217,
		Limit: 73,
		Align: 1, // right
	}

	text.Draw(ctx)

	text = Text{
		X: 967.5,
		// X: 645,
		Y: 120,
		// Y:          80,
		Size:       30,
		LineHeight: 30,
		Color:      "#000000",
		Content:    "2020.03",
		FontStyle:  700,
		FontFamily: "PingFang",
		Align:      1, // right
		// Limit:      217,
		Limit: 73,
	}

	text.Draw(ctx)
}

// CreateImage is api entry
func CreateImage(d []Drawer, g GlobalConfig) {
	Setup()

	c := Canvas{canvas.New(g.Width, g.Height)}

	// 对*canvas.Draw函数传入绘图数据
	c.Draw(d)

	// SavePNG的第二个参数是canvas导出时放大的倍数
	// 尽量导出2x或者3x的尺寸，但坐标是1x的，需要更多测试
	c.SavePNG(g.FileName, 1)

	fmt.Println("create image", len(d), g)
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

	fmt.Printf("enc=[%s]\n", encoded)

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
