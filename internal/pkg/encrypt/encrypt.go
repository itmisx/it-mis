package encrypt

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

// Md5 获取md5
func Md5(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	sum := fmt.Sprintf("%x", h.Sum(nil))
	return sum
}

// Sha256 获取sha256
func Sha256(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	sum := fmt.Sprintf("%x", h.Sum(nil))
	return sum
}

// FileMd5 获取文件的md5
func FileMd5(path string) (string, error) {
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

// FileSha256 获取文件的sha256
func FileSha256(path string) (string, error) {
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
