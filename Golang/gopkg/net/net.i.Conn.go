/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.Conn
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type Conn interface {
     // Read reads data from the connection.
     // Read can be made to time out and return an Error with Timeout() == true
     // after a fixed time limit; see SetDeadline and SetReadDeadline.
     Read(b []byte) (n int, err error)
     // Write writes data to the connection.
     // Write can be made to time out and return an Error with Timeout() == true
     // after a fixed time limit; see SetDeadline and SetWriteDeadline.
     Write(b []byte) (n int, err error)
     // Close closes the connection.
     // Any blocked Read or Write operations will be unblocked and return errors.
     Close() error
     // LocalAddr returns the local network address.
     LocalAddr() Addr
     // RemoteAddr returns the remote network address.
     RemoteAddr() Addr
     // SetDeadline sets the read and write deadlines associated
     // with the connection. It is equivalent to calling both
     // SetReadDeadline and SetWriteDeadline.
     //
     // A deadline is an absolute time after which I/O operations
     // fail with a timeout (see type Error) instead of
     // blocking. The deadline applies to all future and pending
     // I/O, not just the immediately following call to Read or
     // Write. After a deadline has been exceeded, the connection
     // can be refreshed by setting a deadline in the future.
     //
     // An idle timeout can be implemented by repeatedly extending
     // the deadline after successful Read or Write calls.
     //
     // A zero value for t means I/O operations will not time out.
     SetDeadline(t time.Time) error
     // SetReadDeadline sets the deadline for future Read calls
     // and any currently-blocked Read call.
     // A zero value for t means Read will not time out.
     SetReadDeadline(t time.Time) error
     // SetWriteDeadline sets the deadline for future Write calls
     // and any currently-blocked Write call.
     // Even if write times out, it may return n > 0, indicating that
     // some of the data was successfully written.
     // A zero value for t means Write will not time out.
     SetWriteDeadline(t time.Time) error
 }
 Conn is a generic stream-oriented network connection.
 Multiple goroutines may invoke methods on a Conn simultaneously.
 ------------------------------------------------------------------------------------
 **Description:
 type Conn interface {
    // Read从连接中读取数据
    // Read方法可能会在超过某个固定时间限制后超时返回错误，该错误的Timeout()方法返回真
    Read(b []byte) (n int, err error)
    // Write从连接中写入数据
    // Write方法可能会在超过某个固定时间限制后超时返回错误，该错误的Timeout()方法返回真
    Write(b []byte) (n int, err error)
    // Close方法关闭该连接
    // 并会导致任何阻塞中的Read或Write方法不再阻塞并返回错误
    Close() error
    // 返回本地网络地址
    LocalAddr() Addr
    // 返回远端网络地址
    RemoteAddr() Addr
    // 设定该连接的读写deadline，等价于同时调用SetReadDeadline和SetWriteDeadline
    // deadline是一个绝对时间，超过该时间后I/O操作就会直接因超时失败返回而不会阻塞
    // deadline对之后的所有I/O操作都起效，而不仅仅是下一次的读或写操作
    // 参数t为零值表示不设置期限
    SetDeadline(t time.Time) error
    // 设定该连接的读操作deadline，参数t为零值表示不设置期限
    SetReadDeadline(t time.Time) error
    // 设定该连接的写操作deadline，参数t为零值表示不设置期限
    // 即使写入超时，返回值n也可能>0，说明成功写入了部分数据
    SetWriteDeadline(t time.Time) error
 }
 ------------------------------------------------------------------------------------
 **要点总结:
 0. Conn接口代表通用的面向流的网络连接, 多个线程可能会同时调用同一个Conn的方法;
*************************************************************************************/
package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "www.163.com:80")
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
