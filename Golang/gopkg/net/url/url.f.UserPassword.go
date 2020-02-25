/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: url
 **Element: url.UserPassword
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func UserPassword(username, password string) *Userinfo
 ------------------------------------------------------------------------------------
 **Description:
 UserPassword returns a Userinfo containing the provided username and password.
 This functionality should only be used with legacy web sites. RFC 2396 warns that
 interpreting Userinfo this way “is NOT RECOMMENDED, because the passing of
 authentication information in clear text (such as URI) has proven to be a security
 risk in almost every case where it has been used.”
 ------------------------------------------------------------------------------------
 **要点总结:
 1. UserPassword函数返回一个用户名设置为username、密码设置为password的*Userinfo;
 2. 这个函数应该只用于老式的站点，因为风险很大，不建议使用，参见RFC2396。
*************************************************************************************/
package main

import (
	"fmt"
	"net/url"
)

func main() {
	ui := url.UserPassword("axe", "pwtang")
	fmt.Println(ui.String())
}
