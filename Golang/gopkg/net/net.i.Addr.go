/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.Addr
 **Type: interfface
 ------------------------------------------------------------------------------------
 **Definition:
 type Addr interface {
     Network() string // name of the network (for example, "tcp", "udp")
     String() string  // string form of address (for example, "192.0.2.1:25", "[2001:db8::1]:80")
 }
 ------------------------------------------------------------------------------------
 **Description:
 Addr represents a network end point address.
 The two methods Network and String conventionally return strings that can be passed
 as the arguments to Dial, but the exact form and meaning of the strings is up to the
 implementation.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. Addr是一个接口,代表一个网络端的地址
 结构:
	type Addr interface {
   	 	Network() string // 网络名字符串
   	 	String() string  // 地址字符串
    }
*************************************************************************************/
package main

func main() {
}
