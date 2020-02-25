/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: flag
 **Element: flag.convar
 **Type: const,var
 ------------------------------------------------------------------------------------
 **Definition:
 var CommandLine = NewFlagSet(os.Args[0], ExitOnError)
 var ErrHelp = errors.New("flag: help requested")
 var Usage = func() {
     fmt.Fprintf(CommandLine.Output(), "Usage of %s:\n", os.Args[0])
     PrintDefaults()
 }
 ------------------------------------------------------------------------------------
 **Description:
 CommandLine is the default set of command-line flags, parsed from os.Args.
 The top-level functions such as BoolVar, Arg, and so on are wrappers for the
 methods of CommandLine.
 ErrHelp is the error returned if the -help or -h flag is invoked but no such
 flag is defined.
 Usage prints a usage message documenting all defined command-line flags to
 CommandLine's output, which by default is os.Stderr. It is called when an
 error occurs while parsing flags. The function is a variable that may be
 changed to point to a custom function. By default it prints a simple header
 and calls PrintDefaults; for details about the format of the output and how
 to control it, see the documentation for PrintDefaults. Custom usage
 functions may choose to exit the program; by default exiting happens anyway
 as the command line's error handling strategy is set to ExitOnError.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. CommandLine 是命令行标记的默认集合，从 os.Args 解析而来。像 BoolVar、Arg 等这样的
 顶级函数为 CommandLine 方法的包装;
 2. ErrHelp 在 -help 或 -h 标志未定义却调用了它时返回一个错误;
 3. Usage 将所有已定义命令行标记的用法信息文档打印到标准错误输出。它会在解析标记遇到错误时
 调用。该函数是个变量，因此可指向自定义的函数。它默认打印一个简单的开头并调用PrintDefaults；
 4. 关于输出格式的控制及详情，参见 PrintDefaults 的文档。
 5. 在os包中，有Args变量定义如下：
    var Args []string
    Args hold the command-line arguments, starting with the program name。
*************************************************************************************/
package main

import (
	"flag"
	"fmt"
	"os"
)

// 实际中应该用更好的变量名
var (
	h bool

	v, V bool
	t, T bool
	q    *bool

	s string
	p string
	c string
	g string
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")

	flag.BoolVar(&v, "v", false, "show version and exit")
	flag.BoolVar(&V, "V", false, "show version and configure options then exit")

	flag.BoolVar(&t, "t", false, "test configuration and exit")
	flag.BoolVar(&T, "T", false, "test configuration, dump it and exit")

	// 另一种绑定方式
	q = flag.Bool("q", false, "suppress non-error messages during configuration testing")

	// 注意 `signal`。默认是 -s string，有了 `signal` 之后，变为 -s signal
	flag.StringVar(&s, "s", "", "send `signal` to a master process: stop, quit, reopen, reload")
	flag.StringVar(&p, "p", "/usr/local/nginx/", "set `prefix` path")
	flag.StringVar(&c, "c", "conf/nginx.conf", "set configuration `file`")
	flag.StringVar(&g, "g", "conf/nginx.conf", "set global `directives` out of configuration file")

	// 改变默认的 Usage，flag包中的Usage 其实是一个函数类型。这里是覆盖默认函数实现，具体见后面Usage部分的分析
	flag.Usage = usage
}

func main() {
	flag.Parse()

	if h {
		flag.Usage()
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `nginx version: nginx/1.10.0
Usage: nginx [-hvVtTq] [-s signal] [-c filename] [-p prefix] [-g directives]

Options:
`)
	flag.PrintDefaults()
}
