module github.com/luankefei/golang-canvas

go 1.13

require (
	github.com/disintegration/imaging v1.6.2
	github.com/gin-gonic/gin v1.6.2
	github.com/golang/protobuf v1.3.5 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/oliamb/cutter v0.2.2
	github.com/rs/zerolog v1.18.0
	github.com/tdewolff/canvas v0.0.0-20200402151523-385ae42fcf3d
	golang.org/x/sys v0.0.0-20200331124033-c3d80250170d // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	routes v0.0.0
)

replace routes v0.0.0 => ./src/routes
