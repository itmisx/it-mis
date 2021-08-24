package datetime

import (
	"fmt"
	"testing"
)

func TestDate(*testing.T) {
	fmt.Println("默认时区:", Date("y-m-d h:i:s"))
	SetDefaultTimeZone(0)
	fmt.Println("0时区:", Date("y-m-d h:i:s"))
}
