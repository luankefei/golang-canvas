package canvas

import (
	"fmt"
	"image/color"
	"image/jpeg"
	"os"
	"testing"

	"github.com/tdewolff/canvas"
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
		},
		&Image{
			// name: "avatar"
			// Type:         "image",
			ImageURL:     "https://img.laiye.com/cLvgXicdq4RMvFgMeyiarFciatqCEPrkGudP9N6SceHhmA4Tl2unDvK4rNVCFroJZcfqMnUGvnBeDLaZpDYW0TRl9lxmD47gs70.jpg",
			X:            48,
			Y:            48,
			Width:        150,
			Height:       150,
			BorderRadius: 75,
		},
		&Image{
			// name: "qrCode"
			// DrawType: "image",
			ImageURL: "https://img.laiye.com/G0gmxRbVb4EtnpEi1mUK2FkTk5HcCuK6mbX7lj5qFYbU3D0A.png",
			X:        930,
			Y:        930,
			Height:   150,
			Width:    150,
		},
		&Text{
			// name: "chickenMessage"
			// DrawType:   "text",
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
			// Name: "welcomeMessage"
			// DrawType: "text",
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
		FileName: "out1.png",
	}

	// fmt.Println(data, global)
	CreateImage(data, global)
}

func TestBasicDrawImage(t *testing.T) {
	c := Canvas{canvas.New(750, 750)}
	ctx := canvas.NewContext(c)

	// to draw a red background
	ctx.SetFillColor(color.RGBA{0xff, 0x00, 0x00, 0xff})
	ctx.DrawPath(0, 0, canvas.Rectangle(750, 750))

	// image1 at point(0,0)
	bg, _ := os.Open("../static/background.jpg")
	bgImg, _ := jpeg.Decode(bg)
	ctx.DrawImage(0, 0, bgImg, 1)

	// image2 start (-1,-1) trying to fix margin
	head, _ := os.Open("../static/head.jpg")
	img, _ := jpeg.Decode(head)
	ctx.DrawImage(-1, -1, img, 1)

	c.SavePNG("out.png", 1)
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
	// for _, v := range images {
	// 	// TODO: traceId在http header内，image_key用来做兜底方案
	// 	v.fetch()
	// }

	fmt.Println("TestLoadImageFilter finish", len(images))
}
