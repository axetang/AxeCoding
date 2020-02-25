/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.IPConn
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type IPConn struct {
     // contains filtered or unexported fields
 }
 func DialIP(network string, laddr, raddr *IPAddr) (*IPConn, error)
 func ListenIP(network string, laddr *IPAddr) (*IPConn, error)
 func (c *IPConn) Close() error
 func (c *IPConn) File() (f *os.File, err error)
 func (c *IPConn) LocalAddr() Addr
 func (c *IPConn) Read(b []byte) (int, error)
 func (c *IPConn) ReadFrom(b []byte) (int, Addr, error)
 func (c *IPConn) ReadFromIP(b []byte) (int, *IPAddr, error)
 func (c *IPConn) ReadMsgIP(b, oob []byte) (n, oobn, flags int, addr *IPAddr, err error)
 func (c *IPConn) RemoteAddr() Addr
 func (c *IPConn) SetDeadline(t time.Time) error
 func (c *IPConn) SetReadBuffer(bytes int) error
 func (c *IPConn) SetReadDeadline(t time.Time) error
 func (c *IPConn) SetWriteBuffer(bytes int) error
 func (c *IPConn) SetWriteDeadline(t time.Time) error
 func (c *IPConn) SyscallConn() (syscall.RawConn, error)
 func (c *IPConn) Write(b []byte) (int, error)
 func (c *IPConn) WriteMsgIP(b, oob []byte, addr *IPAddr) (n, oobn int, err error)
 func (c *IPConn) WriteTo(b []byte, addr Addr) (int, error)
 func (c *IPConn) WriteToIP(b []byte, addr *IPAddr) (int, error)
 ------------------------------------------------------------------------------------
 **Description:
 IPConn is the implementation of the Conn and PacketConn interfaces for IP network connections.
 DialIP acts like Dial for IP networks.
 The network must be an IP network name; see func Dial for details.
 If laddr is nil, a local address is automatically chosen. If the IP field of raddr is nil or an unspecified IP address, the local system is assumed.
 ListenIP acts like ListenPacket for IP networks.
 The network must be an IP network name; see func Dial for details.
 If the IP field of laddr is nil or an unspecified IP address, ListenIP listens on all available IP addresses of the local system except multicast IP addresses.
 Close closes the connection.
 File returns a copy of the underlying os.File It is the caller's responsibility to close f when finished. Closing c does not affect f, and closing f does not affect c.
 The returned os.File's file descriptor is different from the connection's. Attempting to change properties of the original using this duplicate may or may not have the desired effect.
 LocalAddr returns the local network address. The Addr returned is shared by all invocations of LocalAddr, so do not modify it.
 Read implements the Conn Read method.
 ReadFrom implements the PacketConn ReadFrom method.
 ReadFromIP acts like ReadFrom but returns an IPAddr.
 ReadMsgIP reads a message from c, copying the payload into b and the associated out-of-band data into oob. It returns the number of bytes copied into b, the number of bytes copied into oob, the flags that were set on the message and the source address of the message.
 The packages golang.org/x/net/ipv4 and golang.org/x/net/ipv6 can be used to manipulate IP-level socket options in oob.
 RemoteAddr returns the remote network address. The Addr returned is shared by all invocations of RemoteAddr, so do not modify it.
 SetDeadline implements the Conn SetDeadline method.
 SetReadBuffer sets the size of the operating system's receive buffer associated with the connection.
 SetReadDeadline implements the Conn SetReadDeadline method.
 SetWriteBuffer sets the size of the operating system's transmit buffer associated with the connection.
 SetWriteDeadline implements the Conn SetWriteDeadline method.
 SyscallConn returns a raw network connection. This implements the syscall.Conn interface.
 Write implements the Conn Write method.
 WriteMsgIP writes a message to addr via c, copying the payload from b and the associated out-of-band data from oob. It returns the number of payload and out-of-band bytes written.
 The packages golang.org/x/net/ipv4 and golang.org/x/net/ipv6 can be used to manipulate IP-level socket options in oob.
 WriteTo implements the PacketConn WriteTo method.
 WriteToIP acts like WriteTo but takes an IPAddr.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
