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

	text := Text{
		X:          50,
		Y:          50,
		Size:       20,
		LineHeight: 40,
		Color:      "#000000",
		Content:    "测试的文字，我爱北京天安门。索尼大法好",
		FontStyle:  400,
		FontFamily: "PingFang",
	}

	ctx := canvas.NewContext(c)
	text.Draw(ctx)
}

// CreateImage is api entry
func CreateImage(d []Drawer, g GlobalConfig) {
	// TODO: 临时测试代码
	Setup()

	c := Canvas{canvas.New(750, 750)}
	c.Draw()

	// // 尽量导出2x或者3x的尺寸，但坐标是1x的，需要更多测试
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
