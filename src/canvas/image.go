package canvas

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"

	"github.com/oliamb/cutter"
	"github.com/rs/zerolog/log"
	"github.com/tdewolff/canvas"
)

// Draw image
func (i *Image) Draw(c *canvas.Context) {
	img, _ := ConvertBytesToImage(i.Buffer, i.MimeType)

	// TODO: 研究一下crop的mode参数copy机制，对性能会有帮助
	// 注意裁剪是基于原图尺寸的，例如二位码应该裁剪到430-370, starpos (30, 30)
	if i.Clip.Width > 0 && !i.Resize {
		croppedImg, _ := cutter.Crop(img, cutter.Config{
			Width:  i.Clip.Width,
			Height: i.Clip.Height,
			Anchor: image.Point{i.Clip.X, i.Clip.Y},
		})
		img = croppedImg
	}

	// 获取实际图片尺寸和传入参数之间的比例
	scale := (float64(img.Bounds().Dx()) / i.Width)

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
	img, t, err = loadImageByteFromRemote(i)
	// libs.Log().Info("trace_id(%s) get image(%+v) from remote, err(%+v)", traceId, image, err)
	// }

	i.MimeType = t
	i.Buffer = img
	return nil
}
