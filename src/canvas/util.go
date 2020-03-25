package canvas

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/luankefei/golang-canvas/src/config"
	"github.com/luankefei/golang-canvas/src/libs"
)

// urlIsOss 检查图片url是否是阿里云oss地址
func urlIsOss(imgURL string) bool {
	return strings.Contains(imgURL, "oss")
}

// reformatUrl 将oss地址替换为内网地址，用oss实时缩图
func reformatURL(imageURL string, w uint32, h uint32) (string, error) {
	fmt.Println("reformatURL", w, h)
	if config.IsLocal() == false {
		// 替换为内网地址
		imageURL = strings.Replace(imageURL, "img.laiye.com", "laiye-image.oss-cn-beijing-internal.aliyuncs.com", 1)
		imageURL = strings.Replace(imageURL, "oss-cn-beijing.aliyuncs.com", "oss-cn-beijing-internal.aliyuncs.com", 1)
	}
	if w == uint32(0) && h == uint32(0) && urlIsOss(imageURL) {
		return "", fmt.Errorf("image(%s) not in oss can not assign x, y", imageURL)
	} else if w != uint32(0) && h != uint32(0) && urlIsOss(imageURL) {
		// oss实时缩图
		imageURL += fmt.Sprintf("?x-oss-process=image/resize,m_lfit,h_%d,w_%d", h, w)
	}
	return imageURL, nil
}

// LoadImageByteFromRemote 从远程获取图片的字节流
func loadImageByteFromRemote(imgURL string, w uint32, h uint32) ([]byte, string, error) {
	img := []byte{}
	var t string
	imgURL, err := reformatURL(imgURL, w, h)
	if err != nil {
		return img, t, err
	}
	response, err := libs.
		Get(imgURL).
		// SetTransport(ImageTransport).
		Response()
	if response != nil && response.Body != nil {
		defer response.Body.Close()
	}
	if err != nil {
		// Log().Info("load image(%s) from remote failed, error(%+v)", imgURL, err)
		fmt.Println("error1")
		return img, t, err
	}
	img, err = ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error2")
		// Log().Info("read all from image(%s) response failed, error(%+v)", imgURL, err)
		return img, t, err
	}
	headers := response.Header
	contentTypes := headers["Content-Type"]
	fmt.Println("contentTypes", contentTypes)
	// for _, c := range contentTypes {
	// 	if c == "image/png" {
	// 		t = draw_pb.ImageType_png
	// 		break
	// 	} else if c == "image/jpeg" || c == "image/jpg" {
	// 		t = draw_pb.ImageType_jpg
	// 		break
	// 	}
	// }
	return img, t, nil
}

// fetchOneImage 获取一个图片的字节流和类型
//  (b *FetchPic)

// 兜底图片方案的image_key在Image对象中传递
func fetchOneImage(ctx context.Context, image *Image) error {

	//	span, ctx := libs.FuncTraceInstance.Start(ctx)
	//	defer libs.FuncTraceInstance.Stop(span)
	//	span.LogFields(log.String("image url", image.ImageUrl))
	var err error

	if image.ImageURL == "" {
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

	fmt.Println("before loadImageByteFromRemote", image.Width, image.Height)
	img, t, err = loadImageByteFromRemote(image.ImageURL, image.Width, image.Height)
	// libs.Log().Info("trace_id(%s) get image(%+v) from remote, err(%+v)", traceId, image, err)
	// }

	image.Mime = t
	image.Buffer = img
	return nil
}

// LoadImageFilter drawer
// TODO: 这里后续可以加入buffer_name和expires的检查
func LoadImageFilter(arr []interface{}) []*Image {
	d := make([]*Image, 0)

	for _, v := range arr {
		switch t := v.(type) {
		case *Image:
			d = append(d, v.(*Image))
			break
		default:
			fmt.Println("default type", t)
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
