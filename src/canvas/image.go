package canvas

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"image/draw"
	"os"

	// "github.com/llgcode/draw2d"
	"github.com/rs/zerolog/log"
	"github.com/tdewolff/canvas"
	"github.com/disintegration/imaging"
	// "github.com/golang/freetype/raster"
)

// 兜底图片方案的image_key在Image对象中传递
func (i *Image) fetch() error {
	var err error

	if i.ImageURL == "" {
		err = errors.New("image url can not be empty")
		log.Info().
			Err(err).
			Msg("fetch one image failed, error")
		return err
	}
	var img []byte
	var t string
	// if image.LocalBufferName != "" {
	// 	name := fmt.Sprintf("%s::%s", businessName, image.LocalBufferName)
	// 	img, t, err = b.getLocalBufferImage(name, image)
	// 	libs.Log().Info("trace_id(%s) get image(%+v) from local buffer, err(%+v)", traceId, image, err)
	// } else {
	img, t, err = loadImageByteFromRemote(i.ImageURL, i.Width, i.Height)
	// libs.Log().Info("trace_id(%s) get image(%+v) from remote, err(%+v)", traceId, image, err)
	// }

	i.MimeType = t
	i.Buffer = img
	return nil
}
// imageWidth: number,
// imageHeight: number,
// offsetX: number,
// offsetY: number,
// radius: number | number[]
// func RoundMask(w float64, h float64, r float64) {
// 	// p := &Path{}
// 	// p.MoveTo(rx, 0.0)
// 	// p.ArcTo(rx, ry, 0.0, false, true, -rx, 0.0)
// 	// p.ArcTo(rx, ry, 0.0, false, true, rx, 0.0)
// 	// p.Close()
// 	return p
// }

// ClipPreserve updates the clipping region by intersecting the current
// clipping region with the current path as it would be filled by dc.Fill().
// The path is preserved after this operation.
func ClipPreserve(c *canvas.Context, i *Image) {
	// var mask Path
	mask := Path{canvas.RoundedRectangle(100, 100, 20)}
	tempPath := imaging.Resize(&mask, int(100), int(100), imaging.Linear)

	rect := image.Rect(0, 0, 1000, 1000)
	dst := image.NewAlpha(rect)
	// &Circle{pt, dx / 2}, nil
	// golang-draw传入的是图片的起点，取负值不知道为什么
	pt := image.Pt(-int(0), -int(0))
	// painter := raster.NewAlphaOverPainter(clip)
	// dc.fill(painter)
	// if dc.mask == nil {
	// 	dc.mask = clip
	// } else {
		// mask := image.NewAlpha(image.Rect(0, 0, dc.width, dc.height))


	img, _ := ConvertBytesToImage(i.Buffer, i.MimeType)

  // image.ZP
	draw.DrawMask(dst, dst.Bounds(), img, pt, tempPath, pt, draw.Over)
	// dc.mask = mask
	// }
}


// Draw image
func (i *Image) Draw(c *canvas.Context) {
	// fmt.Println("===1")
	head, _ := os.Open("../static/head.jpeg")
	img, _ := jpeg.Decode(head)
	c.DrawImage(-1, -1, img, 1)

	bg, _ := os.Open("../static/background.jpeg")
	bgImg, _ := jpeg.Decode(bg)
	c.DrawImage(0, 0, bgImg, 1)
	// draw.DrawMask(dst, dst.Bounds(), src, image.ZP, &circle{p, r}, image.ZP, draw.Over)
	// c.DrawImage(0, 0, img, 1)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("===2")

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("===3")
	// // x, y float64, img image.Image, dpm float64
	// // Scale(1.0/dpm, 1.0/dpm)
	// c.DrawImage(0, 0, img, 1)

	// fmt.Println("===4")
	image, _ := ConvertBytesToImage(i.Buffer, i.MimeType)
	fmt.Println("=== image_draw: ", i.MimeType, i.ImageURL, i.Width, i.Height)

	// 获取实际图片尺寸和传入参数之间的比例
	scale := (float64(image.Bounds().Dx()) / i.Width)
	fmt.Println("=== scale", scale)
	// c.DrawImage(i.X, i.Y*-1-i.Height, image, scale)
	c.DrawImage(i.X, i.Y, image, scale)
}

// ConvertBytesToImage 将二进制流转化为图片
func ConvertBytesToImage(imgByte []byte, t string) (image.Image, error) {
	fmt.Println("===11")
	if len(imgByte) == 0 {
		return nil, errors.New("can not convert empty image")
	}
	var img image.Image
	var err error
	imgIoReader := bytes.NewReader(imgByte)
	fmt.Println("===22")
	switch t {
	case "image/jpeg":
		img, err = jpeg.Decode(imgIoReader)
	case "image/png":
		img, err = png.Decode(imgIoReader)
	}
	fmt.Println("===33")
	if err != nil {
		log.Info().Msgf("gzh_test jpg png failed err (%v) img(%d) type(%v)", err, len(imgByte), t)
	} else {
		log.Info().Msgf("gzh_test jpg png success img(%d) type(%v)", len(imgByte), t)
	}

	fmt.Println("===444")
	return img, err
}

// func drawImage(c *canvas.Context) {
// 	head, err := os.Open("../static/head.jpeg")
// 	if err != nil {
// 		panic(err)
// 	}

// 	img, err := jpeg.Decode(head)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// x, y float64, img image.Image, dpm float64
// 	// Scale(1.0/dpm, 1.0/dpm)
// 	c.DrawImage(0, 0, img, 1)
// }
