package config

import (
	"time"
)

// Server 定义
type Server struct {
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// ServerSetting 单例对象
var ServerSetting = &Server{}
