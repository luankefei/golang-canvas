package canvas

import (
	"fmt"
	"testing"
)

// 测试画图主入口的参数传入
func TestCreateImage(t *testing.T) {
	var aaa []Drawer
	fmt.Println(aaa)

	i := Image{
		x: 1,
		y: 2,
	}
	aaa = append(aaa, &i)

	// a := []Image{i}

	global := GlobalConfig{}

	CreateImage(aaa, global)
}
