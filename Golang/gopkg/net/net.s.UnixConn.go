/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.UnixConn
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type UnixConn struct {
     // contains filtered or unexported fields
 }
 UnixConn is an implementation of the Conn interface for connections to Unix domain sockets.
 func DialUnix(network string, laddr, raddr *UnixAddr) (*UnixConn, error)
 DialUnix acts like Dial for Unix networks.
 The network must be a Unix network name; see func Dial for details.
 If laddr is non-nil, it is used as the local address for the connection.
 func ListenUnixgram(network string, laddr *UnixAddr) (*UnixConn, error)
 ListenUnixgram acts like ListenPacket for Unix networks.
 The network must be "unixgram".
 func (c *UnixConn) Close() error
 Close closes the connection.
 func (c *UnixConn) CloseRead() error
 CloseRead shuts down the reading side of the Unix domain connection. Most callers should just use Close.
 func (c *UnixConn) CloseWrite() error
 CloseWrite shuts down the writing side of the Unix domain connection. Most callers should just use Close.
 func (c *UnixConn) File() (f *os.File, err error)
 File returns a copy of the underlying os.File It is the caller's responsibility to close f when finished. Closing c does not affect f, and closing f does not affect c.
 The returned os.File's file descriptor is different from the connection's. Attempting to change properties of the original using this duplicate may or may not have the desired effect.
 func (c *UnixConn) LocalAddr() Addr
 LocalAddr returns the local network address. The Addr returned is shared by all invocations of LocalAddr, so do not modify it.
 func (c *UnixConn) Read(b []byte) (int, error)
 Read implements the Conn Read method.
 func (c *UnixConn) ReadFrom(b []byte) (int, Addr, error)
 ReadFrom implements the PacketConn ReadFrom method.
 func (c *UnixConn) ReadFromUnix(b []byte) (int, *UnixAddr, error)
 ReadFromUnix acts like ReadFrom but returns a UnixAddr.
 func (c *UnixConn) ReadMsgUnix(b, oob []byte) (n, oobn, flags int, addr *UnixAddr, err error)
 ReadMsgUnix reads a message from c, copying the payload into b and the associated out-of-band data into oob. It returns the number of bytes copied into b, the number of bytes copied into oob, the flags that were set on the message and the source address of the message.
 Note that if len(b) == 0 and len(oob) > 0, this function will still read (and discard) 1 byte from the connection.
 func (c *UnixConn) RemoteAddr() Addr
 RemoteAddr returns the remote network address. The Addr returned is shared by all invocations of RemoteAddr, so do not modify it.
 func (c *UnixConn) SetDeadline(t time.Time) error
 SetDeadline implements the Conn SetDeadline method.
 func (c *UnixConn) SetReadBuffer(bytes int) error
 SetReadBuffer sets the size of the operating system's receive buffer associated with the connection.
 func (c *UnixConn) SetReadDeadline(t time.Time) error
 SetReadDeadline implements the Conn SetReadDeadline method.
 func (c *UnixConn) SetWriteBuffer(bytes int) error
 SetWriteBuffer sets the size of the operating system's transmit buffer associated with the connection.
 func (c *UnixConn) SetWriteDeadline(t time.Time) error
 SetWriteDeadline implements the Conn SetWriteDeadline method.
 func (c *UnixConn) SyscallConn() (syscall.RawConn, error)
 SyscallConn returns a raw network connection. This implements the syscall.Conn interface.
 func (c *UnixConn) Write(b []byte) (int, error)
 Write implements the Conn Write method.
 func (c *UnixConn) WriteMsgUnix(b, oob []byte, addr *UnixAddr) (n, oobn int, err error)
 WriteMsgUnix writes a message to addr via c, copying the payload from b and the associated out-of-band data from oob. It returns the number of payload and out-of-band bytes written.
 Note that if len(b) == 0 and len(oob) > 0, this function will still write 1 byte to the connection.
 func (c *UnixConn) WriteTo(b []byte, addr Addr) (int, error)
 WriteTo implements the PacketConn WriteTo method.
 func (c *UnixConn) WriteToUnix(b []byte, addr *UnixAddr) (int, error)
 WriteToUnix acts like WriteTo but takes a UnixAddr.
 ------------------------------------------------------------------------------------
 **Description:
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
