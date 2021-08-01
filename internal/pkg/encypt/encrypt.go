package encrpt

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

// 获取md5
func GetMd5FromString(data string) string {
	h := md5.New()
	io.WriteString(h, data)
	sum := fmt.Sprintf("%x", h.Sum(nil))
	return sum
}

// 获取sha256
func GetSHA256FromString(data string) string {
	h := sha256.New()
	io.WriteString(h, data)
	sum := fmt.Sprintf("%x", data)
	return sum
}

// 获取文件的md5
func GetMd5FromFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	if err != nil {
		return "", err
	}
	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil

}

// 获取文件的sha256
func GetSHA256FromFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	if err != nil {
		return "", err
	}
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	sum := fmt.Sprintf("%x", h.Sum(nil))
	return sum, nil
}
