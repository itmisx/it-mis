package encrypt

import (
	"it-mis/internal/pkg/debugx"
	"testing"
)

func TestMd5(*testing.T) {
	debugx.PrintInfo(Md5("123"))
}

func TestSha256(*testing.T) {
	debugx.PrintInfo(Sha256("123"))
}

func TestFileMd5(*testing.T) {
	debugx.PrintInfo(FileMd5("test.txt"))
}

func TestFileSha256(*testing.T) {
	debugx.PrintInfo(FileSha256("test.txt"))
}
