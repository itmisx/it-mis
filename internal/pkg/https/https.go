package https

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

// GetCertInfo 获取证书信息
func GetCertInfo() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	seedUrl := "https://www.meijingxiu.net"
	resp, err := client.Get(seedUrl)
	if err != nil {
		fmt.Println(seedUrl, " 请求失败")
	}
	defer resp.Body.Close()
	certInfo := resp.TLS.PeerCertificates[0]
	fmt.Println("过期时间:", certInfo.NotAfter)
	fmt.Println("组织信息:", certInfo.Subject)
}
