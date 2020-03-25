package canvas

import (
	"context"
	"fmt"
	"testing"
)

// 测试画图主入口的参数传入
func TestCreateImage(t *testing.T) {
	// TODO: 临时测试代码
	Setup()

	// 塞到[]drawer之前就拆成两个数组
	data := []Drawer{
		&Image{
			X:        1,
			Y:        2,
			Width:    1125,
			Height:   1125,
			ImageURL: "https://img.laiye.com/checkinAlbum_20200309105259_wPkT1VY88V.jpg",
		}, &Text{
			X:          45,
			Y:          277.5,
			Size:       39,
			LineHeight: 39,
			Color:      "#000000",
			Content:    "连续早起",
			FontStyle:  400,
			FontFamily: "PingFang",
		},
		&Text{
			X:          45,
			Y:          336,
			Size:       99,
			LineHeight: 99,
			Color:      "#000000",
			Content:    "154",
			FontStyle:  700,
			FontFamily: "PingFang",
		},
		&Text{
			X:          45,
			Y:          472.5,
			Size:       39,
			LineHeight: 39,
			Color:      "#000000",
			Content:    "今日早起",
			FontStyle:  700,
			FontFamily: "PingFang",
		},
		&Text{
			X:          45,
			Y:          528,
			Size:       99,
			LineHeight: 99,
			Color:      "#000000",
			Content:    "10:22",
			FontStyle:  700,
			FontFamily: "PingFang",
		},
		&Text{
			X:          45,
			Y:          694.5,
			Size:       36,
			LineHeight: 36,
			Color:      "#000000",
			Content:    "10906993人正在参与 比185万人起的早",
			FontStyle:  700,
			FontFamily: "PingFang",
			Limit:      250,
		},
		&Text{
			X:          967.5,
			Y:          25.5,
			Size:       75,
			LineHeight: 75,
			Color:      "#000000",
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
			Color:      "#000000",
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
		FileName: "out.png",
	}

	CreateImage(data, global)
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
			ImageURL: "https://img.laiye.com/qE9MKluetOntzRtRbNltRhIpicn8ktDDNbTPiaGCv1CrIoPQor5Iw7Q6LM78qJft8ncTFCze3S4JHzpLEqiclrCJg.jpg",
		},
	}
	testArr := make([]interface{}, len(arr))
	for i, v := range arr {
		testArr[i] = v
	}

	images := LoadImageFilter(testArr)
	for _, v := range images {
		// TODO: traceId在http header内，image_key用来做兜底方案
		fetchOneImage(context.Background(), "traceId", "businessName", v, "image_key")
	}

	fmt.Println("TestLoadImageFilter finish", images[0])
}
