/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: url
 **Element: url.URL
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type URL struct {
     Scheme     stringw
     Opaque     string    // encoded opaque data
     User       *Userinfo // username and password information
     Host       string    // host or host:port
     Path       string    // path (relative paths may omit leading slash)
     RawPath    string    // encoded path hint (see EscapedPath method)
     ForceQuery bool      // append a query ('?') even if RawQuery is empty
     RawQuery   string    // encoded query values, without '?'
     Fragment   string    // fragment for references, without '#'
 }
 func (u *URL) EscapedPath() string
 func (u *URL) Hostname() string
 func (u *URL) IsAbs() bool
 func (u *URL) MarshalBinary() (text []byte, err error)
 func (u *URL) Parse(ref string) (*URL, error)
 func (u *URL) Port() string
 func (u *URL) Query() Values
 func (u *URL) RequestURI() string
 func (u *URL) ResolveReference(ref *URL) *URL
 func (u *URL) String() string
 func (u *URL) UnmarshalBinary(text []byte) error
 ------------------------------------------------------------------------------------
 **Description:
 A URL represents a parsed URL (technically, a URI reference).
 The general form represented is:
 [scheme:][//[userinfo@]host][/]path[?query][#fragment]
 URLs that do not start with a slash after the scheme are interpreted as:
 scheme:opaque[?query][#fragment]
 Note that the Path field is stored in decoded form: /%47%6f%2f becomes /Go/. A
 consequence is that it is impossible to tell which slashes in the Path were slashes
 in the raw URL and which were %2f. This distinction is rarely important, but when it
 is, code must not use Path directly. The Parse function sets both Path and RawPath in
 the URL it returns, and URL's String method uses RawPath if it is a valid encoding of
 Path, by calling the EscapedPath method.
 EscapedPath returns the escaped form of u.Path. In general there are multiple
 possible escaped forms of any path. EscapedPath returns u.RawPath when it is a
 valid escaping of u.Path. Otherwise EscapedPath ignores u.RawPath and computes an
 escaped form on its own. The String and RequestURI methods use EscapedPath to
 construct their results. In general, code should call EscapedPath instead of
 reading u.RawPath directly.
 Hostname returns u.Host, without any port number.
 If Host is an IPv6 literal with a port number, Hostname returns the IPv6 literal
 without the square brackets. IPv6 literals may include a zone identifier.
 IsAbs reports whether the URL is absolute. Absolute means that it has a non-empty
 scheme.
 Parse parses a URL in the context of the receiver. The provided URL may be relative
 or absolute. Parse returns nil, err on parse failure, otherwise its return value is
 the same as ResolveReference.
 Port returns the port part of u.Host, without the leading colon. If u.Host doesn't
 contain a port, Port returns an empty string.
 Query parses RawQuery and returns the corresponding values. It silently discards
 malformed value pairs. To check errors use ParseQuery.
 RequestURI returns the encoded path?query or opaque?query string that would be used
 in an HTTP request for u.
 ResolveReference resolves a URI reference to an absolute URI from an absolute base
 URI u, per RFC 3986 Section 5.2. The URI reference may be relative or absolute.
 ResolveReference always returns a new URL instance, even if the returned URL is
 identical to either the base or reference. If ref is an absolute URL, then
 ResolveReference ignores base and returns a copy of ref.
 String reassembles the URL into a valid URL string. The general form of the result
 is one of:
 scheme:opaque?query#fragment
 scheme://userinfo@host/path?query#fragment
 If u.Opaque is non-empty, String uses the first form; otherwise it uses the second
 form. Any non-ASCII characters in host are escaped. To obtain the path, String uses
 u.EscapedPath().
 In the second form, the following rules apply:
 - if u.Scheme is empty, scheme: is omitted.
 - if u.User is nil, userinfo@ is omitted.
 - if u.Host is empty, host/ is omitted.
 - if u.Scheme and u.Host are empty and u.User is nil,
    the entire scheme://userinfo@host/ is omitted.
 - if u.Host is non-empty and u.Path begins with a /,
    the form host/path does not add its own /.
 - if u.RawQuery is empty, ?query is omitted.
 - if u.Fragment is empty, #fragment is omitted.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. URL类型代表一个解析后的URL（或者说，一个URL参照）。URL基本格式如下：
      scheme://[userinfo@]host/path[?query][#fragment]
    scheme后不是冒号加双斜线的URL被解释为如下格式：
      scheme:opaque[?query][#fragment]
    注意路径字段是以解码后的格式保存的，如/%47%6f%2f会变成/Go/。这导致我们无法确定Path字段中的斜线
    是来自原始URL还是解码前的%2f。除非一个客户端必须使用其他程序/函数来解析原始URL或者重构原始URL，
    这个区别并不重要。此时，HTTP服务端可以查询req.RequestURI，而HTTP客户端可以使用
    URL{Host: "example.com", Opaque: "//example.com/Go%2f"}代替
    {Host: "example.com", Path: "/Go/"}
 1. EscapedPath方法返回URL.Path的escaped form；
 2. HostName方法返回Hostname；
 3. IsAbs方法判断u *URL是否是绝对路径地址；
 4. MarshalBinary方法
 5. Parse方法解析字符串参数ref string为一个合法的URL,ref可以是相对的，也可以是绝对的；
 6. Port方法返回Hostname的端口号；
 7. Query方法解析URL的RawQuery字符串并返回一个Value类型，Value类型为一个字符串-字符串map类型，
    Value定义如下：
      type Values map[string][]string
 8. RequestURI方法返回编码好的path?query或opaque?query字符串，用在HTTP请求里;
 9. ResolveReference方法
 10. String方法组合并返回合法的URL字符串；
 11. UnmarshalBinary方法
*************************************************************************************/
package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	//a example for URL
	u, err := url.Parse("http://bing.com/search?q=dotnet")
	if err != nil {
		log.Fatal(err)
	}
	u.Scheme = "https"
	u.Host = "google.com"
	q := u.Query()
	q.Set("q", "golang")
	u.RawQuery = q.Encode()
	fmt.Println(u)

	fmt.Println("----------IsAbs，Parse，String，Hostname,Port-----------------")
	fmt.Println("http://bing.com/search?q=dotnet IsAbs?", u.IsAbs())
	u1, err := url.Parse("//bing.com/search?q=dotnet")
	fmt.Println("//bing.com/search?q=dotnet IsAbs?", u1.IsAbs())
	u2 := url.URL{Host: "example.com:5000", Path: "foo"}
	fmt.Println(u2.IsAbs())
	u2.Scheme = "http"
	fmt.Println(u2.IsAbs())
	fmt.Println(u2.String())
	fmt.Println(u2.Hostname())
	fmt.Println(u2.Port())

	fmt.Println("----------Query,-----------")
	u, err = url.Parse("https://example.org/?a=1&a=2&b=&=3&&c=4&=5&")
	fmt.Println("u.RawQuery=", u.RawQuery)
	if err != nil {
		log.Fatal(err)
	}
	q = u.Query()
	fmt.Println(q["a"])
	fmt.Println(q["b"])
	fmt.Println(q["c"])
	fmt.Println(q.Get("b"))
	fmt.Println(q.Get(""))

	fmt.Println("----------------RequestURI---------------")
	u, err = url.Parse("https://example.org/path?foo=bar")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.RequestURI())

	fmt.Println("-----------------Parse--------------")
	u, err = url.Parse("https://example.org")
	if err != nil {
		log.Fatal(err)
	}
	rel, err := u.Parse("/foo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rel)
	_, err = u.Parse(":foo")
	if _, ok := err.(*url.Error); !ok {
		fmt.Println("error")
		log.Fatal(err)
	}

	fmt.Println("----------------Hostname&Port-------------")
	u, err = url.Parse("https://example.org:8000/path")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Hostname())
	u, err = url.Parse("https://[2001:0db8:85a3:0000:0000:8a2e:0370:7334]:17000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Hostname(), u.Port())

	fmt.Println("\n---------EscapedPath--------------")
	u, err = url.Parse("http://example.com/path with spaces")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.EscapedPath())

	fmt.Println("\n---------MarshalBinary--------------")
	u, _ = url.Parse("https://example.org")
	b, err := u.MarshalBinary()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%b\n", b)
	fmt.Printf("%s\n", b)

	fmt.Println("\n---------UnmarshalBinary--------------")
	u = &url.URL{}
	err = u.UnmarshalBinary([]byte("https://example.org/foo"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", u.String())

}
