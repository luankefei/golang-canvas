package canvas

import (
	"errors"
	"fmt"

	"github.com/tdewolff/canvas"
)

// 兜底图片方案的image_key在Image对象中传递
func (i *Image) fetch() error {
	//	span, ctx := libs.FuncTraceInstance.Start(ctx)
	//	defer libs.FuncTraceInstance.Stop(span)
	//	span.LogFields(log.String("image url", image.ImageUrl))
	var err error

	if i.ImageURL == "" {
		err = errors.New("image url can not be empty")
		// libs.Log().Info("fetch one image failed, error(%v)", err)
		return err
	}
	var img []byte
	var t string
	// if image.LocalBufferName != "" {
	// 	name := fmt.Sprintf("%s::%s", businessName, image.LocalBufferName)
	// 	img, t, err = b.getLocalBufferImage(name, image)
	// 	libs.Log().Info("trace_id(%s) get image(%+v) from local buffer, err(%+v)", traceId, image, err)
	// } else {

	fmt.Println("before loadImageByteFromRemote", i.Width, i.Height)
	img, t, err = loadImageByteFromRemote(i.ImageURL, i.Width, i.Height)
	// libs.Log().Info("trace_id(%s) get image(%+v) from remote, err(%+v)", traceId, image, err)
	// }

	i.Mime = t
	i.Buffer = img
	return nil
}

// Draw image
func (i *Image) Draw(c *canvas.Context) {
	fmt.Printf("%v image draw", &i)
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
