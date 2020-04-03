module github.com/luankefei/golang-canvas

go 1.13

require (
	github.com/disintegration/imaging v1.6.2
	github.com/gin-gonic/gin v1.6.2
	github.com/golang/protobuf v1.3.5 // indirect
	github.com/rs/zerolog v1.18.0
	github.com/tdewolff/canvas v0.0.0-20200402151523-385ae42fcf3d
	golang.org/x/sys v0.0.0-20200331124033-c3d80250170d // indirect
	golang.org/x/tools/gopls v0.3.4 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.29.1 // indirect
	routes v0.0.0
)

replace routes v0.0.0 => ./src/routes
