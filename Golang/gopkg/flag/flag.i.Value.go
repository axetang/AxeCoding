/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: flag
 **Element: flag.Value
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type Value interface {
    String() string
    Set(string) error
 }
 ------------------------------------------------------------------------------------
 **Description:
 Value is the interface to the dynamic value stored in a flag. (The default
   value is represented as a string.)
 If a Value has an IsBoolFlag() bool method returning true, the command-line
 parser makes -name equivalent to -name=true rather than using the next
 command-line argument.
 Set is called once, in command line order, for each flag present. The flag
 package may call the String method with a zero-valued receiver, such as a
 nil pointer.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Value是一个接口，包含了两个方法String和Set；

*************************************************************************************/
package main

type MyValue struct {
	s string
}

func (mv *MyValue) String() string {
	return mv.s
}
func (mv *MyValue) Set(s string) error {
	mv.s = s
	return nil
}
func main() {
  
}
