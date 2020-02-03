package canvas

// Drawer is general interface
type Drawer interface {
	Draw()
}

// GlobalConfig type
type GlobalConfig struct{}

// ImageClip is clip image config
type ImageClip struct {
	width, height, x, y int32
}

// Image config
type Image struct {
	x, y, width, height, opacity, borderRadius int32
	imageURL                                   string
	clip                                       ImageClip
}
