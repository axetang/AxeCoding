/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.Flags
 **Type: type
 ------------------------------------------------------------------------------------
 **Definition:
 type Flags uint
 const (
     FlagUp           Flags = 1 << iota // interface is up
     FlagBroadcast                      // interface supports broadcast access capability
     FlagLoopback                       // interface is a loopback interface
     FlagPointToPoint                   // interface belongs to a point-to-point link
     FlagMulticast                      // interface supports multicast access capability
 )
 func (f Flags) String() string
 ------------------------------------------------------------------------------------
 **Description:
 const (
    FlagUp           Flags = 1 << iota // 接口在活动状态
    FlagBroadcast                      // 接口支持广播
    FlagLoopback                       // 接口是环回的
    FlagPointToPoint                   // 接口是点对点的
    FlagMulticast                      // 接口支持组播
 )
 ------------------------------------------------------------------------------------
 **要点总结:
 0. Flags类型记录接口状态，golang预定义了一系列常数来代表不同的状态；
 1. String方法返回Flags类型的字符串字面量表达。
*************************************************************************************/
package main

import (
	"fmt"
	"net"
)

func main() {
	f := net.FlagUp
	fmt.Println(f.String())
	f = net.FlagMulticast
	fmt.Println(f.String())
	f = net.FlagPointToPoint
	fmt.Println(f.String())
}
