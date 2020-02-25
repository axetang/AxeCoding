/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: smtp
 **Element: smtp.Auth
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type Auth interface {
     // Start begins an authentication with a server.
     // It returns the name of the authentication protocol
     // and optionally data to include in the initial AUTH message
     // sent to the server. It can return proto == "" to indicate
     // that the authentication should be skipped.
     // If it returns a non-nil error, the SMTP client aborts
     // the authentication attempt and closes the connection.
     Start(server *ServerInfo) (proto string, toServer []byte, err error)
     // Next continues the authentication. The server has just sent
     // the fromServer data. If more is true, the server expects a
     // response, which Next should return as toServer; otherwise
     // Next should return toServer == nil.
     // If Next returns a non-nil error, the SMTP client aborts
     // the authentication attempt and closes the connection.
     Next(fromServer []byte, more bool) (toServer []byte, err error)
 }
 Auth is implemented by an SMTP authentication mechanism.
 ------------------------------------------------------------------------------------
 **Description:
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Auth接口应被每一个SMTP认证机制实现；
*************************************************************************************/
package main

func main() {
}
