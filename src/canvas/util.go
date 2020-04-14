package canvas

import (
	"encoding/hex"
	"fmt"
	"image/color"
	"io/ioutil"
	"net/http"

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
func reformatURL(imageURL string, r float64, resize bool, w float64, h float64) (string, error) {
	if config.IsLocal() == false {
		// 替换为内网地址
		imageURL = strings.Replace(imageURL, "img.laiye.com", "laiye-image.oss-cn-beijing-internal.aliyuncs.com", 1)
		imageURL = strings.Replace(imageURL, "oss-cn-beijing.aliyuncs.com", "oss-cn-beijing-internal.aliyuncs.com", 1)
	}

	// early return
	if !urlIsOss(imageURL) || w == float64(0) || h == float64(0) {
		return imageURL, nil
	}

	// 需要调用oss接口对图片做处理
	if r > 0 || resize {
		imageURL += "?x-oss-process=image"

		// TODO: 如果使用oss resize，二维码会变形，原因未知
		// @see https://help.aliyun.com/document_detail/44688.html?spm=a2c4g.11186623.6.1367.57f117f1UQGsW3
		if resize {
			// 调用 resize，默认是不允许放大。即如果请求的图片比原图大，那么返回的仍然是原图。如果想取到放大的图片，即增加参数调用 limit_0
			// oss实时缩图，注意oss不支持float，所以这里要强转
			imageURL += fmt.Sprintf("/resize,m_lfit,h_%d,w_%d,limit_0", uint32(h), uint32(w))
		}

		// 图片切圆角
		// @see https://help.aliyun.com/document_detail/44694.html?spm=a2c4g.11186623.6.1372.5959c1f6Sz0KTS
		if r > 0 {
			imageURL += fmt.Sprintf("/rounded-corners,r_%d/format,png", uint32(r))
		}
	}

	fmt.Println("zzzz: ", imageURL)
	return imageURL, nil
}

func getFileContentType(buffer []byte) (string, error) {
	// Only the first 512 bytes are used to sniff the content type.
	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer[:512])

	return contentType, nil
}

// LoadImageByteFromRemote 从远程获取图片的字节流
func loadImageByteFromRemote(i *Image) ([]byte, string, error) {
	img := []byte{}
	var t string

	// TODO: 使用阿里云oss处理图片 如果使用oss拉取图片，就变形了
	imgURL, err := reformatURL(i.ImageURL, i.BorderRadius, i.Resize, i.Width, i.Height)
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

	t, err = getFileContentType(img)

	// http content-type field is untructed，this only be used for fallback
	if err != nil {
		t = headers["Content-Type"][0]
	}
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
