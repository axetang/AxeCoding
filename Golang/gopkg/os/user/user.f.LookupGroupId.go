/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: user
 **Element: user.LookupGroupId
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func LookupGroupId(gid string) (*Group, error)
 ------------------------------------------------------------------------------------
 **Description:
 LookupGroupId looks up a group by groupid. If the group cannot be found, the
 returned error is of type UnknownGroupIdError.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

import (
	"fmt"
	"os/user"
)

func main() {
	gr, err := user.LookupGroupId("20")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Gid:%s,Name:%s\n", gr.Gid, gr.Name)
}
