package canvas

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"

	// "os"

	// "github.com/llgcode/draw2d"
	"github.com/disintegration/imaging"
	"github.com/oliamb/cutter"
	"github.com/rs/zerolog/log"
	"github.com/tdewolff/canvas"
	// "github.com/golang/freetype/raster"
)

// Draw image
func (i *Image) Draw(c *canvas.Context) {
	img, _ := ConvertBytesToImage(i.Buffer, i.MimeType)

	// fmt.Println("=== scale", scale, i.Width)

	// test 图片切圆角功能
	// if i.BorderRadius > 0 {
	// 	dst := image.NewRGBA(image.Rect(0, 0, int(i.Width), int(i.Height)))
	// 	img, _ = ClipPreserve(c, i, dst)

	// 	fmt.Println("=====ClipPreserve", img.Bounds())
	// }

	// test 图片边缘裁剪
	// 获取实际图片尺寸和传入参数之间的比例
	scale := (float64(img.Bounds().Dx()) / i.Width)

	// TODO: 阿里云oss resize不会对图片进行放大,只能缩小，所以绘制时仍然要依赖绘图库进行图片缩放
	// TODO: 2020.4.7 考虑将计算缩放比率放到切图之后
	if i.Clip.Width > 0 {
		croppedImg, _ := cutter.Crop(img, cutter.Config{
			Width:  i.Clip.Width,
			Height: i.Clip.Height,
			Anchor: image.Point{i.Clip.X, i.Clip.Y},
		})
		img = croppedImg
	}

	fmt.Println("=== image_draw: ", i.MimeType, i.ImageURL, i.Width, i.Height, scale, img.Bounds().Dx(), img.Bounds())
	c.DrawImage(i.X, i.Y*-1-i.Height, img, scale)
}

// ConvertBytesToImage 将二进制流转化为图片
func ConvertBytesToImage(imgByte []byte, t string) (image.Image, error) {
	if len(imgByte) == 0 {
		return nil, errors.New("can not convert empty image")
	}
	var img image.Image
	var err error
	imgIoReader := bytes.NewReader(imgByte)
	switch t {
	case "image/jpeg":
		img, err = jpeg.Decode(imgIoReader)
	case "image/png":
		img, err = png.Decode(imgIoReader)
	}
	if err != nil {
		log.Info().Msgf("gzh_test jpg png failed err (%v) img(%d) type(%v)", err, len(imgByte), t)
	} else {
		log.Info().Msgf("gzh_test jpg png success img(%d) type(%v)", len(imgByte), t)
	}

	return img, err
}

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
	// TODO: 图片缓存
	// if image.LocalBufferName != "" {
	// 	name := fmt.Sprintf("%s::%s", businessName, image.LocalBufferName)
	// 	img, t, err = b.getLocalBufferImage(name, image)
	// 	libs.Log().Info("trace_id(%s) get image(%+v) from local buffer, err(%+v)", traceId, image, err)
	// } else {
	img, t, err = loadImageByteFromRemote(i.ImageURL, i.Resize, i.Width, i.Height)
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
func ClipPreserve(c *canvas.Context, i *Image, dst *image.RGBA) (draw.Image, error) {
	// fmt.Println("ClipPreserve === 1", c.Width(), c.Height(), i.BorderRadius)
	// var mask Path
	mask := Path{*canvas.RoundedRectangle(i.Width, i.Height, i.BorderRadius), i}
	// fmt.Println("ClipPreserve === 2", mask)
	tempPath := imaging.Resize(&mask, int(i.Width), int(i.Height), imaging.Linear)

	// fmt.Println("ClipPreserve === TODO 3")
	// fmt.Println("ClipPreserve === 4", i.Y, i.Height)
	// &Circle{pt, dx / 2}, nil
	// golang-draw传入的是图片的起点，取负值不知道为什么
	// pt := image.Pt(-int(i.X), -int(i.Y))

	// zero := image.Pt(0, 0)

	pt := image.Pt(-int(i.X), -int(i.Y))
	// painter := raster.NewAlphaOverPainter(clip)
	// dc.fill(painter)
	// if dc.mask == nil {
	// 	dc.mask = clip
	// } else {
	// mask := image.NewAlpha(image.Rect(0, 0, dc.width, dc.height))

	img, _ := ConvertBytesToImage(i.Buffer, i.MimeType)

	// fmt.Println("ClipPreserve === 5")
	// image.ZP
	draw.DrawMask(dst, dst.Bounds(), img, pt, tempPath, pt, draw.Over)
	// dc.mask = mask
	// }

	return dst, nil
}
