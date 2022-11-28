package socket

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func cConnHandler(c *net.TCPConn) {
	reader := bufio.NewReader(os.Stdin)
	buf := make([]byte, 1024)

	cnt, err := c.Read(buf)
	flag := string(buf[0:cnt])
	if flag == "true" {
		err = c.SetNoDelay(true)
	} else if flag == "false" {
		err = c.SetNoDelay(false)
	}
	if err != nil {
		fmt.Println("客户端获取配置失败", err)
	}
	for {
		fmt.Println("输入数据:")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		_, err := c.Write([]byte(input))
		if err != nil {
			fmt.Println("发送失败", err)
		}
		cnt, err := c.Read(buf)
		if err != nil {
			fmt.Printf("客户端读取数据失败", err)
			continue
		}
		fmt.Print("服务器端回复" + string(buf[0:cnt]))
	}
}

func ClientSocket() {
	target := "localhost:8080"
	raddr, err := net.ResolveTCPAddr("tcp", target)
	if err != nil {
		fmt.Println("连接失败", err)
	}
	conn, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		fmt.Println("连接失败", err)
	}
	cConnHandler(conn)
}
