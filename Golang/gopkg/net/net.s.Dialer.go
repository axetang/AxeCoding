/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.Dialer
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Dialer struct {
     // Timeout is the maximum amount of time a dial will wait for
     // a connect to complete. If Deadline is also set, it may fail
     // earlier.
     //
     // The default is no timeout.
     //
     // When using TCP and dialing a host name with multiple IP
     // addresses, the timeout may be divided between them.
     //
     // With or without a timeout, the operating system may impose
     // its own earlier timeout. For instance, TCP timeouts are
     // often around 3 minutes.
     Timeout time.Duration
     // Deadline is the absolute point in time after which dials
     // will fail. If Timeout is set, it may fail earlier.
     // Zero means no deadline, or dependent on the operating system
     // as with the Timeout option.
     Deadline time.Time
     // LocalAddr is the local address to use when dialing an
     // address. The address must be of a compatible type for the
     // network being dialed.
     // If nil, a local address is automatically chosen.
     LocalAddr Addr
     // DualStack previously enabled RFC 6555 Fast Fallback
     // support, also known as "Happy Eyeballs", in which IPv4 is
     // tried soon if IPv6 appears to be misconfigured and
     // hanging.
     //
     // Deprecated: Fast Fallback is enabled by default. To
     // disable, set FallbackDelay to a negative value.
     DualStack bool
     // FallbackDelay specifies the length of time to wait before
     // spawning a RFC 6555 Fast Fallback connection. That is, this
     // is the amount of time to wait for IPv6 to succeed before
     // assuming that IPv6 is misconfigured and falling back to
     // IPv4.
     //
     // If zero, a default delay of 300ms is used.
     // A negative value disables Fast Fallback support.
     FallbackDelay time.Duration
     // KeepAlive specifies the keep-alive period for an active
     // network connection.
     // If zero, keep-alives are enabled if supported by the protocol
     // and operating system. Network protocols or operating systems
     // that do not support keep-alives ignore this field.
     // If negative, keep-alives are disabled.
     KeepAlive time.Duration
     // Resolver optionally specifies an alternate resolver to use.
     Resolver *Resolver
     // Cancel is an optional channel whose closure indicates that
     // the dial should be canceled. Not all types of dials support
     // cancelation.
     //
     // Deprecated: Use DialContext instead.
     Cancel <-chan struct{}
     // If Control is not nil, it is called after creating the network
     // connection but before actually dialing.
     //
     // Network and address parameters passed to Control method are not
     // necessarily the ones passed to Dial. For example, passing "tcp" to Dial
     // will cause the Control function to be called with "tcp4" or "tcp6".
     Control func(network, address string, c syscall.RawConn) error
 }
 func (d *Dialer) Dial(network, address string) (Conn, error)
 func (d *Dialer) DialContext(ctx context.Context, network, address string) (Conn, error)
 ------------------------------------------------------------------------------------
 **Description:
 A Dialer contains options for connecting to an address.
 The zero value for each field is equivalent to dialing without that option. Dialing
 with the zero value of Dialer is therefore equivalent to just calling the Dial
 function.
 Dial connects to the address on the named network.
 See func Dial for a description of the network and address parameters.
 DialContext connects to the address on the named network using the provided context.
 The provided Context must be non-nil. If the context expires before the connection is
 complete, an error is returned. Once successfully connected, any expiration of the
 context will not affect the connection.
 When using TCP, and the host in the address parameter resolves to multiple network
 addresses, any dial timeout (from d.Timeout or ctx) is spread over each consecutive
 dial, such that each is given an appropriate fraction of the time to connect. For
 example, if a host has 4 IP addresses and the timeout is 1 minute, the connect to
 each single address will be given 15 seconds to complete before trying the next one.
 See func Dial for a description of the network and address parameters.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. Dialer类型包含与某个地址建立连接时的参数, 每一个字段的零值都等价于没有该字段, 因此调用Dialer
 零值的Dial方法等价于调用Dial函数，具体定义如下：
 type Dialer struct {
    // Timeout是dial操作等待连接建立的最大时长，默认值代表没有超时。
    // 如果Deadline字段也被设置了，dial操作也可能更早失败。
    // 不管有没有设置超时，操作系统都可能强制执行它的超时设置。
    // 例如，TCP（系统）超时一般在3分钟左右。
    Timeout time.Duration
    // Deadline是一个具体的时间点期限，超过该期限后，dial操作就会失败。
    // 如果Timeout字段也被设置了，dial操作也可能更早失败。
    // 零值表示没有期限，即遵守操作系统的超时设置。
    Deadline time.Time
    // LocalAddr是dial一个地址时使用的本地地址。
    // 该地址必须是与dial的网络相容的类型。
    // 如果为nil，将会自动选择一个本地地址。
    LocalAddr Addr
    // DualStack允许单次dial操作在网络类型为"tcp"，
    // 且目的地是一个主机名的DNS记录具有多个地址时，
    // 尝试建立多个IPv4和IPv6连接，并返回第一个建立的连接。
    DualStack bool
    // KeepAlive指定一个活动的网络连接的生命周期；如果为0，会禁止keep-alive。
    // 不支持keep-alive的网络连接会忽略本字段。
    KeepAlive time.Duration
 }
 1. Dial方法在指定的网络上连接指定的地址。参见Dial函数获取网络和地址参数的描述；
 2. DialContext方法
*************************************************************************************/
package main

func main() {
}
