/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: syslog
 **Element: syslog.Writer
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Writer struct {
    // contains filtered or unexported fields
 }
 func Dial(network, raddr string, priority Priority, tag string) (*Writer, error)
 func New(priority Priority, tag string) (*Writer, error)
 func (w *Writer) Alert(m string) error
 func (w *Writer) Close() error
 func (w *Writer) Crit(m string) error
 func (w *Writer) Debug(m string) error
 func (w *Writer) Emerg(m string) error
 func (w *Writer) Err(m string) error
 func (w *Writer) Info(m string) error
 func (w *Writer) Notice(m string) error
 func (w *Writer) Warning(m string) error
 func (w *Writer) Write(b []byte) (int, error)
 ------------------------------------------------------------------------------------
 **Description:
 A Writer is a connection to a syslog server.
 Dial establishes a connection to a log daemon by connecting to address raddr on the
 specified network. Each write to the returned writer sends a log message with the
 facility and severity (from priority) and tag. If tag is empty, the os.Args[0] is
 used. If network is empty, Dial will connect to the local syslog server. Otherwise,
 see the documentation for net.Dial for valid values of network and raddr.
 New establishes a new connection to the system log daemon. Each write to the
 returned writer sends a log message with the given priority (a combination of the
 syslog facility and severity) and prefix tag. If tag is empty, the os.Args[0] is used.
 Alert logs a message with severity LOG_ALERT, ignoring the severity passed to New.
 Close closes a connection to the syslog daemon.
 Crit logs a message with severity LOG_CRIT, ignoring the severity passed to New.
 Debug logs a message with severity LOG_DEBUG, ignoring the severity passed to New.
 Emerg logs a message with severity LOG_EMERG, ignoring the severity passed to New.
 Err logs a message with severity LOG_ERR, ignoring the severity passed to New.
 Info logs a message with severity LOG_INFO, ignoring the severity passed to New.
 Notice logs a message with severity LOG_NOTICE, ignoring the severity passed to New.
 Warning logs a message with severity LOG_WARNING, ignoring the severity passed to New.
 Write sends a log message to the syslog daemon.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
