package canvas

import (
	"fmt"
)

// Draw image
func (i *Image) Draw() {
	fmt.Printf("%v image draw", &i)
}
