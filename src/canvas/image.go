package canvas

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"

	"github.com/rs/zerolog/log"
	"github.com/tdewolff/canvas"
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

// Draw image
func (i *Image) Draw(c *canvas.Context) {
	// fmt.Println("===1")
	// head, err := os.Open("../static/head.jpeg")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("===2")

	// img, err := jpeg.Decode(head)
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
	c.DrawImage(i.X, i.Y*-1-i.Height, image, scale)
	// c.DrawImage(0, 0, image, 1)
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
