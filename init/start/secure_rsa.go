package start

import (
	"errors"
	"it-mis/internal/pkg/rsax"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

// GenSecureRSA 生成安全RSA证书s
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
	return rsax.GenRSA(1024, outputPath)
}
