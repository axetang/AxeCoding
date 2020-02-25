/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: net
 **Element: net.DNSError
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type DNSError struct {
     Err         string // description of the error
     Name        string // name looked for
     Server      string // server used
     IsTimeout   bool   // if true, timed out; not all timeouts set this
     IsTemporary bool   // if true, error is temporary; not all errors set this
 }
 func (e *DNSError) Error() string
 func (e *DNSError) Temporary() bool
 func (e *DNSError) Timeout() bool
 ------------------------------------------------------------------------------------
 **Description:
 DNSError represents a DNS lookup error.
 Temporary reports whether the DNS error is known to be temporary. This is not always
 known; a DNS lookup may fail due to a temporary error and return a DNSError for which
 Temporary returns false.
 Timeout reports whether the DNS lookup is known to have timed out. This is not always
 known; a DNS lookup may fail due to a timeout and return a DNSError for which Timeout
 returns false.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. DNSError代表DNS查询的错误；
 1. Error方法返回错误信息字符串；
 2. Temporary和Timeout方法判断错误性质.
*************************************************************************************/
package main

import (
	"fmt"
	"net"
)

func main() {
	dnserror := net.DNSError{
		Err:         "dns error",
		Name:        "dnserr",
		Server:      "dns search",
		IsTimeout:   true,
		IsTemporary: true,
	}
	fmt.Println(dnserror.Error())     //lookup dnserr on dns search: dns error
	fmt.Println(dnserror.Temporary()) //true
	fmt.Println(dnserror.Timeout())   //true
	fmt.Println(dnserror.Server)      //dns search
	fmt.Println(dnserror.Err)         //dns error
	fmt.Println(dnserror.Name)        //dnserr
	fmt.Println(dnserror.IsTimeout)   //true
	fmt.Println(dnserror.IsTemporary) //true
}
