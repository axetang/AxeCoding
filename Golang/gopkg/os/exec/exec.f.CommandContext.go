/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: exec
 **Element: exec.CommandContext
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func CommandContext(ctx context.Context, name string, arg ...string) *Cmd
 ------------------------------------------------------------------------------------
 **Description:
 CommandContext is like Command but includes a context.
 The provided context is used to kill the process (by calling os.Process.Kill) if the
 context becomes done before the command completes on its own.
 CombinedOutput runs the command and returns its combined standard output and standard
 error.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. CommandContext函数和Command函数功能类似，但包含一个上下文context；
 2. 当参数ctx context.Context代表的上下文在命令完成前已经完成，则调用它杀死进程；
 3. 此函数使用方法要再查阅源代码更新。
*************************************************************************************/
package main

import (
	"context"
	"os/exec"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	if err := exec.CommandContext(ctx, "sleep", "5").Run(); err != nil {
		// This will fail after 100 milliseconds. The 5 second sleep
		// will be interrupted.
	}
}
