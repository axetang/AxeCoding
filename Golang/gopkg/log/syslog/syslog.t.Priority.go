/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: syslog
 **Element: syslog.Priority
 **Type: type
 ------------------------------------------------------------------------------------
 **Definition:
 type Priority int
 const (
    // From /usr/include/sys/syslog.h.
    // These are the same on Linux, BSD, and OS X.
    LOG_EMERG Priority = iota
    LOG_ALERT
    LOG_CRIT
    LOG_ERR
    LOG_WARNING
    LOG_NOTICE
    LOG_INFO
    LOG_DEBUG
 )
 const (
    // From /usr/include/sys/syslog.h.
    // These are the same up to LOG_FTP on Linux, BSD, and OS X.
    LOG_KERN Priority = iota << 3
    LOG_USER
    LOG_MAIL
    LOG_DAEMON
    LOG_AUTH
    LOG_SYSLOG
    LOG_LPR
    LOG_NEWS
    LOG_UUCP
    LOG_CRON
    LOG_AUTHPRIV
    LOG_FTP
    LOG_LOCAL0
    LOG_LOCAL1
    LOG_LOCAL2
    LOG_LOCAL3
    LOG_LOCAL4
    LOG_LOCAL5
    LOG_LOCAL6
    LOG_LOCAL7
 )
 ------------------------------------------------------------------------------------
 **Description:
 The Priority is a combination of the syslog facility and severity. For example,
 LOG_ALERT | LOG_FTP sends an alert severity message from the FTP facility. The
 default severity is LOG_EMERG; the default facility is LOG_KERN.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
