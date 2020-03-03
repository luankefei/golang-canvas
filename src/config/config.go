package config

import (
	"time"
)

// Setup initialize the configuration instance
func Setup() {
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}

// // SetupFont 测试配置文件导入
// func SetupFont() {
// 	file, _ := os.Open("./font.json")
// 	defer file.Close()

// 	// read our opened jsonFile as a byte array.
// 	bytes, _ := ioutil.ReadAll(file)
// 	fmt.Println("bytes", bytes)

// 	conf := make([]canvas.Font, 0)

// 	err := json.Unmarshal(bytes, &conf)
// 	fmt.Println("conf:", conf[0].FileName)

// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}
// }
