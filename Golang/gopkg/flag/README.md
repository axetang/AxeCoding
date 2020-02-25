# Golang package源码
这个目录下所有文件和子目录是Golang的标准库的相应标准包的源代码，供学习过程中查阅和参考q 
**标准库在线文档链接如下：**  
- [English Version](https://godoc.org/)
- [中文链接](http://docscn.studygolang.com/pkg/)
  
# package flag
## import "flag"

Package flag implements command-line flag parsing.

## Usage
```
Define flags using flag.String(), Bool(), Int(), etc.
```
This declares an integer flag, -flagname, stored in the pointer ip, with type *int.
```
import "flag"
var ip = flag.Int("flagname", 1234, "help message for flagname")
```
If you like, you can bind the flag to a variable using the Var() functions.

```
var flagvar int
func init() {
	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
}
```
Or you can create custom flags that satisfy the Value interface (with pointer receivers) and couple them to flag parsing by

```
flag.Var(&flagVal, "name", "help message for flagname")
```
For such flags, the default value is just the initial value of the variable.

After all flags are defined, call
```
flag.Parse()
```
to parse the command line into the defined flags.

Flags may then be used directly. If you're using the flags themselves, they are all pointers; if you bind to variables, they're values.

```
fmt.Println("ip has value ", *ip)
fmt.Println("flagvar has value ", flagvar)
```
After parsing, the arguments following the flags are available as the slice flag.Args() or individually as flag.Arg(i). The arguments are indexed from 0 through flag.NArg()-1.

## Command line flag syntax

The following forms are permitted:
```
-flag
-flag=x
-flag x  // non-boolean flags only
```
One or two minus signs may be used; they are equivalent. The last form is not permitted for boolean flags because the meaning of the command

```
cmd -x *
```
where * is a Unix shell wildcard, will change if there is a file called 0, false, etc. You must use the -flag=false form to turn off a boolean flag.

Flag parsing stops just before the first non-flag argument ("-" is a non-flag argument) or after the terminator "--".

Integer flags accept 1234, 0664, 0x1234 and may be negative. Boolean flags may be:

```
1, 0, t, f, T, F, true, false, TRUE, FALSE, True, False
```
Duration flags accept any input valid for time.ParseDuration.

The default set of command-line flags is controlled by top-level functions. The FlagSet type allows one to define independent sets of flags, such as to implement subcommands in a command-line interface. The methods of FlagSet are analogous to the top-level functions for the command-line flag set.

## Index
```
Variables
func Arg(i int) string
func Args() []string
func Bool(name string, value bool, usage string) *bool
func BoolVar(p *bool, name string, value bool, usage string)
func Duration(name string, value time.Duration, usage string) *time.Duration
func DurationVar(p *time.Duration, name string, value time.Duration, usage string)
func Float64(name string, value float64, usage string) *float64
func Float64Var(p *float64, name string, value float64, usage string)
func Int(name string, value int, usage string) *int
func Int64(name string, value int64, usage string) *int64
func Int64Var(p *int64, name string, value int64, usage string)
func IntVar(p *int, name string, value int, usage string)
func NArg() int
func NFlag() int
func Parse()
func Parsed() bool
func PrintDefaults()
func Set(name, value string) error
func String(name string, value string, usage string) *string
func StringVar(p *string, name string, value string, usage string)
func Uint(name string, value uint, usage string) *uint
func Uint64(name string, value uint64, usage string) *uint64
func Uint64Var(p *uint64, name string, value uint64, usage string)
func UintVar(p *uint, name string, value uint, usage string)
func UnquoteUsage(flag *Flag) (name string, usage string)
func Var(value Value, name string, usage string)
func Visit(fn func(*Flag))
func VisitAll(fn func(*Flag))
type ErrorHandling
type Flag
func Lookup(name string) *Flag
type FlagSet
func NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet
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
type Getter
type Value
```
## Package Files

flag.go

