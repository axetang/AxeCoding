/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.PacketConn
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type PacketConn interface {
     // ReadFrom reads a packet from the connection,
     // copying the payload into p. It returns the number of
     // bytes copied into p and the return address that
     // was on the packet.
     // It returns the number of bytes read (0 <= n <= len(p))
     // and any error encountered. Callers should always process
     // the n > 0 bytes returned before considering the error err.
     // ReadFrom can be made to time out and return
     // an Error with Timeout() == true after a fixed time limit;
     // see SetDeadline and SetReadDeadline.
     ReadFrom(p []byte) (n int, addr Addr, err error)
     // WriteTo writes a packet with payload p to addr.
     // WriteTo can be made to time out and return
     // an Error with Timeout() == true after a fixed time limit;
     // see SetDeadline and SetWriteDeadline.
     // On packet-oriented connections, write timeouts are rare.
     WriteTo(p []byte, addr Addr) (n int, err error)
     // Close closes the connection.
     // Any blocked ReadFrom or WriteTo operations will be unblocked and return errors.
     Close() error
     // LocalAddr returns the local network address.
     LocalAddr() Addr
     // SetDeadline sets the read and write deadlines associated
     // with the connection. It is equivalent to calling both
     // SetReadDeadline and SetWriteDeadline.
     //
     // A deadline is an absolute time after which I/O operations
     // fail with a timeout (see type Error) instead of
     // blocking. The deadline applies to all future and pending
     // I/O, not just the immediately following call to ReadFrom or
     // WriteTo. After a deadline has been exceeded, the connection
     // can be refreshed by setting a deadline in the future.
     //
     // An idle timeout can be implemented by repeatedly extending
     // the deadline after successful ReadFrom or WriteTo calls.
     //
     // A zero value for t means I/O operations will not time out.
     SetDeadline(t time.Time) error
     // SetReadDeadline sets the deadline for future ReadFrom calls
     // and any currently-blocked ReadFrom call.
     // A zero value for t means ReadFrom will not time out.
     SetReadDeadline(t time.Time) error
     // SetWriteDeadline sets the deadline for future WriteTo calls
     // and any currently-blocked WriteTo call.
     // Even if write times out, it may return n > 0, indicating that
     // some of the data was successfully written.
     // A zero value for t means WriteTo will not time out.
     SetWriteDeadline(t time.Time) error
 }
 ------------------------------------------------------------------------------------
 **Description:
 type PacketConn interface {
    // ReadFrom方法从连接读取一个数据包，并将有效信息写入b
    // ReadFrom方法可能会在超过某个固定时间限制后超时返回错误，该错误的Timeout()方法返回真
    // 返回写入的字节数和该数据包的来源地址
    ReadFrom(b []byte) (n int, addr Addr, err error)
    // WriteTo方法将有效数据b写入一个数据包发送给addr
    // WriteTo方法可能会在超过某个固定时间限制后超时返回错误，该错误的Timeout()方法返回真
    // 在面向数据包的连接中，写入超时非常罕见
    WriteTo(b []byte, addr Addr) (n int, err error)
    // Close方法关闭该连接
    // 会导致任何阻塞中的ReadFrom或WriteTo方法不再阻塞并返回错误
    Close() error
    // 返回本地网络地址
    LocalAddr() Addr
    // 设定该连接的读写deadline
    SetDeadline(t time.Time) error
    // 设定该连接的读操作deadline，参数t为零值表示不设置期限
    // 如果时间到达deadline，读操作就会直接因超时失败返回而不会阻塞
    SetReadDeadline(t time.Time) error
    // 设定该连接的写操作deadline，参数t为零值表示不设置期限
    // 如果时间到达deadline，写操作就会直接因超时失败返回而不会阻塞
    // 即使写入超时，返回值n也可能>0，说明成功写入了部分数据
    SetWriteDeadline(t time.Time) error
 }
 ------------------------------------------------------------------------------------
 **要点总结:
 0. PacketConn接口代表通用的面向数据包的网络连接。多个线程可能会同时调用同一个Conn的方法。
*************************************************************************************/
package main

func main() {
}
