/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: url
 **Element: url.ParseRequestURI
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func ParseRequestURI(rawurl string) (*URL, error)
 ------------------------------------------------------------------------------------
 **Description:
 ParseRequestURI parses rawurl into a URL structure. It assumes that rawurl was
 received in an HTTP request, so the rawurl is interpreted only as an absolute URI
 or an absolute path. The string rawurl is assumed not to have a #fragment suffix.
 (Web browsers strip #fragment before sending the URL to a web server.)
 ------------------------------------------------------------------------------------
 **要点总结:
 1. ParseRequestURI对参数字符串rawurl进行解析，并生成一个URL结构体返回；
 2. 它假定rawurl是从一个HTTP request中获得，因此rawurl只会被翻译成一个绝对路径的URI或者一个
 绝对路径；
 3. rawurl字符串假定是没有#fragment后缀段的，因为web浏览器会在吧URL发给webserver前
 把#fragment抽离。
*************************************************************************************/
package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u, err := url.ParseRequestURI("https://example.org/#path?foo=bar")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)
	fmt.Println(u.Scheme, "\n", u.Opaque, "\n", u.User, "\n", u.Host, "\n", u.Path, "\n", u.RawPath, "\n",
		u.ForceQuery, "\n", u.RawQuery, "\n", u.Fragment)
	/*
		Scheme     stringw
		Opaque     string    // encoded opaque data
		User       *Userinfo // username and password information
		Host       string    // host or host:port
		Path       string    // path (relative paths may omit leading slash)
		RawPath    string    // encoded path hint (see EscapedPath method)
		ForceQuery bool      // append a query ('?') even if RawQuery is empty
		RawQuery   string    // encoded query values, without '?'
		Fragment   string    // fragment for references, without '#'

	*/
}
