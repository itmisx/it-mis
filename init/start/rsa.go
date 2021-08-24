package start

import (
	"it-mis/internal/pkg/debugx"
	"it-mis/internal/pkg/pathx"
	"it-mis/internal/pkg/rsax"
	"os"
	"path"
)

// GenSecureRSA 生成安全RSA证书s
func GenSecureRSA() error {
	outputPath := path.Join(pathx.RootPath, "/assets/secure-rsa")
	// 判断secre-rsa是否存在
	if _, err := os.Stat(path.Join(outputPath, "/private.pem")); !os.IsNotExist(err) {
		debugx.PrintInfo("rsa file already exists")
		return nil
	}
	if _, err := os.Stat(path.Join(outputPath, "/public.pem")); !os.IsNotExist(err) {
		debugx.PrintInfo("rsa file already exists")
		return nil
	}
	return rsax.GenRSA(1024, outputPath)
}
