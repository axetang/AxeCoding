/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.Chtimes
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Chtimes(name string, atime time.Time, mtime time.Time) error
  ------------------------------------------------------------------------------------
 **Description:
 Chtimes changes the access and modification times of the named file, similar to the
 Unix utime() or utimes() functions. The underlying filesystem may truncate or round
 the values to a less precise time unit. If there is an error, it will be of type
 *PathError.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Chtimes修改参数name指定的文件的atime和mtime；
 2. 出错返回*PathError。
*************************************************************************************/
package main

import (
	"log"
	"os"
	"time"
)

func main() {
	mtime := time.Date(2006, time.February, 1, 3, 4, 5, 0, time.UTC)
	atime := time.Date(2007, time.March, 2, 4, 5, 6, 0, time.UTC)
	if err := os.Chtimes("./_iofiles/Hello.go", atime, mtime); err != nil {
		log.Fatal(err)
	}
}
