/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: os
 **Element: os.Signal
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type Signal interface {
     String() string
     Signal() // to distinguish from other Stringers
 }
 var (
     Interrupt Signal = syscall.SIGINT
     Kill      Signal = syscall.SIGKILL
 )
 ------------------------------------------------------------------------------------
 **Description:
 A Signal represents an operating system signal. The usual underlying implementation
 is operating system-dependent: on Unix it is syscall.Signal.
 The only signal values guaranteed to be present in the os package on all systems are
 Interrupt (send the process an interrupt) and Kill (force the process to exit).
 Interrupt is not implemented on Windows; using it with os.Process.Signal will return
 an error.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main
