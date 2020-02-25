/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: url
 **Element: url.UserInfo
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 // The Userinfo type is an immutable encapsulation of username and
 // password details for a URL. An existing Userinfo value is guaranteed
 // to have a username set (potentially empty, as allowed by RFC 2396),
 // and optionally a password.
 type Userinfo struct {
     // contains filtered or unexported fields
	username    string
	password    string
	passwordSet bool
 }
 func (u *Userinfo) Password() (string, bool)
 func (u *Userinfo) String() string
 func (u *Userinfo) Username() string
 ------------------------------------------------------------------------------------
 **Description:
 The Userinfo type is an immutable encapsulation of username and password details for
 a URL. An existing Userinfo value is guaranteed to have a username set (potentially
 empty, as allowed by RFC 2396), and optionally a password.
 Password returns the password in case it is set, and whether it is set.
 String returns the encoded userinfo information in the standard form of
 "username[:password]".
 Username returns the username.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. URL结构体拥有成员User *Userinfo；
 1. Userinfo类型是一个URL的用户名+密码细节的不可改变封装体。一个已经存在的Userinfo必然拥有一个
 用户名集合，可能是空的集合），而password是可选项；
 2. Password返回用户密码；
 3. String返回编码后的userinfo信息，信息格式：“username:[:password]；
 4. Username返回username string。
*************************************************************************************/
package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	usera := url.User("axe")
	pwd, b := usera.Password()
	fmt.Println(usera.Username(), pwd, b, usera.String())

	// Parse + String preserve the original encoding.
	u, err := url.Parse("https://example.com/foo%2fbar")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.User.Username())
}
