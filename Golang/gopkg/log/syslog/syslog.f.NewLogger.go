/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: syslog
 **Element: syslog.NewLogger
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func NewLogger(p Priority, logFlag int) (*log.Logger, error)
 ------------------------------------------------------------------------------------
 **Description:
 NewLogger creates a log.Logger whose output is written to the system log service
 with the specified priority, a combination of the syslog facility and severity.
 The logFlag argument is the flag set passed through to log.New to create the Logger.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

import (
	"log/syslog"
)

func main() {
	nl, _ := syslog.NewLogger(syslog.LOG_NOTICE, 19)
	nl.Print("this is a syslog created by axe!")

}
