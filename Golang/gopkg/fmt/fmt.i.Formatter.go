/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: fmt
 **Element: fmt.Formatter
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type Formatter interface {
    Format(f State, c rune)
 }
 ------------------------------------------------------------------------------------
 **Description:
 Formatter is the interface implemented by values with a custom formatter. The
 implementation of Format may call Sprint(f) or Fprint(f) etc. to generate its
 output.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Formatter接口由带有定制的格式化器的值所实现;
 2. Format的实现可调用Sprintf或Fprintf(f)等函数来生成其输出;
 3. State接口定义如下：
 State represents the printer state passed to custom formatters.
 It provides access to the io.Writer interface plus information about the flags and
 options for the operand's format specifier.
 type State interface {
	// Write is the function to call to emit formatted output to be printed.
	Write(b []byte) (n int, err error)
	// Width returns the value of the width option and whether it has been set.
	Width() (wid int, ok bool)
	// Precision returns the value of the precision option and whether it has been set.
	Precision() (prec int, ok bool)
	// Flag reports whether the flag c, a character, has been set.
	Flag(c int) bool
 }
*************************************************************************************/
package main

func main() {
}
