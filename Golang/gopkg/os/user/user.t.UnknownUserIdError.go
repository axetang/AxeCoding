/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: user
 **Element: user.UnknownUserIdError
 **Type: type
 ------------------------------------------------------------------------------------
 **Definition:
 type UnknownUserIdError int
 func (e UnknownUserIdError) Error() string
 ------------------------------------------------------------------------------------
 **Description:
 UnknownUserIdError is returned by LookupId when a user cannot be found.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. UnknownUserIdError在LookupID函数找不到一个相应的用户时返回；
 2. Error方法返回错误信息字符串。
*************************************************************************************/
package main

import (
	"fmt"
	"os/user"
)

func main() {
	var err user.UnknownUserIdError
	fmt.Println(err.Error())

}
