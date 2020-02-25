/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: user
 **Element: user.LookupGroup
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func LookupGroup(name string) (*Group, error)
 ------------------------------------------------------------------------------------
 **Description:
 LookupGroup looks up a group by name. If the group cannot be found, the returned
 error is of type UnknownGroupError.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

import (
	"fmt"
	"os/user"
)

func main() {
	gr, err := user.LookupGroup("staff")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Gid", gr.Gid)
	fmt.Println("Name", gr.Name)
}
