package libs

import (
	"fmt"
	"testing"
	// "github.com/luankefei/golang-canvas/src/canvas"
	// "github.com/luankefei/golang-canvas/src/canvas"
)

func TestLoadJSON(t *testing.T) {
	value := make([]string, 0)
	filepath := "../config/font.json"
	LoadConfigFromJSON(filepath, &value)

	fmt.Println("value:", value)
}
