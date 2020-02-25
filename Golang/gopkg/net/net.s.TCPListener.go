/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.TCPListener
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type TCPListener struct {
     // contains filtered or unexported fields
 }
 TCPListener is a TCP network listener. Clients should typically use variables of type Listener instead of assuming TCP.
 func ListenTCP(network string, laddr *TCPAddr) (*TCPListener, error)
 ListenTCP acts like Listen for TCP networks.
 The network must be a TCP network name; see func Dial for details.
 If the IP field of laddr is nil or an unspecified IP address, ListenTCP listens on all available unicast and anycast IP addresses of the local system. If the Port field of laddr is 0, a port number is automatically chosen.
 func (l *TCPListener) Accept() (Conn, error)
 Accept implements the Accept method in the Listener interface; it waits for the next call and returns a generic Conn.
 func (l *TCPListener) AcceptTCP() (*TCPConn, error)
 AcceptTCP accepts the next incoming call and returns the new connection.
 func (l *TCPListener) Addr() Addr
 Addr returns the listener's network address, a *TCPAddr. The Addr returned is shared by all invocations of Addr, so do not modify it.
 func (l *TCPListener) Close() error
 Close stops listening on the TCP address. Already Accepted connections are not closed.
 func (l *TCPListener) File() (f *os.File, err error)
 File returns a copy of the underlying os.File. It is the caller's responsibility to close f when finished. Closing l does not affect f, and closing f does not affect l.
 The returned os.File's file descriptor is different from the connection's. Attempting to change properties of the original using this duplicate may or may not have the desired effect.
 func (l *TCPListener) SetDeadline(t time.Time) error
 SetDeadline sets the deadline associated with the listener. A zero time value disables the deadline.
 func (l *TCPListener) SyscallConn() (syscall.RawConn, error)
 SyscallConn returns a raw network connection. This implements the syscall.Conn interface.
 The returned RawConn only supports calling Control. Read and Write return an error.
 ------------------------------------------------------------------------------------
 **Description:
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
