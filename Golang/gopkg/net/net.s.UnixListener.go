/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.UnixListener
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type UnixListener struct {
     // contains filtered or unexported fields
 }
 UnixListener is a Unix domain socket listener. Clients should typically use variables of type Listener instead of assuming Unix domain sockets.
 func ListenUnix(network string, laddr *UnixAddr) (*UnixListener, error)
 ListenUnix acts like Listen for Unix networks.
 The network must be "unix" or "unixpacket".
 func (l *UnixListener) Accept() (Conn, error)
 Accept implements the Accept method in the Listener interface. Returned connections will be of type *UnixConn.
 func (l *UnixListener) AcceptUnix() (*UnixConn, error)
 AcceptUnix accepts the next incoming call and returns the new connection.
 func (l *UnixListener) Addr() Addr
 Addr returns the listener's network address. The Addr returned is shared by all invocations of Addr, so do not modify it.
 func (l *UnixListener) Close() error
 Close stops listening on the Unix address. Already accepted connections are not closed.
 func (l *UnixListener) File() (f *os.File, err error)
 File returns a copy of the underlying os.File. It is the caller's responsibility to close f when finished. Closing l does not affect f, and closing f does not affect l.
 The returned os.File's file descriptor is different from the connection's. Attempting to change properties of the original using this duplicate may or may not have the desired effect.
 func (l *UnixListener) SetDeadline(t time.Time) error
 SetDeadline sets the deadline associated with the listener. A zero time value disables the deadline.
 func (l *UnixListener) SetUnlinkOnClose(unlink bool)
 SetUnlinkOnClose sets whether the underlying socket file should be removed from the file system when the listener is closed.
 The default behavior is to unlink the socket file only when package net created it. That is, when the listener and the underlying socket file were created by a call to Listen or ListenUnix, then by default closing the listener will remove the socket file. but if the listener was created by a call to FileListener to use an already existing socket file, then by default closing the listener will not remove the socket file.
 func (l *UnixListener) SyscallConn() (syscall.RawConn, error)
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
