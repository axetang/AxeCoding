# Golang package源码
这个目录下所有文件和子目录是Golang的标准库的相应标准包的源代码，供学习过程中查阅和参考q 
**标准库在线文档链接如下：**  
- [English Version](https://godoc.org/)
- [中文链接](http://docscn.studygolang.com/pkg/)
  
# package package net
## import "net"

Package net provides a portable interface for network I/O, including TCP/IP, UDP, domain name resolution, and Unix domain sockets.

Although the package provides access to low-level networking primitives, most clients will need only the basic interface provided by the Dial, Listen, and Accept functions and the associated Conn and Listener interfaces. The crypto/tls package uses the same interfaces and similar Dial and Listen functions.

The Dial function connects to a server:
```
conn, err := net.Dial("tcp", "golang.org:80")
if err != nil {
    // handle error
}
fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
status, err := bufio.NewReader(conn).ReadString('\n')
// ...
```
The Listen function creates servers:
```
ln, err := net.Listen("tcp", ":8080")
if err != nil {
    // handle error
}
for {
    conn, err := ln.Accept()
    if err != nil {
        // handle error
    }
    go handleConnection(conn)
}
```
## Name Resolution

The method for resolving domain names, whether indirectly with functions like Dial or directly with functions like LookupHost and LookupAddr, varies by operating system.

On Unix systems, the resolver has two options for resolving names. It can use a pure Go resolver that sends DNS requests directly to the servers listed in /etc/resolv.conf, or it can use a cgo-based resolver that calls C library routines such as getaddrinfo and getnameinfo.

By default the pure Go resolver is used, because a blocked DNS request consumes only a goroutine, while a blocked C call consumes an operating system thread. When cgo is available, the cgo-based resolver is used instead under a variety of conditions: on systems that do not let programs make direct DNS requests (OS X), when the LOCALDOMAIN environment variable is present (even if empty), when the RES_OPTIONS or HOSTALIASES environment variable is non-empty, when the ASR_CONFIG environment variable is non-empty (OpenBSD only), when /etc/resolv.conf or /etc/nsswitch.conf specify the use of features that the Go resolver does not implement, and when the name being looked up ends in .local or is an mDNS name.

The resolver decision can be overridden by setting the netdns value of the GODEBUG environment variable (see package runtime) to go or cgo, as in:

export GODEBUG=netdns=go    # force pure Go resolver
export GODEBUG=netdns=cgo   # force cgo resolver
The decision can also be forced while building the Go source tree by setting the netgo or netcgo build tag.

A numeric netdns setting, as in GODEBUG=netdns=1, causes the resolver to print debugging information about its decisions. To force a particular resolver while also printing debugging information, join the two settings by a plus sign, as in GODEBUG=netdns=go+1.

On Plan 9, the resolver always accesses /net/cs and /net/dns.

On Windows, the resolver always uses C library functions, such as GetAddrInfo and DnsQuery.

## Index
```
Constants
Variables
func JoinHostPort(host, port string) string
func LookupAddr(addr string) (names []string, err error)
func LookupCNAME(host string) (cname string, err error)
func LookupHost(host string) (addrs []string, err error)
func LookupPort(network, service string) (port int, err error)
func LookupTXT(name string) ([]string, error)
func ParseCIDR(s string) (IP, *IPNet, error)
func Pipe() (Conn, Conn)
func SplitHostPort(hostport string) (host, port string, err error)
type Addr
func InterfaceAddrs() ([]Addr, error)
type AddrError
func (e *AddrError) Error() string
func (e *AddrError) Temporary() bool
func (e *AddrError) Timeout() bool
type Buffers
func (v *Buffers) Read(p []byte) (n int, err error)
func (v *Buffers) WriteTo(w io.Writer) (n int64, err error)
type Conn
func Dial(network, address string) (Conn, error)
func DialTimeout(network, address string, timeout time.Duration) (Conn, error)
func FileConn(f *os.File) (c Conn, err error)
type DNSConfigError
func (e *DNSConfigError) Error() string
func (e *DNSConfigError) Temporary() bool
func (e *DNSConfigError) Timeout() bool
type DNSError
func (e *DNSError) Error() string
func (e *DNSError) Temporary() bool
func (e *DNSError) Timeout() bool
type Dialer
func (d *Dialer) Dial(network, address string) (Conn, error)
func (d *Dialer) DialContext(ctx context.Context, network, address string) (Conn, error)
type Error
type Flags
func (f Flags) String() string
type HardwareAddr
func ParseMAC(s string) (hw HardwareAddr, err error)
func (a HardwareAddr) String() string
type IP
func IPv4(a, b, c, d byte) IP
func LookupIP(host string) ([]IP, error)
func ParseIP(s string) IP
func (ip IP) DefaultMask() IPMask
func (ip IP) Equal(x IP) bool
func (ip IP) IsGlobalUnicast() bool
func (ip IP) IsInterfaceLocalMulticast() bool
func (ip IP) IsLinkLocalMulticast() bool
func (ip IP) IsLinkLocalUnicast() bool
func (ip IP) IsLoopback() bool
func (ip IP) IsMulticast() bool
func (ip IP) IsUnspecified() bool
func (ip IP) MarshalText() ([]byte, error)
func (ip IP) Mask(mask IPMask) IP
func (ip IP) String() string
func (ip IP) To16() IP
func (ip IP) To4() IP
func (ip *IP) UnmarshalText(text []byte) error
type IPAddr
func ResolveIPAddr(network, address string) (*IPAddr, error)
func (a *IPAddr) Network() string
func (a *IPAddr) String() string
type IPConn
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
type IPMask
func CIDRMask(ones, bits int) IPMask
func IPv4Mask(a, b, c, d byte) IPMask
func (m IPMask) Size() (ones, bits int)
func (m IPMask) String() string
type IPNet
func (n *IPNet) Contains(ip IP) bool
func (n *IPNet) Network() string
func (n *IPNet) String() string
type Interface
func InterfaceByIndex(index int) (*Interface, error)
func InterfaceByName(name string) (*Interface, error)
func Interfaces() ([]Interface, error)
func (ifi *Interface) Addrs() ([]Addr, error)
func (ifi *Interface) MulticastAddrs() ([]Addr, error)
type InvalidAddrError
func (e InvalidAddrError) Error() string
func (e InvalidAddrError) Temporary() bool
func (e InvalidAddrError) Timeout() bool
type ListenConfig
func (lc *ListenConfig) Listen(ctx context.Context, network, address string) (Listener, error)
func (lc *ListenConfig) ListenPacket(ctx context.Context, network, address string) (PacketConn, error)
type Listener
func FileListener(f *os.File) (ln Listener, err error)
func Listen(network, address string) (Listener, error)
type MX
func LookupMX(name string) ([]*MX, error)
type NS
func LookupNS(name string) ([]*NS, error)
type OpError
func (e *OpError) Error() string
func (e *OpError) Temporary() bool
func (e *OpError) Timeout() bool
type PacketConn
func FilePacketConn(f *os.File) (c PacketConn, err error)
func ListenPacket(network, address string) (PacketConn, error)
type ParseError
func (e *ParseError) Error() string
type Resolver
func (r *Resolver) LookupAddr(ctx context.Context, addr string) (names []string, err error)
func (r *Resolver) LookupCNAME(ctx context.Context, host string) (cname string, err error)
func (r *Resolver) LookupHost(ctx context.Context, host string) (addrs []string, err error)
func (r *Resolver) LookupIPAddr(ctx context.Context, host string) ([]IPAddr, error)
func (r *Resolver) LookupMX(ctx context.Context, name string) ([]*MX, error)
func (r *Resolver) LookupNS(ctx context.Context, name string) ([]*NS, error)
func (r *Resolver) LookupPort(ctx context.Context, network, service string) (port int, err error)
func (r *Resolver) LookupSRV(ctx context.Context, service, proto, name string) (cname string, addrs []*SRV, err error)
func (r *Resolver) LookupTXT(ctx context.Context, name string) ([]string, error)
type SRV
func LookupSRV(service, proto, name string) (cname string, addrs []*SRV, err error)
type TCPAddr
func ResolveTCPAddr(network, address string) (*TCPAddr, error)
func (a *TCPAddr) Network() string
func (a *TCPAddr) String() string
type TCPConn
func DialTCP(network string, laddr, raddr *TCPAddr) (*TCPConn, error)
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
type TCPListener
func ListenTCP(network string, laddr *TCPAddr) (*TCPListener, error)
func (l *TCPListener) Accept() (Conn, error)
func (l *TCPListener) AcceptTCP() (*TCPConn, error)
func (l *TCPListener) Addr() Addr
func (l *TCPListener) Close() error
func (l *TCPListener) File() (f *os.File, err error)
func (l *TCPListener) SetDeadline(t time.Time) error
func (l *TCPListener) SyscallConn() (syscall.RawConn, error)
type UDPAddr
func ResolveUDPAddr(network, address string) (*UDPAddr, error)
func (a *UDPAddr) Network() string
func (a *UDPAddr) String() string
type UDPConn
func DialUDP(network string, laddr, raddr *UDPAddr) (*UDPConn, error)
func ListenMulticastUDP(network string, ifi *Interface, gaddr *UDPAddr) (*UDPConn, error)
func ListenUDP(network string, laddr *UDPAddr) (*UDPConn, error)
func (c *UDPConn) Close() error
func (c *UDPConn) File() (f *os.File, err error)
func (c *UDPConn) LocalAddr() Addr
func (c *UDPConn) Read(b []byte) (int, error)
func (c *UDPConn) ReadFrom(b []byte) (int, Addr, error)
func (c *UDPConn) ReadFromUDP(b []byte) (int, *UDPAddr, error)
func (c *UDPConn) ReadMsgUDP(b, oob []byte) (n, oobn, flags int, addr *UDPAddr, err error)
func (c *UDPConn) RemoteAddr() Addr
func (c *UDPConn) SetDeadline(t time.Time) error
func (c *UDPConn) SetReadBuffer(bytes int) error
func (c *UDPConn) SetReadDeadline(t time.Time) error
func (c *UDPConn) SetWriteBuffer(bytes int) error
func (c *UDPConn) SetWriteDeadline(t time.Time) error
func (c *UDPConn) SyscallConn() (syscall.RawConn, error)
func (c *UDPConn) Write(b []byte) (int, error)
func (c *UDPConn) WriteMsgUDP(b, oob []byte, addr *UDPAddr) (n, oobn int, err error)
func (c *UDPConn) WriteTo(b []byte, addr Addr) (int, error)
func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error)
type UnixAddr
func ResolveUnixAddr(network, address string) (*UnixAddr, error)
func (a *UnixAddr) Network() string
func (a *UnixAddr) String() string
type UnixConn
func DialUnix(network string, laddr, raddr *UnixAddr) (*UnixConn, error)
func ListenUnixgram(network string, laddr *UnixAddr) (*UnixConn, error)
func (c *UnixConn) Close() error
func (c *UnixConn) CloseRead() error
func (c *UnixConn) CloseWrite() error
func (c *UnixConn) File() (f *os.File, err error)
func (c *UnixConn) LocalAddr() Addr
func (c *UnixConn) Read(b []byte) (int, error)
func (c *UnixConn) ReadFrom(b []byte) (int, Addr, error)
func (c *UnixConn) ReadFromUnix(b []byte) (int, *UnixAddr, error)
func (c *UnixConn) ReadMsgUnix(b, oob []byte) (n, oobn, flags int, addr *UnixAddr, err error)
func (c *UnixConn) RemoteAddr() Addr
func (c *UnixConn) SetDeadline(t time.Time) error
func (c *UnixConn) SetReadBuffer(bytes int) error
func (c *UnixConn) SetReadDeadline(t time.Time) error
func (c *UnixConn) SetWriteBuffer(bytes int) error
func (c *UnixConn) SetWriteDeadline(t time.Time) error
func (c *UnixConn) SyscallConn() (syscall.RawConn, error)
func (c *UnixConn) Write(b []byte) (int, error)
func (c *UnixConn) WriteMsgUnix(b, oob []byte, addr *UnixAddr) (n, oobn int, err error)
func (c *UnixConn) WriteTo(b []byte, addr Addr) (int, error)
func (c *UnixConn) WriteToUnix(b []byte, addr *UnixAddr) (int, error)
type UnixListener
func ListenUnix(network string, laddr *UnixAddr) (*UnixListener, error)
func (l *UnixListener) Accept() (Conn, error)
func (l *UnixListener) AcceptUnix() (*UnixConn, error)
func (l *UnixListener) Addr() Addr
func (l *UnixListener) Close() error
func (l *UnixListener) File() (f *os.File, err error)
func (l *UnixListener) SetDeadline(t time.Time) error
func (l *UnixListener) SetUnlinkOnClose(unlink bool)
func (l *UnixListener) SyscallConn() (syscall.RawConn, error)
type UnknownNetworkError
func (e UnknownNetworkError) Error() string
func (e UnknownNetworkError) Temporary() bool
func (e UnknownNetworkError) Timeout() bool
Bugs
```
## Package Files

addrselect.go cgo_linux.go cgo_resnew.go cgo_socknew.go cgo_unix.go conf.go dial.go dnsclient.go dnsclient_unix.go dnsconfig_unix.go error_posix.go error_unix.go fd_unix.go file.go file_unix.go hook.go hook_unix.go hosts.go interface.go interface_aix.go interface_linux.go ip.go iprawsock.go iprawsock_posix.go ipsock.go ipsock_posix.go lookup.go lookup_unix.go mac.go net.go nss.go parse.go pipe.go port.go port_unix.go rawconn.go sendfile_linux.go sock_cloexec.go sock_linux.go sock_posix.go sockaddr_posix.go sockopt_aix.go sockopt_linux.go sockopt_posix.go sockoptip_linux.go sockoptip_posix.go splice_linux.go tcpsock.go tcpsock_posix.go tcpsockopt_posix.go tcpsockopt_unix.go udpsock.go udpsock_posix.go unixsock.go unixsock_posix.go writev_unix.go