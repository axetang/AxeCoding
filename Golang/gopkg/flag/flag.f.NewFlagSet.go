/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: flag
 **Element: flag.NewFlagSet
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet
 ------------------------------------------------------------------------------------
 **Description:
 NewFlagSet returns a new, empty flag set with the specified name and error
 handling property. If the name is not empty, it will be printed in the
 default usage message and in error messages.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. NewFlagSet创建一个新的、名为name，采用errorHandling为错误处理策略的FlagSet;
 2. 如果name为空，它会打印出缺省usage信息和error信息。
*************************************************************************************/
package main

import (
	"flag"
	"fmt"
	"net/url"
)

type URLValue struct {
	URL *url.URL
}

func (v URLValue) String() string {
	if v.URL != nil {
		return v.URL.String()
	}
	return ""
}

func (v URLValue) Set(s string) error {
	if u, err := url.Parse(s); err != nil {
		return err
	} else {
		*v.URL = *u
	}
	return nil
}

var u = &url.URL{}

func main() {
	fs := flag.NewFlagSet("ExampleValue", flag.ExitOnError)
	//fs := flag.NewFlagSet("", flag.ExitOnError)
	fs.Var(&URLValue{u}, "url", "URL to parse")

	fs.Parse([]string{"-url", "https://golang.org/pkg/flag/"})
	fmt.Printf(`{scheme: %q, host: %q, path: %q}`, u.Scheme, u.Host, u.Path)
	println()
}
