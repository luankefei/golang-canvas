module github.com/luankefei/golang-canvas

go 1.13

require (
	github.com/astaxie/beego v1.9.3-0.20171218111859-f16688817aa4
	github.com/gin-gonic/gin v1.5.0
	github.com/golang/protobuf v1.3.3 // indirect
	github.com/tdewolff/canvas v0.0.0-20200112234221-d90d8eee919d
	golang.org/x/sys v0.0.0-20190921204832-2dccfee4fd3e // indirect
	routes v0.0.0
)

replace routes v0.0.0 => ./src/routes
