/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: user
 **Element: user.UnknownGroupError
 **Type: type
 ------------------------------------------------------------------------------------
 **Definition:
 type UnknownGroupError string
 func (e UnknownGroupError) Error() string
 ------------------------------------------------------------------------------------
 **Description:
 UnknownGroupError is returned by LookupGroup when a group cannot be found.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. UnknownGroupError在LookupGroup函数找不到一个相应的用户群组时返回；
 2. Error方法返回错误信息字符串。
*************************************************************************************/
package main

import (
	"fmt"
	"os/user"
)

func main() {
	var err user.UnknownGroupError
	fmt.Println(err.Error())

}
