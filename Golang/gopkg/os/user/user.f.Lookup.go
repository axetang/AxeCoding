/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: user
 **Element: user.Lookup
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Lookup(username string) (*User, error)
 ------------------------------------------------------------------------------------
 **Description:
 Lookup looks up a user by username. If the user cannot be found, the returned error
 is of type UnknownUserError.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Lookup函数返回参数username string指定的用户名的用户信息*User
*************************************************************************************/
package main

import (
	"fmt"
	"os/user"
)

func main() {
	findedUser, err := user.Lookup("axe")
	if err != nil {
		fmt.Println(err)
	}
	//我所在环境为Mac os X 10.8.2
	fmt.Println("用户id:", findedUser.Uid)      //用户id: 501
	fmt.Println("用户所在组id:", findedUser.Gid)   //    用户所在组id: 20
	fmt.Println("用户名:", findedUser.Username)  //用户名: lsdev
	fmt.Println("用户全名:", findedUser.Name)     //用户全名: 吴 佳煌
	fmt.Println("用户家目录:", findedUser.HomeDir) //用户家目录: /Users/lsdev

}
