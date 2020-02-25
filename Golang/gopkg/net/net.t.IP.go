/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.IP
 **Type: type
 ------------------------------------------------------------------------------------
 **Definition:
 type IP []byte
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
 ------------------------------------------------------------------------------------
 **Description:
 An IP is a single IP address, a slice of bytes. Functions in this package accept
 either 4-byte (IPv4) or 16-byte (IPv6) slices as input.
 Note that in this documentation, referring to an IP address as an IPv4 address or
 an IPv6 address is a semantic property of the address, not just the length of the
 byte slice: a 16-byte slice can still be an IPv4 address.
 DefaultMask returns the default IP mask for the IP address ip. Only IPv4 addresses
 have default masks; DefaultMask returns nil if ip is not a valid IPv4 address.
 Equal reports whether ip and x are the same IP address. An IPv4 address and that
 same address in IPv6 form are considered to be equal.
 IsGlobalUnicast reports whether ip is a global unicast address.
 The identification of global unicast addresses uses address type identification as
 defined in RFC 1122, RFC 4632 and RFC 4291 with the exception of IPv4 directed
 broadcast addresses. It returns true even if ip is in IPv4 private address space
 or local IPv6 unicast address space.
 IsInterfaceLocalMulticast reports whether ip is an interface-local multicast address.
 IsLinkLocalMulticast reports whether ip is a link-local multicast address.
 IsLinkLocalUnicast reports whether ip is a link-local unicast address.
 IsLoopback reports whether ip is a loopback address.
 IsMulticast reports whether ip is a multicast address.
 IsUnspecified reports whether ip is an unspecified address, either the IPv4 address
 "0.0.0.0" or the IPv6 address "::".
 MarshalText implements the encoding.TextMarshaler interface. The encoding is the
 same as returned by String, with one exception: When len(ip) is zero, it returns
 an empty slice.
 Mask returns the result of masking the IP address ip with mask.
 String returns the string form of the IP address ip. It returns one of 4 forms:
 - "<nil>", if ip has length 0
 - dotted decimal ("192.0.2.1"), if ip is an IPv4 or IP4-mapped IPv6 address
 - IPv6 ("2001:db8::1"), if ip is a valid IPv6 address
 - the hexadecimal form of ip, without punctuation, if no other cases apply
 To16 converts the IP address ip to a 16-byte representation. If ip is not an IP
 address (it is the wrong length), To16 returns nil.
 To4 converts the IPv4 address ip to a 4-byte representation. If ip is not an IPv4
 address, To4 returns nil.
 UnmarshalText implements the encoding.TextUnmarshaler interface. The IP address is
 expected in a form accepted by ParseIP.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. IP类型是代表单个IP地址的[]byte切片。本包的函数都可以接受4字节（IPv4）和16字节（IPv6）的切片
 作为输入。注意，IP地址是IPv4地址还是IPv6地址是语义上的属性，而不取决于切片的长度：16字节的切片
 也可以是IPv4地址;
 1. DefaultMask方法返回IP地址ip的默认子网掩码。只有IPv4有默认子网掩码；如果ip不是合法的IPv4地址，
 会返回nil；
 2. Equal方法判断ip和x是否代表同一个IP地址，代表同一地址的IPv4地址和IPv6地址也被认为是相等的；
 3. IsGlobalUnicast方法判断ip是否是全局单播地址；
 4. IsInterfaceLocalMulticast方法判断ip是否是接口本地组播地址；
 5. IsLinkLocalMulticast方法判断ip是否是链路本地组播地址；
 6. IsLinkLocalUnicast方法判断ip是否是链路本地单播地址；
 7. IsLoopback方法判断ip是否是环回地址；
 8. IsMulticast方法判断ip是否是组播地址;
 9. IsUnspecified方法判断ip是否是未指定地址；
 10. MarshalText方法实现了encoding.TextMarshaler接口，返回值和String方法一样；
 11. Mask方法认为mask为ip的子网掩码，返回ip的网络地址部分的ip,主机地址部分都置0;
 12. String方法返回IP地址ip的字符串表示。如果ip是IPv4地址，返回值的格式为点分隔的，
 如"74.125.19.99"；否则表示为IPv6格式，如"2001:4860:0:2001::68"；
 13. To16方法将一个IP地址转换为16字节表示。如果ip不是一个IP地址（长度错误），To16会返回nil；
 14. To4方法To4将一个IPv4地址转换为4字节表示。如果ip不是IPv4地址，To4会返回nil；
 15. UnmarshalText实现了encoding.TextUnmarshaler接口,IP地址字符串是ParseIP函数可以接受的格式。
*************************************************************************************/
package main

import (
	"fmt"
	"net"
)

func main() {
	ip := net.IPv4(byte(192), byte(168), byte(0), byte(1))
	fmt.Println(ip) //返回

	ipp := net.ParseIP("74.125.19.99")
	fmt.Printf("IP:%#v\n", ipp)

	ipm := ip.DefaultMask()
	fmt.Println(ipm.String())

	fmt.Println("-------IsGlobalUnicast,IsLinkLocalMulticaset,IsLinkLocalMulticaset-----------------")
	fmt.Println(ip.IsGlobalUnicast(), ipp.IsGlobalUnicast(), net.ParseIP("255.255.256.0").IsGlobalUnicast())
	fmt.Println(ip.IsLinkLocalMulticast(), ip.IsLinkLocalUnicast(), ip.IsInterfaceLocalMulticast())
	fmt.Println(ipp.IsLinkLocalMulticast(), ipp.IsLinkLocalUnicast(), ipp.IsInterfaceLocalMulticast())

	fmt.Println("-------IsLoopback,IsUnspecified,IsMulticast--------------------")
	fmt.Println(ip.IsMulticast(), ip.IsLoopback(), ip.IsUnspecified())
	fmt.Println(ipp.IsMulticast(), ipp.IsLoopback(), ipp.IsUnspecified())

	fmt.Println("-------------MarshalTex,String,Equal--------------")
	fmt.Println(ip.String(), ipp.String())
	fmt.Println(ip.Equal(ipp), ip.Equal(ip))
	mt1, _ := ip.MarshalText()
	mt2, _ := ipp.MarshalText()
	fmt.Println(string(mt1), string(mt2))
	ip.UnmarshalText([]byte("192.188.1.1"))
	fmt.Println(ip.String())

	fmt.Println("-------------MarshalTex,String,Equal--------------")
	to4 := ip.To4()
	to16 := ip.To16()
	fmt.Println(to4.String(), to16.String())
	fmt.Printf("%#v\n", to4)
	fmt.Printf("%#v\n", to16)
}
