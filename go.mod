module github.com/luankefei/golang-canvas

go 1.13

require (
	github.com/disintegration/imaging v1.6.2
	github.com/gin-gonic/gin v1.6.2
	github.com/golang/protobuf v1.3.5 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/rs/zerolog v1.18.0
	github.com/tdewolff/canvas v0.0.0-00010101000000-000000000000
	// github.com/oliverpool/canvas v0.0.0-a81c43fb9787f2330ab4775028a08c0a6d0215d5
	golang.org/x/sys v0.0.0-20200331124033-c3d80250170d // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	routes v0.0.0

)

// replace github.com/tdewolff/canvas => github.com/oliverpool/canvas@fix_render_image
replace routes v0.0.0 => ./src/routes

// replace github.com/tdewolff/canvas => github.com/oliverpool/canvas v0.0.0-20200414091838-a81c43fb9787
replace github.com/tdewolff/canvas => github.com/oliverpool/canvas v0.0.0-20200414143817-2158c3b733b4

// replace github.com/tdewolff/canvas => github.com/oliverpool/canvas fix_render_image
