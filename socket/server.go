package socket

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func connHandler(c net.Conn) {
	if c == nil {
		fmt.Println("连接失败")
	}
	_, err := c.Write([]byte("true")) // 配置是否开启angle
	if err != nil {
		fmt.Println("服务端发送失败", err)
	}
	buf := make([]byte, 4096)
	for {
		cnt, err := c.Read(buf)
		if cnt == 0 || err != nil {
			err := c.Close()
			if err != nil {
				fmt.Println(err)
			}
			break
		}
		inStr := strings.TrimSpace(string(buf[0:cnt]))
		res := getResult(inStr)
		_, err = c.Write([]byte("服务器端回复" + res + "\n"))
		if err != nil {
			fmt.Println("服务端回复失败", err)
		}
	}
}

func ServerSocket() {
	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("开启socket失败", err)
	}
	fmt.Println("已开启Server")
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("连接出错", err)
		}
		go connHandler(conn)
	}
}

func getResult(str string) string {
	rs := []rune(str)
	res := "0"
	a := 0
	b := 0
	flag := false
	alreadyAdd := false
	flaga := false
	flagb := false
	for i := 0; i < len(rs); i++ {
		if rs[i] == '+' {
			if !alreadyAdd {
				alreadyAdd = true
			} else {
				return res
			}
			if flaga == false {
				return res
			} else {
				flag = true
			}
		} else if rs[i] >= '0' && rs[i] <= '9' {
			if !flag {
				flaga = true
				a = a*10 + int(rs[i]-'0')
			} else {
				flagb = true
				b = b*10 + int(rs[i]-'0')
			}
		} else {
			return res
		}
	}
	if flaga && flagb {
		res = strconv.Itoa(a + b)
	}
	return res
}
