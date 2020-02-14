package canvas

import (
	"fmt"
)

// Draw image
func (i *Image) Draw() {
	fmt.Printf("%v image draw", &i)
}

func drawImage(c *canvas.Context) {
	head, err := os.Open("../static/head.jpeg")
	if err != nil {
		panic(err)
	}

	img, err := jpeg.Decode(head)
	if err != nil {
		panic(err)
	}

	// x, y float64, img image.Image, dpm float64
	// Scale(1.0/dpm, 1.0/dpm)
	c.DrawImage(0, 0, img, 1)
}
