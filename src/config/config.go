package config

import (
	"time"
)

// IsLocal 识别当前环境是否是本地环境
// TODO: IsLocal
func IsLocal() bool {
	return true
	// return getMacAddr() == "6c:96:cf:dd:18:4d"
}

// Setup initialize the configuration instance
func Setup() {
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}
