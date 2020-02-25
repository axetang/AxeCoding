/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.Interface
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Interface struct {
     Index        int          // positive integer that starts at one, zero is never used
     MTU          int          // maximum transmission unit
     Name         string       // e.g., "en0", "lo0", "eth0.100"
     HardwareAddr HardwareAddr // IEEE MAC-48, EUI-48 and EUI-64 form
     Flags        Flags        // e.g., FlagUp, FlagLoopback, FlagMulticast
 }
 func (ifi *Interface) Addrs() ([]Addr, error)
 func (ifi *Interface) MulticastAddrs() ([]Addr, error)
 ------------------------------------------------------------------------------------
 **Description:
 Addrs returns a list of unicast interface addresses for a specific interface.
 MulticastAddrs returns a list of multicast, joined group addresses for a specific
 interface.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. Interface结构体类型代表一个网络接口（系统与网络的一个接点），包含接口索引到名字的映射，也包含
 接口的设备信息；
	type Interface struct {
	    Index        int          // 应该是一个正整数,如果为0的话则代表未启用
	    MTU          int          // 最大传输单位
	    Name         string       // 比如 "en0", "lo0", "eth0.100"
	    HardwareAddr HardwareAddr // IEEE MAC-48, EUI-48 and EUI-64 form
	    Flags        Flags        // 比如 FlagUp, FlagLoopback, FlagMulticast
	}
 1. Addrs方法获取该网络接口下所有地址；
 2. MulticastAddrs方法获取该网络接口下所有多播地址。
*************************************************************************************/
package main

import (
	"fmt"
	"net"
)

func main() {
	eths, err := net.Interfaces()
	if err != nil {
		fmt.Println("error")
	}

	for _, eth := range eths {
		ad, _ := eth.Addrs()
		ma, _ := eth.MulticastAddrs()
		fmt.Println("Interface:", eth.Index, eth.Name, eth.MTU, eth.Flags,
			eth.HardwareAddr)
		fmt.Println("----Addr---------------", ad)
		fmt.Println("----MulticastAddr------", ma)
	}
}
