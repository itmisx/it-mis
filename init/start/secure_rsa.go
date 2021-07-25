package start

import (
	"errors"
	"mis/internal/pkg"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func GenSecureRSA() error {
	_, Dir, _, _ := runtime.Caller(0)
	pwd := filepath.Dir(Dir)
	outputPath := path.Join(pwd + "/../../secure-rsa")
	// 判断secre-rsa是否存在
	if _, err := os.Stat(path.Join(outputPath, "/private.pem")); !os.IsNotExist(err) {
		return errors.New("rsa file already exists")
	}
	if _, err := os.Stat(path.Join(outputPath, "/public.pem")); !os.IsNotExist(err) {
		return errors.New("rsa file already exists")
	}
	return pkg.GenRSA(1024, outputPath)
}
