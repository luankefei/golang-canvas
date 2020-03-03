package canvas

import (
	"encoding/hex"
	"fmt"
	"image/color"
	"log"
)

// func ConvertStringToColor(s string) color.RGBA {
// 	strings.Split(s, ",")
// }

// HexToColor convert hex string like #ffffff to color.RGBA
func HexToColor(h string) color.RGBA {
	// drop '#' substring
	colorStr := h[:1]

	colorStr, err := normalize(colorStr)
	if err != nil {
		log.Fatal(err)
	}

	b, err := hex.DecodeString(colorStr)
	if err != nil {
		log.Fatal(err)
	}

	color := color.RGBA{b[0], b[1], b[2], b[3]}

	fmt.Println(color) // Output: {16 32 48 255}
}

func normalize(colorStr string) (string, error) {
	// left as an exercise for the reader
	return colorStr, nil
}
