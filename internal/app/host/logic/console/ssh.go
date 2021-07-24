package console

import (
	"fmt"
	"log"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

type SSHCli struct {
	IP         string      //IP地址
	Username   string      //用户名
	Password   string      //密码
	Port       int         //端口号
	client     *ssh.Client //ssh客户端
	session    *ssh.Session
	LastResult string //最近一次Run的结果
	Channel    ssh.Channel
}

//创建命令行对象
//@param ip IP地址
//@param username 用户名
//@param password 密码
//@param port 端口号,默认22
func New(ip string, username string, password string, port ...int) *SSHCli {
	cli := new(SSHCli)
	cli.IP = ip
	cli.Username = username
	cli.Password = password
	if len(port) <= 0 {
		cli.Port = 22
	} else {
		cli.Port = port[0]
	}
	// 连接ssh服务端
	if err := cli.connect(); err != nil {
		return cli
	}
	// 创建会话
	var err error
	if cli.session, err = cli.client.NewSession(); err != nil {
		return cli
	}
	return cli
}

//连接
func (c *SSHCli) connect() error {
	config := ssh.ClientConfig{
		User: c.Username,
		Auth: []ssh.AuthMethod{ssh.Password(c.Password)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: 10 * time.Second,
	}
	addr := fmt.Sprintf("%s:%d", c.IP, c.Port)
	sshClient, err := ssh.Dial("tcp", addr, &config)
	if err != nil {
		return err
	}
	c.client = sshClient
	return nil
}

func (c *SSHCli) Run() {
	session := c.session
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // 禁用回显（0禁用，1启动）
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, //output speed = 14.4kbaud
	}
	if err := session.RequestPty("linux", 32, 160, modes); err != nil {
		log.Fatalf("request pty error: %s", err.Error())
	}
	if err := session.Shell(); err != nil {
		log.Fatalf("start shell error: %s", err.Error())
	}
}
