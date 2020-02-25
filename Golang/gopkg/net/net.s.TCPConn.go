/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.TCPConn
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type TCPConn struct {
     // contains filtered or unexported fields
	conn
 }
 func (c *TCPConn) Close() error
 func (c *TCPConn) CloseRead() error
 func (c *TCPConn) CloseWrite() error
 func (c *TCPConn) File() (f *os.File, err error)
 func (c *TCPConn) LocalAddr() Addr
 func (c *TCPConn) Read(b []byte) (int, error)
 func (c *TCPConn) ReadFrom(r io.Reader) (int64, error)
 func (c *TCPConn) RemoteAddr() Addr
 func (c *TCPConn) SetDeadline(t time.Time) error
 func (c *TCPConn) SetKeepAlive(keepalive bool) error
 func (c *TCPConn) SetKeepAlivePeriod(d time.Duration) error
 func (c *TCPConn) SetLinger(sec int) error
 func (c *TCPConn) SetNoDelay(noDelay bool) error
 func (c *TCPConn) SetReadBuffer(bytes int) error
 func (c *TCPConn) SetReadDeadline(t time.Time) error
 func (c *TCPConn) SetWriteBuffer(bytes int) error
 func (c *TCPConn) SetWriteDeadline(t time.Time) error
 func (c *TCPConn) SyscallConn() (syscall.RawConn, error)
 func (c *TCPConn) Write(b []byte) (int, error)
 ------------------------------------------------------------------------------------
 **Description:
 TCPConn is an implementation of the Conn interface for TCP network connections.
 Close closes the connection.
 CloseRead shuts down the reading side of the TCP connection. Most callers should
 just use Close.
 CloseWrite shuts down the writing side of the TCP connection. Most callers should
 just use Close.
 File returns a copy of the underlying os.File It is the caller's responsibility to
 close f when finished. Closing c does not affect f, and closing f does not affect c.
 The returned os.File's file descriptor is different from the connection's. Attempting
 to change properties of the original using this duplicate may or may not have the
 desired effect.
 LocalAddr returns the local network address. The Addr returned is shared by all
 invocations of LocalAddr, so do not modify it.
 Read implements the Conn Read method.
 ReadFrom implements the io.ReaderFrom ReadFrom method.
 RemoteAddr returns the remote network address. The Addr returned is shared by all
 invocations of RemoteAddr, so do not modify it.
 SetDeadline implements the Conn SetDeadline method.
 SetKeepAlive sets whether the operating system should send keepalive messages on
 the connection.
 SetKeepAlivePeriod sets period between keep alives.
 SetLinger sets the behavior of Close on a connection which still has data waiting
 to be sent or to be acknowledged.
 If sec < 0 (the default), the operating system finishes sending the data in the
 background.
 If sec == 0, the operating system discards any unsent or unacknowledged data.
 If sec > 0, the data is sent in the background as with sec < 0. On some operating
 systems after sec seconds have elapsed any remaining unsent data may be discarded.
 SetNoDelay controls whether the operating system should delay packet transmission
 in hopes of sending fewer packets (Nagle's algorithm). The default is true (no
 delay), meaning that data is sent as soon as possible after a Write.
 SetReadBuffer sets the size of the operating system's receive buffer associated
 with the connection.
 SetReadDeadline implements the Conn SetReadDeadline method.
 SetWriteBuffer sets the size of the operating system's transmit buffer associated
 with the connection.
 SetWriteDeadline implements the Conn SetWriteDeadline method.
 SyscallConn returns a raw network connection. This implements the syscall.Conn
 interface.
 Write implements the Conn Write method.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. TCPConn和IPConn等一致，都是对内部结构体conn的一个封装，内部结构体conn实现了Conn接口：
 type TCPConn struct {
     // contains filtered or unexported fields
	conn
 }
 type conn struct {
	fd *netFD
 }
// Network file descriptor.
type netFD struct {
	pfd poll.FD

	// immutable until Close
	family      int
	sotype      int
	isConnected bool // handshake completed or use of association with peer
	net         string
	laddr       Addr
	raddr       Addr
}
1.
*************************************************************************************/
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var logger = log.New(os.Stdout, "AxeLogger: ", log.Lshortfile)

func main() {
	go listenTCP()

	dialInTCP()
	time.Sleep(10 * time.Second)
}

func dialInTCP() {
	logger.Print("entered into dailInTCP")
	conn, err := net.Dial("tcp", "127.0.0.1:8080") //开启TCP连接端口
	if err != nil {
		log.Fatal(err)
	}
	logger.Print("Success in net.Dial 127.0.0.1:8080")
	defer conn.Close()
	msg := []byte("This is message from dialInTCP!")
	if tconn, ok := conn.(*net.TCPConn); ok { //这里可以断言，因为是tcp连接，其实值类型就是TCPConn
		fmt.Println("assert success")
		tconn.CloseRead() //尝试调用TCPConn的函数
		_, err = tconn.Write(msg)
		logger.Print("Write to conn with the message by dialInTCP")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		_, err = conn.Write(msg)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func listenTCP() {
	logger.Print("entered into listenTCP")
	l, err := net.Listen("tcp", ":8080") //监听TCP，8080端口
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept() //接受连接，当上面Dial的时候，这就会接受连接
		if err != nil {
			log.Fatal(err)
		}
		logger.Print("Success in accept the connecting application from dialInTCP")
		go func(c net.Conn) { //使用goroutine，并发
			buf := make([]byte, 256) //接受字段
			_, err = c.Read(buf)     //读取内容
			logger.Print("Read the msg with conn from dialInTCP")
			logger.Print(string(buf))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(buf))
			c.Close()
		}(conn)
	}
}
