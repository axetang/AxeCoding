/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: syslog
 **Element: syslog.New
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func New(priority Priority, tag string) (*Writer, error)
 ------------------------------------------------------------------------------------
 **Description:
 New establishes a new connection to the system log daemon. Each write to the returned 
 writer sends a log message with the given priority (a combination of the syslog 
 facility and severity) and prefix tag. If tag is empty, the os.Args[0] is used.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
