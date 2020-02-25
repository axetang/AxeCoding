/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.UDPConn
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type UDPConn struct {
     // contains filtered or unexported fields
 }
 UDPConn is the implementation of the Conn and PacketConn interfaces for UDP network connections.
 func DialUDP(network string, laddr, raddr *UDPAddr) (*UDPConn, error)
 DialUDP acts like Dial for UDP networks.
 The network must be a UDP network name; see func Dial for details.
 If laddr is nil, a local address is automatically chosen. If the IP field of raddr is nil or an unspecified IP address, the local system is assumed.
 func ListenMulticastUDP(network string, ifi *Interface, gaddr *UDPAddr) (*UDPConn, error)
 ListenMulticastUDP acts like ListenPacket for UDP networks but takes a group address on a specific network interface.
 The network must be a UDP network name; see func Dial for details.
 ListenMulticastUDP listens on all available IP addresses of the local system including the group, multicast IP address. If ifi is nil, ListenMulticastUDP uses the system-assigned multicast interface, although this is not recommended because the assignment depends on platforms and sometimes it might require routing configuration. If the Port field of gaddr is 0, a port number is automatically chosen.
 ListenMulticastUDP is just for convenience of simple, small applications. There are golang.org/x/net/ipv4 and golang.org/x/net/ipv6 packages for general purpose uses.
 func ListenUDP(network string, laddr *UDPAddr) (*UDPConn, error)
 ListenUDP acts like ListenPacket for UDP networks.
 The network must be a UDP network name; see func Dial for details.
 If the IP field of laddr is nil or an unspecified IP address, ListenUDP listens on all available IP addresses of the local system except multicast IP addresses. If the Port field of laddr is 0, a port number is automatically chosen.
 func (c *UDPConn) Close() error
 Close closes the connection.
 func (c *UDPConn) File() (f *os.File, err error)
 File returns a copy of the underlying os.File It is the caller's responsibility to close f when finished. Closing c does not affect f, and closing f does not affect c.
 The returned os.File's file descriptor is different from the connection's. Attempting to change properties of the original using this duplicate may or may not have the desired effect.
 func (c *UDPConn) LocalAddr() Addr
 LocalAddr returns the local network address. The Addr returned is shared by all invocations of LocalAddr, so do not modify it.
 func (c *UDPConn) Read(b []byte) (int, error)
 Read implements the Conn Read method.
 func (c *UDPConn) ReadFrom(b []byte) (int, Addr, error)
 ReadFrom implements the PacketConn ReadFrom method.
 func (c *UDPConn) ReadFromUDP(b []byte) (int, *UDPAddr, error)
 ReadFromUDP acts like ReadFrom but returns a UDPAddr.
 func (c *UDPConn) ReadMsgUDP(b, oob []byte) (n, oobn, flags int, addr *UDPAddr, err error)
 ReadMsgUDP reads a message from c, copying the payload into b and the associated out-of-band data into oob. It returns the number of bytes copied into b, the number of bytes copied into oob, the flags that were set on the message and the source address of the message.
 The packages golang.org/x/net/ipv4 and golang.org/x/net/ipv6 can be used to manipulate IP-level socket options in oob.
 func (c *UDPConn) RemoteAddr() Addr
 RemoteAddr returns the remote network address. The Addr returned is shared by all invocations of RemoteAddr, so do not modify it.
 func (c *UDPConn) SetDeadline(t time.Time) error
 SetDeadline implements the Conn SetDeadline method.
 func (c *UDPConn) SetReadBuffer(bytes int) error
 SetReadBuffer sets the size of the operating system's receive buffer associated with the connection.
 func (c *UDPConn) SetReadDeadline(t time.Time) error
 SetReadDeadline implements the Conn SetReadDeadline method.
 func (c *UDPConn) SetWriteBuffer(bytes int) error
 SetWriteBuffer sets the size of the operating system's transmit buffer associated with the connection.
 func (c *UDPConn) SetWriteDeadline(t time.Time) error
 SetWriteDeadline implements the Conn SetWriteDeadline method.
 func (c *UDPConn) SyscallConn() (syscall.RawConn, error)
 SyscallConn returns a raw network connection. This implements the syscall.Conn interface.
 func (c *UDPConn) Write(b []byte) (int, error)
 Write implements the Conn Write method.
 func (c *UDPConn) WriteMsgUDP(b, oob []byte, addr *UDPAddr) (n, oobn int, err error)
 WriteMsgUDP writes a message to addr via c if c isn't connected, or to c's remote address if c is connected (in which case addr must be nil). The payload is copied from b and the associated out-of-band data is copied from oob. It returns the number of payload and out-of-band bytes written.
 The packages golang.org/x/net/ipv4 and golang.org/x/net/ipv6 can be used to manipulate IP-level socket options in oob.
 func (c *UDPConn) WriteTo(b []byte, addr Addr) (int, error)
 WriteTo implements the PacketConn WriteTo method.
 func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error)
 WriteToUDP acts like WriteTo but takes a UDPAddr.
 ------------------------------------------------------------------------------------
 **Description:
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
