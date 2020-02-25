/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.OpenFile
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func OpenFile(name string, flag int, perm FileMode) (*File, error)
 ------------------------------------------------------------------------------------
 **Description:
 OpenFile is the generalized open call; most users will use Open or Create instead. It
 opens the named file with specified flag (O_RDONLY etc.) and perm (before umask), if
 applicable. If successful, methods on the returned File can be used for I/O. If there
 is an error, it will be of type *PathError.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. OpenFile方法打开一个文件，打开的方式和权限通过参数flag和perm控制；
*************************************************************************************/
package main

import (
	"fmt"
	"os"
)

func main() {
	fi, err := os.OpenFile("./_iofiles/Hello.go", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer fi.Close()
	data := make([]byte, 100)
	fi.Read(data)
	fmt.Printf("The %s's content is: \n%s \n", fi.Name(), data)
	fi.WriteString("come on!!\n")
	fi.Seek(0, 0)
	fi.Read(data)
	fmt.Printf("Now the %s's content is:\n%s \n", fi.Name(), data)
}
