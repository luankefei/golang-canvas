package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/luankefei/golang-canvas/src/canvas"
)

// Setup initialize the configuration instance
func Setup() {
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}

// SetupFont 测试配置文件导入
func SetupFont() {
	file, _ := os.Open("./font.json")
	defer file.Close()

	// read our opened jsonFile as a byte array.
	bytes, _ := ioutil.ReadAll(file)
	fmt.Println("bytes", bytes)

	conf := make([]canvas.Font, 0)

	err := json.Unmarshal(bytes, &conf)
	fmt.Println("conf:", conf[0].FileName)

	// err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
