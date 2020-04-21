package canvas

import (
	"fmt"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"testing"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/rasterizer"
)

// 测试画图主入口的参数传入
func TestCreateImage(t *testing.T) {
	// TODO: 临时测试代码
	Setup()

	// 塞到[]drawer之前就拆成两个数组
	data := []Drawer{
		&Image{
			X:        0,
			Y:        0,
			Width:    1125,
			Height:   1125,
			ImageURL: "https://img.laiye.com/checkinAlbum_20200316083737_htqvLACrln.jpg",
			Resize:   true,
		},
		&Image{
			Name: "avatar",
			ImageURL:     "https://img.laiye.com/cLvgXicdq4RMvFgMeyiarFciatqCEPrkGudP9N6SceHhmA4Tl2unDvK4rNVCFroJZcfqMnUGvnBeDLaZpDYW0TRl9lxmD47gs70.jpg",
			X:            48,
			Y:            48,
			Width:        150,
			Height:       150,
			BorderRadius: 75,
			Resize:       true,
		},
		&Image{
			ImageURL: "https://img.laiye.com/material_20200407053339_aiGLseDwFQ.jpeg",
			X:      930,
			Y:      930,
			Height: 150,
			Width:  150,
			Resize: true,
			Clip: ImageClip{
				Width:  370,
				Height: 370,
				X:      30,
				Y:      30,
			},
		},
		&Text{
			Name: "chickenMessage"
			Content:    "之所以能，是因为相信能",
			X:          270,
			Y:          975,
			Limit:      615,
			Size:       39,
			Align:      1,
			LineHeight: 45,
			FontStyle:  700,
			Color:      "#ffffff",
			FontFamily: "PingFang",
		},
		&Text{
			Name: "welcomeMessage"
			Content:    "扫码和我互道早安",
			X:          570,
			Y:          1035,
			Size:       39,
			FontStyle:  400,
			Color:      "#ffffff",
			FontFamily: "PingFang",
		},
		&Text{
			X:          45,
			Y:          277.5,
			Size:       39,
			LineHeight: 39,
			Color:      "#ffffff",
			Content:    "连续早起",
			FontStyle:  400,
			FontFamily: "PingFang",
		},
		&Text{
			X:          45,
			Y:          336,
			Size:       99,
			LineHeight: 99,
			Color:      "#ffffff",
			Content:    "154",
			FontStyle:  700,
			FontFamily: "PingFang",
		},
		&Text{
			X:          45,
			Y:          472.5,
			Size:       39,
			LineHeight: 39,
			Color:      "#ffffff",
			Content:    "今日早起",
			FontStyle:  700,
			FontFamily: "PingFang",
		},
		&Text{
			X:          45,
			Y:          528,
			Size:       99,
			LineHeight: 99,
			Color:      "#ffffff",
			Content:    "10:22",
			FontStyle:  700,
			FontFamily: "PingFang",
		},
		&Text{
			X:          45,
			Y:          694.5,
			Size:       36,
			LineHeight: 36,
			Color:      "#ffffff",
			Content:    "10906993人正在参与 比185万人起的早",
			FontStyle:  700,
			FontFamily: "PingFang",
			Limit:      500,
		},
		&Text{
			X:          967.5,
			Y:          25.5,
			Size:       75,
			LineHeight: 75,
			Color:      "#ffffff",
			Content:    "13",
			FontStyle:  700,
			FontFamily: "PingFang",
			Limit:      73,
			Align:      1, // right
		},
		&Text{
			X:          967.5,
			Y:          120,
			Size:       30,
			LineHeight: 30,
			Color:      "#ffffff",
			Content:    "2020.03",
			FontStyle:  700,
			FontFamily: "PingFang",
			Align:      1, // right
			Limit:      73,
		},
	}

	global := GlobalConfig{
		Width:    1125,
		Height:   1125,
		FileName: "test_mask.png",
	}

	CreateImage(data, global)
}

func TestCreateQcode(t *testing.T) {
	// 塞到[]drawer之前就拆成两个数组
	data := []Drawer{
		&Image{
			ImageURL: "https://img.laiye.com/material_20200407053339_aiGLseDwFQ.jpeg",
			X:        0,
			Y:        0,
			Height:   100,
			Width:    100,
		},
	}

	global := GlobalConfig{
		Width:    400,
		Height:   400,
		FileName: "test_qcode.png",
	}

	// fmt.Println(data, global)
	CreateImage(data, global)

}

func TestCropImage(t *testing.T) {
	c := Canvas{canvas.New(750, 750)}
	ctx := canvas.NewContext(c)

	// to draw a red background
	ctx.SetFillColor(color.RGBA{0xff, 0x00, 0x00, 0xff})
	ctx.DrawPath(0, 0, canvas.Rectangle(750, 750))

	head, _ := os.Open("../static/crop_head.png")
	img, _ := png.Decode(head)
	ctx.DrawImage(0, 0, img, 1)

	qcode, _ := os.Open("../static/crop_qrcode.jpeg")
	qcodeImg, _ := jpeg.Decode(qcode)
	ctx.DrawImage(0, 200, qcodeImg, 1)

	c.WriteFile("test_crop.png", rasterizer.PNGWriter(1))
}

func TestRGBAToColor(t *testing.T) {
	fmt.Println(RGBAToColor("rgba(255,255,255,1)"))
}

func TestHexToColor(t *testing.T) {
	fmt.Println(HexToColor("#F58C4BFF"))
}

func TestLoadImageFilter(t *testing.T) {
	var arr = []Drawer{
		&Image{
			X:        0,
			Y:        0,
			Width:    132,
			Height:   132,
			ImageURL: "https://img.laiye.com/qE9MKluetOntzRtRbNltRhIpicn8ktDDNbTPiaGCv1CrIoPQor5Iw7Q6LM78qJft8ncTFCze3S4JHzpLEqiclrCJg.jpg",
		},
	}
	testArr := make([]interface{}, len(arr))
	for i, v := range arr {
		testArr[i] = v
	}

	images := LoadImageFilter(testArr)

	fmt.Println("TestLoadImageFilter finish", len(images))
}

// func TestCropImage(t *testing.T) {
// 	prepareImage(d)

// 	c := Canvas{canvas.New(1125, 1125)}

// 	ctx := canvas.NewContext(c)

// 	// to draw a red background
// 	ctx.SetFillColor(color.RGBA{0xff, 0x00, 0x00, 0xff})
// 	ctx.DrawPath(0, 0, canvas.Rectangle(1125, 1125))

// 	// 对*canvas.Draw函数传入绘图数据
// 	c.Draw(d)

// 	// SavePNG的第二个参数是canvas导出时放大的倍数
// 	// 尽量导出2x或者3x的尺寸，但坐标是1x的，需要更多测试
// 	c.SavePNG("test_crop.png", 1)
// }
