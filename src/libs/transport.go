package libs

import (
	"net"
	"net/http"
	"time"
)

// 由于
// 1.http.Transport是连接池的管理单元
// 2.Transport更改Dial->DialContext, 为兼容保留Dial，但是DialContext优先级更高
// 3.beego的SetTimeout基于Dial属性检查
// 4.beego自动生成的Dial未设置keepAlive参数
// 对不同host单独设置Tranport，不在统一使用一个连接池

const (
	IdleConnTimeout     = 90   // keep-alive Idle回收时间，默认无线大，最好还是有个回收时间, 和keepAlive保持一致
	MaxIdleConns        = 5000 // 总连接数，默认无限大，实际单机三个进程，目前短连接高峰30k+
	MaxIdleConnsPerHost = 5000 // 每个host最大连接数，只能以调用量最大的微信为准
)

var QcodeTransport = &http.Transport{
	DialContext: (&net.Dialer{ // http（非https）连接, 如果要tcp保持连接 KeepAlive != 0
		Timeout:   10 * time.Second, // 优先级高于Dial参数（将被抛弃），会覆盖beego setTimeout，可以提个PR改进beego
		KeepAlive: 30 * time.Second,
	}).DialContext,
	MaxIdleConns:        MaxIdleConns,
	MaxIdleConnsPerHost: MaxIdleConnsPerHost,
	IdleConnTimeout:     IdleConnTimeout * time.Second,
}

var ImageTransport = &http.Transport{
	DialContext: (&net.Dialer{ // http（非https）连接, 如果要tcp保持连接 KeepAlive != 0
		Timeout:   5 * time.Second, // 优先级高于Dial参数（将被抛弃），会覆盖beego setTimeout，可以提个PR改进beego
		KeepAlive: 30 * time.Second,
	}).DialContext,
	MaxIdleConns:        MaxIdleConns,
	MaxIdleConnsPerHost: MaxIdleConnsPerHost,
	IdleConnTimeout:     IdleConnTimeout * time.Second,
}
