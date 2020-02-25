/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: template
 **Element: tempalte.
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Error struct {
    // ErrorCode describes the kind of error.
    ErrorCode ErrorCode
    // Node is the node that caused the problem, if known.
    // If not nil, it overrides Name and Line.
    Node parse.Node
    // Name is the name of the template in which the error was encountered.
    Name string
    // Line is the line number of the error in the template source or 0.
    Line int
    // Description is a human-readable description of the problem.
    Description string
 }
 func (e *Error) Error() string
 ------------------------------------------------------------------------------------
 **Description:
 Error describes a problem encountered during template Escaping.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Error结构体秒数了在template转义操作时发生的错误；
 2. Error方法返回错误信息字符串；
*************************************************************************************/
package main

func main() {
}
