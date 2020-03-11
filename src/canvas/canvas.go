package canvas

import (
	"fmt"

	"github.com/luankefei/golang-canvas/src/libs"
	"github.com/tdewolff/canvas"
)

// Canvas is alias of *canvas.Canvas
type Canvas struct {
	*canvas.Canvas
}

// Draw canvas.draw
func (c *Canvas) Draw() {
	fmt.Println("canvas_draw", c)

	ctx := canvas.NewContext(c)

	matrix := canvas.Identity.Translate(0, c.H)
	ctx.SetView(matrix)

	text := Text{
		X:          30,
		Y:          193,
		Size:       24,
		LineHeight: 24,
		Color:      "#000000",
		Content:    "连续早起",
		FontStyle:  400,
		FontFamily: "PingFang",
	}

	text.Draw(ctx)

	text = Text{
		X:          30,
		Y:          227,
		Size:       66,
		LineHeight: 66,
		Color:      "#000000",
		Content:    "999",
		FontStyle:  700,
		FontFamily: "PingFang",
	}

	text.Draw(ctx)

	text = Text{
		X:          30,
		Y:          323,
		Size:       24,
		LineHeight: 24,
		Color:      "#000000",
		Content:    "今日早起",
		FontStyle:  400,
		FontFamily: "PingFang",
	}

	text.Draw(ctx)

	text = Text{
		X:          30,
		Y:          357,
		Size:       66,
		LineHeight: 66,
		Color:      "#000000",
		Content:    "06:00",
		FontStyle:  700,
		FontFamily: "PingFang",
	}

	text.Draw(ctx)

	text = Text{
		X:          30,
		Y:          466,
		Size:       24,
		LineHeight: 30,
		Color:      "#000000",
		Content:    "2424160人正在参与 比90%的人起的早",
		FontStyle:  400,
		FontFamily: "PingFang",
		Limit:      230,
	}

	text.Draw(ctx)

	text = Text{
		X:          652,
		Y:          24,
		Size:       50,
		LineHeight: 50,
		Color:      "#000000",
		Content:    "08",
		FontStyle:  700,
		FontFamily: "PingFang",
		// Limit:      217,
		Limit: 73,
		Align: 1,
	}

	text.Draw(ctx)

	text = Text{
		X:          652,
		Y:          83,
		Size:       20,
		LineHeight: 20,
		Color:      "#000000",
		Content:    "2019.07",
		FontStyle:  700,
		FontFamily: "PingFang",
		Align:      1,
		// Limit:      217,
		Limit: 73,
	}

	text.Draw(ctx)
}

// CreateImage is api entry
func CreateImage(d []Drawer, g GlobalConfig) {
	// TODO: 临时测试代码
	Setup()

	c := Canvas{canvas.New(750, 750)}
	c.Draw()

	// 尽量导出2x或者3x的尺寸，但坐标是1x的，需要更多测试
	c.SavePNG("out.png", 1.0)

	fmt.Println("create image", len(d), g)
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
