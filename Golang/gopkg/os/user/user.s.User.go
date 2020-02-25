/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: user
 **Element: user.User
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type User struct {
     // Uid is the user ID.
     // On POSIX systems, this is a decimal number representing the uid.
     // On Windows, this is a security identifier (SID) in a string format.
     // On Plan 9, this is the contents of /dev/user.
     Uid string
     // Gid is the primary group ID.
     // On POSIX systems, this is a decimal number representing the gid.
     // On Windows, this is a SID in a string format.
     // On Plan 9, this is the contents of /dev/user.
     Gid string
     // Username is the login name.
     Username string
     // Name is the user's real or display name.
     // It might be blank.
     // On POSIX systems, this is the first (or only) entry in the GECOS field
     // list.
     // On Windows, this is the user's display name.
     // On Plan 9, this is the contents of /dev/user.
     Name string
     // HomeDir is the path to the user's home directory (if they have one).
     HomeDir string
 }
 func (u *User) GroupIds() ([]string, error)
 ------------------------------------------------------------------------------------
 **Description:
 User represents a user account.
 GroupIds returns the list of group IDs that the user is a member of.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. User结构体包含五个成员：Uid，Gid，Username，Name和HomeDir；
 2. GroupIds方法返回用户所属群组的ID的字符串切片[]string。
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
	s, _ := user.GroupIds()
	fmt.Printf("GroupIds:%v\n", s) //用户id: 501

}
