package libs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// LoadConfigFromJSON 从文件读取json配置
func LoadConfigFromJSON(filepath string, value interface{}) interface{} {
	file, _ := os.Open(filepath)
	defer file.Close()

	// read our opened jsonFile as a byte array.
	bytes, _ := ioutil.ReadAll(file)
	fmt.Println("bytes", bytes)

	conf := value

	err := json.Unmarshal(bytes, &conf)
	// fmt.Println("conf:", conf[0].FileName)

	if err != nil {
		fmt.Println("Error:", err)
	}

	return value
}
