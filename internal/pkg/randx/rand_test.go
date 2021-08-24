package randx

import (
	"fmt"
	"it-mis/internal/pkg/debugx"
	"os"
	"testing"
)

func TestRandString(*testing.T) {
	debugx.PrintInfo(RandString(20))
}

func TestTempDir(*testing.T) {
	fmt.Println(os.TempDir())
}
