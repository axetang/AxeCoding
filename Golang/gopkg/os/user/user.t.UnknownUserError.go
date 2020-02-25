/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: user
 **Element: user.UnknownUserError
 **Type: type
 ------------------------------------------------------------------------------------
 **Definition:
 type UnknownUserError string
 func (e UnknownUserError) Error() string
 ------------------------------------------------------------------------------------
 **Description:
 UnknownUserError is returned by Lookup when a user cannot be found.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. UnknownUserError在Lookup函数找不到一个相应的用户时返回；
 2. Error方法返回错误信息字符串。
*************************************************************************************/
package main

import (
	"fmt"
	"os/user"
)

func main() {
	var err user.UnknownUserError
	fmt.Println(err.Error())

}
