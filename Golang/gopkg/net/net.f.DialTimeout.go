/************************************************************************************
**Author:  Axe Tang; Email: axetang@163.com
**Package: net
**Element: net.DialTimeout
**Type: func
------------------------------------------------------------------------------------
**Definition:
func DialTimeout(network, address string, timeout time.Duration) (Conn, error)
------------------------------------------------------------------------------------
**Description:
DialTimeout acts like Dial but takes a timeout.
The timeout includes name resolution, if required. When using TCP, and the host in the address parameter resolves to multiple IP addresses, the timeout is spread over each consecutive dial, such that each is given an appropriate fraction of the time to connect.
See func Dial for a description of the network and address parameters.
------------------------------------------------------------------------------------
**要点总结:
1. DialTimeout类似Dial但采用了超时。timeout参数如果必要可包含名称解析。
*************************************************************************************/
package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.DialTimeout("tcp", "www.163.com:80", time.Second)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", conn) //返回类似 &net.TCPConn{fd:(*net.netFD)(0xf840069000)}
	fmt.Println(conn.LocalAddr())
	fmt.Println(conn.RemoteAddr())
	n, err := conn.Write([]byte("This is msg from axe!"))
	fmt.Println(n, err)
	buf := make([]byte, 10)
	conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	n, err = conn.Read(buf)
	fmt.Println(buf)
}
