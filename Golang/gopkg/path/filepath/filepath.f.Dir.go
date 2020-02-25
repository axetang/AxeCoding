/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: filepath
 **Element: filepath.Dir
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Dir(path string) string
 -----------------------------------------------------------------------------------
 **Description:
 Dir returns all but the last element of path, typically the path's directory.
 After dropping the final element, Dir calls Clean on the path and trailing
 slashes are removed. If the path is empty, Dir returns ".". If the path
 consists entirely of separators, Dir returns a single separator. The returned
 path does not end in a separator unless it is the root directory.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Dir函数和path.Dir函数功能用法一致，参见package path.
*************************************************************************************/
package main

func main() {
}
