/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: flag
 **Element: flag.FlagSet
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type FlagSet struct {
	// Usage is the function called when an error occurs while parsing flags.
	// The field is a function (not a method) that may be changed to point to
	// a custom error handler. What happens after Usage is called depends
	// on the ErrorHandling setting; for the command line, this defaults
	// to ExitOnError, which exits the program after calling Usage.
	Usage func()

    // contains filtered or unexported fields
	name          string
	parsed        bool
	actual        map[string]*Flag
	formal        map[string]*Flag
	args          []string // arguments after flags
	errorHandling ErrorHandling
	output        io.Writer // nil means stderr; use out() accessor
 }
 //func NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet
 func (f *FlagSet) Arg(i int) string
 func (f *FlagSet) Args() []string
 func (f *FlagSet) Bool(name string, value bool, usage string) *bool
 func (f *FlagSet) BoolVar(p *bool, name string, value bool, usage string)
 func (f *FlagSet) Duration(name string, value time.Duration, usage string) *time.Duration
 func (f *FlagSet) DurationVar(p *time.Duration, name string, value time.Duration, usage string)
 func (f *FlagSet) ErrorHandling() ErrorHandling
 func (f *FlagSet) Float64(name string, value float64, usage string) *float64
 func (f *FlagSet) Float64Var(p *float64, name string, value float64, usage string)
 func (f *FlagSet) Init(name string, errorHandling ErrorHandling)
 func (f *FlagSet) Int(name string, value int, usage string) *int
 func (f *FlagSet) Int64(name string, value int64, usage string) *int64
 func (f *FlagSet) Int64Var(p *int64, name string, value int64, usage string)
 func (f *FlagSet) IntVar(p *int, name string, value int, usage string)
 func (f *FlagSet) Lookup(name string) *Flag
 func (f *FlagSet) NArg() int
 func (f *FlagSet) NFlag() int
 func (f *FlagSet) Name() string
 func (f *FlagSet) Output() io.Writer
 func (f *FlagSet) Parse(arguments []string) error
 func (f *FlagSet) Parsed() bool
 func (f *FlagSet) PrintDefaults()
 func (f *FlagSet) Set(name, value string) error
 func (f *FlagSet) SetOutput(output io.Writer)
 func (f *FlagSet) String(name string, value string, usage string) *string
 func (f *FlagSet) StringVar(p *string, name string, value string, usage string)
 func (f *FlagSet) Uint(name string, value uint, usage string) *uint
 func (f *FlagSet) Uint64(name string, value uint64, usage string) *uint64
 func (f *FlagSet) Uint64Var(p *uint64, name string, value uint64, usage string)
 func (f *FlagSet) UintVar(p *uint, name string, value uint, usage string)
 func (f *FlagSet) Var(value Value, name string, usage string)
 func (f *FlagSet) Visit(fn func(*Flag))
 func (f *FlagSet) VisitAll(fn func(*Flag))
 ------------------------------------------------------------------------------------
 **Description:
 A FlagSet represents a set of defined flags. The zero value of a FlagSet has no name
 and has ContinueOnError error handling.
 Arg returns the i'th argument. Arg(0) is the first remaining argument after flags
 have been processed. Arg returns an empty string if the requested element does not
 exist.
 Args returns the non-flag arguments.
 Bool defines a bool flag with specified name, default value, and usage string. The
 return value is the address of a bool variable that stores the value of the flag.
 BoolVar defines a bool flag with specified name, default value, and usage string.
 The argument p points to a bool variable in which to store the value of the flag.
 Duration defines a time.Duration flag with specified name, default value, and usage
 string. The return value is the address of a time.Duration variable that stores the
 value of the flag. The flag accepts a value acceptable to time.ParseDuration.
 DurationVar defines a time.Duration flag with specified name, default value, and
 usage string. The argument p points to a time.Duration variable in which to store
 the value of the flag. The flag accepts a value acceptable to time.ParseDuration.
 ErrorHandling returns the error handling behavior of the flag set.
 Float64 defines a float64 flag with specified name, default value, and usage string.
 The return value is the address of a float64 variable that stores the value of the
 flag.
 Float64Var defines a float64 flag with specified name, default value, and usage
 string. The argument p points to a float64 variable in which to store the value of
 the flag.
 Init sets the name and error handling property for a flag set. By default, the zero
 FlagSet uses an empty name and the ContinueOnError error handling policy.
 Int defines an int flag with specified name, default value, and usage string. The
 return value is the address of an int variable that stores the value of the flag.
 Int64 defines an int64 flag with specified name, default value, and usage string.
 The return value is the address of an int64 variable that stores the value of the
 flag.
 Int64Var defines an int64 flag with specified name, default value, and usage string.
 The argument p points to an int64 variable in which to store the value of the flag.
 IntVar defines an int flag with specified name, default value, and usage string.
 The argument p points to an int variable in which to store the value of the flag.
 Lookup returns the Flag structure of the named flag, returning nil if none exists.
 NArg is the number of arguments remaining after flags have been processed.
 NFlag returns the number of flags that have been set.
 Name returns the name of the flag set.
 Output returns the destination for usage and error messages. os.Stderr is returned
 if output was not set or was set to nil.
 Parse parses flag definitions from the argument list, which should not include the
 command name. Must be called after all flags in the FlagSet are defined and before
 flags are accessed by the program. The return value will be ErrHelp if -help or -h
 were set but not defined.
 Parsed reports whether f.Parse has been called.
 PrintDefaults prints, to standard error unless configured otherwise, the default
 values of all defined command-line flags in the set. See the documentation for the
 global function PrintDefaults for more information.
 Set sets the value of the named flag.
 SetOutput sets the destination for usage and error messages. If output is nil,
 os.Stderr is used.
 String defines a string flag with specified name, default value, and usage string.
 The return value is the address of a string variable that stores the value of the
 flag.
 StringVar defines a string flag with specified name, default value, and usage
 string. The argument p points to a string variable in which to store the value of
 the flag.
 Uint defines a uint flag with specified name, default value, and usage string. The
 return value is the address of a uint variable that stores the value of the flag.
 Uint64 defines a uint64 flag with specified name, default value, and usage string.
 The return value is the address of a uint64 variable that stores the value of the
 flag.
 Uint64Var defines a uint64 flag with specified name, default value, and usage string.
 The argument p points to a uint64 variable in which to store the value of the flag.
 UintVar defines a uint flag with specified name, default value, and usage string.
 The argument p points to a uint variable in which to store the value of the flag.
 Var defines a flag with the specified name and usage string. The type and value of
 the flag are represented by the first argument, of type Value, which typically holds
 a user-defined implementation of Value. For instance, the caller could create a flag
 that turns a comma-separated string into a slice of strings by giving the slice the
 methods of Value; in particular, Set would decompose the comma-separated string into
 the slice.
 Visit visits the flags in lexicographical order, calling fn for each. It visits only
 those flags that have been set.
 VisitAll visits the flags in lexicographical order, calling fn for each. It visits
 all flags, even those not set.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. FlagSet代表一个已注册的flag的集合。FlagSet零值没有名字，采用ContinueOnError错误处理策略;
 1. FlagSet的各个方法和flag包中函数功能基本一致，只不过方法接收方是底层的FlagSet结构体，不在赘述；
 2. Init方法设置flag集合f的名字和错误处理属性。FlagSet零值没有名字，默认采用ContinueOnError。
