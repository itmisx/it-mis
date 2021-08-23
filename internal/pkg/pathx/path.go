package pathx

import (
	"os"
	"path"
	"runtime"
)

var RootPath string

func init() {
	_, RootPath, _, _ = runtime.Caller(0)
	RootPath = path.Dir(path.Dir(path.Dir(path.Dir(RootPath))))
	_ = RootPath
}

func Mkdir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0777)
		if err != nil {
			return err
		}
	}
	return nil
}
