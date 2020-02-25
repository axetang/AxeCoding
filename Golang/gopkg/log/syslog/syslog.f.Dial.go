/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: syslog
 **Element: syslog.Dial
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Dial(network, raddr string, priority Priority, tag string) (*Writer, error)
 ------------------------------------------------------------------------------------
 **Description:
 Dial establishes a connection to a log daemon by connecting to address raddr on the
 specified network. Each write to the returned writer sends a log message with the
 facility and severity (from priority) and tag. If tag is empty, the os.Args[0] is
 used. If network is empty, Dial will connect to the local syslog server. Otherwise,
 see the documentation for net.Dial for valid values of network and raddr.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
