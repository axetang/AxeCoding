# Golang package源码
这个目录下所有文件和子目录是Golang的标准库的相应标准包的源代码，供学习过程中查阅和参考q 
**标准库在线文档链接如下：**  
- [English Version](https://godoc.org/)
- [中文链接](http://docscn.studygolang.com/pkg/)
  
# package expvar
## import "expvar"

Package expvar provides a standardized interface to public variables, such as operation counters in servers. It exposes these variables via HTTP at /debug/vars in JSON format.

Operations to set or modify these public variables are atomic.

In addition to adding the HTTP handler, this package registers the following variables:

cmdline   os.Args
memstats  runtime.Memstats
The package is sometimes only imported for the side effect of registering its HTTP handler and the above variables. To use it this way, link this package into your program:

import _ "expvar"

## Index
```
func Do(f func(KeyValue))
func Handler() http.Handler
func Publish(name string, v Var)
type Float
func NewFloat(name string) *Float
func (v *Float) Add(delta float64)
func (v *Float) Set(value float64)
func (v *Float) String() string
func (v *Float) Value() float64
type Func
func (f Func) String() string
func (f Func) Value() interface{}
type Int
func NewInt(name string) *Int
func (v *Int) Add(delta int64)
func (v *Int) Set(value int64)
func (v *Int) String() string
func (v *Int) Value() int64
type KeyValue
type Map
func NewMap(name string) *Map
func (v *Map) Add(key string, delta int64)
func (v *Map) AddFloat(key string, delta float64)
func (v *Map) Delete(key string)
func (v *Map) Do(f func(KeyValue))
func (v *Map) Get(key string) Var
func (v *Map) Init() *Map
func (v *Map) Set(key string, av Var)
func (v *Map) String() string
type String
func NewString(name string) *String
func (v *String) Set(value string)
func (v *String) String() string
func (v *String) Value() string
type Var
func Get(name string) Var
```

## Package Files

expvar.go

  