*************************************************************************************/
package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	fs := flag.NewFlagSet("myfs", flag.ContinueOnError)
	fmt.Println(fs)
	fs.Init("myFlagSet", flag.ExitOnError)
	fmt.Println(fs)
	fmt.Println("NFlag() =", fs.NFlag())
	fl := fs.Lookup("myFlagSet")
	if fl == nil {
		fmt.Println("Lookup() return nil")
	}
	fmt.Println("NArg() =", fs.NArg(), "Args() =", fs.Args(), "Arg(0) =", fs.Arg(0))
	println("----------------------")
	var (
		bb, bbb bool
		ff64    float64
		uui     uint
		uui64   uint64
		ii      int
		ii64    int64
		sstr    string
		dd      time.Duration
	)

	var b = fs.Bool("b", true, "-b=false")
	var f64 = fs.Float64("f64", 3.14, "-f64=3.14")
	var ui = fs.Uint("ui", 0, "-ui=0")
	var ui64 = fs.Uint64("ui64", 0, "-ui64=0")
	var i = fs.Int("i", 0, "-i=0")
	var i64 = fs.Int64("i64", 0, "-i64=0")
	var str = fs.String("str", "axe", "-str=axe")
	var d = fs.Duration("d", time.Second*1000, "-d=1000s")

	fs.BoolVar(&bb, "bb", false, "-bb=true")
	fs.BoolVar(&bbb, "bbb", false, "-bbb=false")
	fs.Float64Var(&ff64, "ff64", 3.14, "-ff64=3.14")
	fs.UintVar(&uui, "uui", 0, "-uui=0")
	fs.Uint64Var(&uui64, "uui64", 0, "-uui64=0")
	fs.IntVar(&ii, "ii", 0, "-ii=0")
	fs.Int64Var(&ii64, "ii64", 0, "-ii64=0")
	fs.StringVar(&sstr, "sstr", "axetang", "-sstr=axe")
	fs.DurationVar(&dd, "dd", time.Second*1000, "-dd=1000s")

	fmt.Println("fs Is Parsed?", fs.Parsed())
	fs.Parse([]string{
		"-b=false",
		"-bb=true",
		"-f64=1.11",
		"-ff64=2.222",
		"-ui=100",
		"-uui=101",
		"-ui64=111",
		"-uui64=888",
		"-i=999",
		"-ii=998",
		"-i64=1234",
		"-ii64=4321",
		"-str=axe",
		"-sstr=tang",
		"-d=1000s",
		"-dd=1000s",
		"a",
		"b",
		"c",
	})
	fs.SetOutput(os.Stdout)
	fs.PrintDefaults()
	println("----------Methods of FlagSet----------------")
	fmt.Println("b=", *b, "bb=", bb)
	fmt.Println("f64=", *f64, "ff64=", ff64)
	fmt.Println("ui=", *ui, "uui=", uui)
	fmt.Println("ui64=", *ui64, "uui64=", uui64)
	fmt.Println("i=", *i, "ii=", ii)
	fmt.Println("i64=", *i64, "ii64=", ii64)
	fmt.Println("str=", *str, "sstr=", sstr)
	fmt.Printf("d=%f, dd=%f\n", d.Seconds(), dd.Seconds())
	fmt.Println("fs Is Parsed?", fs.Parsed())
	fmt.Println("NArg() =", fs.NArg(), "Args() =", fs.Args(), "Arg(0) =", fs.Arg(0))

	println("--------Other Methods---------------")
	fmt.Println("fs NFlag is", fs.NFlag())
	fl = fs.Lookup("sstr")
	fmt.Println("sstr flag", fl.Name, fl.Usage, fl.DefValue, fl.Value)

	println("-------Visit&VisitAll--------------")
	fs.Visit(PrintFlag)
	fs.VisitAll(PrintFlag)

}

func PrintFlag(f *flag.Flag) {
	fmt.Printf("Name:%s, Usage:%s,Value:%s,DefValue:%s\n", f.Name, f.Usage, f.Value, f.DefValue)
}
