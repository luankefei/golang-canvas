package canvas

import (
	"encoding/hex"
	"fmt"
	"image/color"
	"log"
	"strconv"
	"strings"
)

// LoadImageFilter drawer
// TODO: 这里后续可以加入buffer_name和expires的检查
func LoadImageFilter(arr []interface{}) []Image {
	d := make([]Image, 0)

	for _, v := range arr {
		switch v.(type) {
		case Image:
			{
				d = append(d, v.(Image))
			}
		}
	}

	return d
}

// RGBAToColor convert rgba string like rgba(255, 255, 255, 1) to color.RGBA
func RGBAToColor(s string) color.RGBA {
	var t2 = []byte{}

	// drop "rgba(" and ")"
	t := strings.Split(s[5:len(s)-1], ",")

	for _, i := range t {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		t2 = append(t2, byte(j))
	}
	fmt.Println(t2)

	return color.RGBA{t2[0], t2[1], t2[2], t2[3]}
}

// HexToColor convert hex string like #ffffff to color.RGBA
func HexToColor(h string) color.RGBA {
	// TODO: 需要自动补全最后两位
	if len(h) == 7 {
		h = h + "ff"
	}
	// drop '#' substring
	colorStr := h[1:]

	fmt.Println("colorStr", colorStr)

	colorStr, err := normalize(colorStr)
	if err != nil {
		log.Fatal(err)
	}

	b, err := hex.DecodeString(colorStr)
	if err != nil {
		log.Fatal(err)
	}

	color := color.RGBA{b[0], b[1], b[2], b[3]}

	return color

	// fmt.Println(color) // Output: {16 32 48 255}
}

func normalize(colorStr string) (string, error) {
	// left as an exercise for the reader
	return colorStr, nil
}

// func (b *FetchPic) fetcheImage(traceId string, businessName string, image *draw_pb.Image,
// 	image_key string) error {

// }
