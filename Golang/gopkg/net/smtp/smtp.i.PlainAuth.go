/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: smtp
 **Element: smtp.PlainAuth
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func PlainAuth(identity, username, password, host string) Auth
 ------------------------------------------------------------------------------------
 **Description:
 PlainAuth returns an Auth that implements the PLAIN authentication mechanism as
 defined in RFC 4616. The returned Auth uses the given username and password to
 authenticate to host and act as identity. Usually identity should be the empty
 string, to act as username.
 PlainAuth will only send the credentials if the connection is using TLS or is
 connected to localhost. Otherwise authentication will fail with an error, without
 sending the credentials.
  ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
