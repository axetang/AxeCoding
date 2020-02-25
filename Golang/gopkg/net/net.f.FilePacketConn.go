/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.FilePacketConn
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func FilePacketConn(f *os.File) (c PacketConn, err error)
 ------------------------------------------------------------------------------------
 **Description:
 FilePacketConn returns a copy of the packet network connection corresponding to the
 open file f. It is the caller's responsibility to close f when finished. Closing c
 does not affect f, and closing f does not affect c.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. FilePacketConn函数返回一个下层为文件f的数据包网络连接的拷贝。调用者有责任在结束程序前关闭f;
 2. 关闭c不会影响f，关闭f也不会影响c。本函数与各种实现了PacketConn接口的类型的File方法是对应的。
*************************************************************************************/
package main

func main() {
}