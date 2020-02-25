/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.Executable
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Executable() (string, error)
 ------------------------------------------------------------------------------------
 **Description:
 Executable returns the path name for the executable that started the current process.
 There is no guarantee that the path is still pointing to the correct executable. If
 a symlink was used to start the process, depending on the operating system, the
 result might be the symlink or the path it pointed to. If a stable result is needed,
 path/filepath.EvalSymlinks might help.
 Executable returns an absolute path unless an error occurred.
 The main use case is finding resources located relative to an executable.
 Executable is not supported on nacl.
 ------------------------------------------------------------------------------------
 **要点总结:

*************************************************************************************/
package main
