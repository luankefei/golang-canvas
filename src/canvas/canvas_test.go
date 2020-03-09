package canvas

import (
	"fmt"
	"testing"
)

// 测试画图主入口的参数传入
func TestCreateImage(t *testing.T) {
	var data []Drawer

	i := Image{
		x: 1,
		y: 2,
	}
	data = append(data, &i)
	global := GlobalConfig{}

	CreateImage(data, global)
}

func TestRGBAToColor(t *testing.T) {
	fmt.Println(RGBAToColor("rgba(255,255,255,1)"))
}

func TestHexToColor(t *testing.T) {
	fmt.Println(HexToColor("#F58C4BFF"))
}
