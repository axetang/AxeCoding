/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.UserCacheDir
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func UserCacheDir() (string, error)
 ------------------------------------------------------------------------------------
 **Description:
 UserCacheDir returns the default root directory to use for user-specific cached data.
 Users should create their own application-specific subdirectory within this one and
 use that.
 On Unix systems, it returns $XDG_CACHE_HOME as specified by
 https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html
 if non-empty, else $HOME/.cache. On Darwin, it returns $HOME/Library/Caches.
 On Windows, it returns %LocalAppData%. On Plan 9, it returns $home/lib/cache.
 If the location cannot be determined (for example, $HOME is not defined), then it
 will return an error.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. UserCacheDir函数返回指定用户存储cache数据的目录；
 2. Unix系统下，返回$XDG_CACHE_HOME;
 3. Mac系统下，返回$HOME/Library/Caches;
 4. Windows, %LocalAppData%。
*************************************************************************************/
package main

import (
	"fmt"
	"os"
)

func main() {
	s, _ := os.UserCacheDir()
	fmt.Printf("UserCacheDir is: %s\n", s)
}
