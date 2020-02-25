package libs

import (
	"github.com/tdewolff/canvas"
)

// LoadFont 从本地文件注册字体
func LoadFont(filepath string, name string, style canvas.FontStyle) {
	font := canvas.NewFontFamily(name)

	if err := font.LoadFontFile(filepath, style); err != nil {
		panic(err)
	}
}
