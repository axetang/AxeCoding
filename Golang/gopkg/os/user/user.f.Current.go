/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: user
 **Element: user.Current
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Current() (*User, error)
 ------------------------------------------------------------------------------------
 **Description:
 Current returns the current user.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Current函数返回当前user的全部信息*User；
 2. User结构体参见相关代码。
*************************************************************************************/
package main

import (
	"fmt"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}
	//我所在环境为Mac os X 10.8.2
	fmt.Println("用户id:", user.Uid)      //用户id: 501
	fmt.Println("用户所在组id:", user.Gid)   //    用户所在组id: 20
	fmt.Println("用户名:", user.Username)  //用户名: lsdev
	fmt.Println("用户全名:", user.Name)     //用户全名: 吴 佳煌
	fmt.Println("用户家目录:", user.HomeDir) //用户家目录: /Users/lsdev

}
