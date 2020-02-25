/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: user
 **Element: user.UnknownGroupIdError
 **Type: type
 ------------------------------------------------------------------------------------
 **Definition:
 type UnknownGroupIdError string
 func (e UnknownGroupIdError) Error() string
 ------------------------------------------------------------------------------------
 **Description:
 UnknownGroupIdError is returned by LookupGroupId when a group cannot be found.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. UnknownGroupIdError在LookupGroupId函数找不到一个相应的用户时返回；
 2. Error方法返回错误信息字符串。
*************************************************************************************/
package main

import (
	"fmt"
	"os/user"
)

func main() {
	var err user.UnknownGroupIdError
	fmt.Println(err.Error())
}
