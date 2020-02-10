# golang-canvas
canvas api demo

### 目录

```
src
	canvas              绘图主业务逻辑
	config              配置文件
	libs                字体加载，数据库连接，阿里云等
	main                项目入口
	middleware          web框架相关中间件
	protos              protobuf，预留
	routes              api路由
	services            api相关数据逻辑
	static              静态资源
	test                demo
	views               页面模板
	

go.mod
.pre-commit-config.yaml, .pre-commit-hooks.yaml
```


### Testing
Unit tests:

```
$ go test -v
```


### 元素类型

此处设计会尽量向前兼容，靠近canvas-api

```

// clip字段决定图片是否被裁剪成圆角
// transparent 字段会将图片信息重写，抹去灰度过低的颜色
Image {
  x: number
  y: number
  height: number
  width: number
  imageUrl: string
  clip?: boolean
  transparent?: boolean
}

// lineHeight 和 limit 主要用于文字多行需要计算折行的情况
// 如果文字是居中对齐，x和y值需要传入水平居中的中心点坐标
// EAlign: left | center | right | justify
// EFontWeight: bold | regular | normal
IText {
  x: number
  y: number
  content: string
  color: string
  size: number
  align?: EAlign
  fontWeight?: EFontWeight
  lineHeight?: number
  limit?: number
}

ILine {
  lineWidth: number
  color: string
  x1: number
  y1: number
  x2: number
  y2: number
}
```


### TODO

- 规范字体命名
- 规范log，全链路监控
- 测试用例
- 消息队列
- 定时重启
- 单机压测数据
- modern properties(shadow round opacity mask transform)

### Format of the commit message
- feat (feature)
- fix (bug fix)
- docs (documentation)
- style (formatting, missing semi colons, …)
- refactor
- test (when adding missing tests)
- chore (maintain)

### 依赖方案

+ [gin](https://github.com/gin-gonic/gin)
+ [kafka](https://github.com/NervJS/taro)
+ [canvas](https://github.com/tdewolff/canvas)


