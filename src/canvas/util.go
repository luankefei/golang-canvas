package canvas

import (
	"encoding/hex"
	"fmt"
	"image/color"
	"io/ioutil"

	"strconv"
	"strings"

	"github.com/luankefei/golang-canvas/src/config"
	"github.com/luankefei/golang-canvas/src/request"
	"github.com/rs/zerolog/log"
)

// urlIsOss 检查图片url是否是阿里云oss地址
func urlIsOss(imgURL string) bool {
	return strings.Contains(imgURL, "oss") || strings.Contains(imgURL, "img.laiye.com")
}

// reformatUrl 将oss地址替换为内网地址，用oss实时缩图
func reformatURL(imageURL string, w float64, h float64) (string, error) {
	fmt.Println("reformatURL", w, h)
	if config.IsLocal() == false {
		// 替换为内网地址
		imageURL = strings.Replace(imageURL, "img.laiye.com", "laiye-image.oss-cn-beijing-internal.aliyuncs.com", 1)
		imageURL = strings.Replace(imageURL, "oss-cn-beijing.aliyuncs.com", "oss-cn-beijing-internal.aliyuncs.com", 1)
	}
	if w == float64(0) && h == float64(0) && urlIsOss(imageURL) {
		return "", fmt.Errorf("image(%s) not in oss can not assign w, h", imageURL)
	} else if w != float64(0) && h != float64(0) && urlIsOss(imageURL) {
		// oss实时缩图，注意oss不支持float，所以这里要强转
		imageURL += fmt.Sprintf("?x-oss-process=image/resize,m_lfit,h_%d,w_%d", uint32(h), uint32(w))
	}

	fmt.Println("reformatURL", imageURL)

	return imageURL, nil
}

// LoadImageByteFromRemote 从远程获取图片的字节流
func loadImageByteFromRemote(imgURL string, w float64, h float64) ([]byte, string, error) {
	img := []byte{}
	var t string
	imgURL, err := reformatURL(imgURL, w, h)
	if err != nil {
		return img, t, err
	}
	response, err := request.
		Get(imgURL).
		SetTransport(config.ImageTransport).
		Response()
	if response != nil && response.Body != nil {
		defer response.Body.Close()
	}
	if err != nil {
		log.Info().
			Err(err).
			Msgf("load image(%s) from remote failed", imgURL)
		return img, t, err
	}
	img, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Info().
			Msgf("read all from image(%s) response failed, error(%+v)", imgURL, err)
		return img, t, err
	}
	headers := response.Header
	t = headers["Content-Type"][0]
	fmt.Println("loadImageRemoteFromByte", headers["Content-Type"])
	return img, t, nil
}

// fetchOneImage 获取一个图片的字节流和类型
//  (b *FetchPic)

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
			fmt.Println("default_type", t)
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

	colorStr, err := normalize(colorStr)
	if err != nil {
		log.Fatal().
			Err(err)
	}

	b, err := hex.DecodeString(colorStr)
	if err != nil {
		log.Fatal().
			Err(err)
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
