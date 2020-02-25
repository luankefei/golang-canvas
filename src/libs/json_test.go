package libs

import (
	"fmt"
	"testing"

	"github.com/luankefei/golang-canvas/src/canvas"
)

func TestLoadJSON(t *testing.T) {
	value := make([]canvas.Font, 0)
	filepath := "../config/font.json"
	LoadConfigFromJSON(filepath, &value)

	fmt.Println("value:", value)
}
