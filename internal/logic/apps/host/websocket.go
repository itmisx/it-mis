package host

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Console struct {
	SSHClient *SSHCli
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (cs Console) Connect(c *gin.Context) {
	// 升级get请求为webSocket协议
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	// 初始化ssh客户端
	cs.SSHClient = New("", "", "", 22)
	defer conn.Close()
	w, _ := cs.SSHClient.session.StdinPipe()
	r, _ := cs.SSHClient.session.StdoutPipe()
	cs.SSHClient.Run()

	go func() {
		for {
			var buf [65 * 10000]byte
			n, _ := r.Read(buf[0:])
			if n > 0 {
				conn.WriteMessage(1, buf[:n])
			}
		}
	}()
	for {
		// 读取conn中的数据
		_, message, err := conn.ReadMessage()
		if err != nil {
			break
		}
		_, err1 := w.Write(message)
		if err1 != nil {
			fmt.Println(err1)
		}
	}
}
