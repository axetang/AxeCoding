/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: time
 **Element: time.Timer
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 type Timer struct {
     C <-chan Time
     // contains filtered or unexported fields
 }
 func (t *Timer) Reset(d Duration) bool
 func (t *Timer) Stop() bool
 ------------------------------------------------------------------------------------
 **Description:
 The Timer type represents a single event. When the Timer expires, the current time will
 be sent on C, unless the Timer was created by AfterFunc. A Timer must be created with
 NewTimer or AfterFunc.
 Reset changes the timer to expire after duration d. It returns true if the timer had
 been active, false if the timer had expired or been stopped.
 Resetting a timer must take care not to race with the send into t.C that happens when
 the current timer expires. If a program has already received a value from t.C, the
 timer is known to have expired, and t.Reset can be used directly. If a program has
 not yet received a value from t.C, however, the timer must be stopped and—if Stop
 reports that the timer expired before being stopped—the channel explicitly drained:
 if !t.Stop() {
 	<-t.C
 }
 t.Reset(d)
 This should not be done concurrent to other receives from the Timer's channel.
 Note that it is not possible to use Reset's return value correctly, as there is a race
 condition between draining the channel and the new timer expiring. Reset should always
 be invoked on stopped or expired channels, as described above. The return value exists
 to preserve compatibility with existing programs.
 Stop prevents the Timer from firing. It returns true if the call stops the timer,
 false if the timer has already expired or been stopped. Stop does not close the
 channel, to prevent a read from the channel succeeding incorrectly.
 To prevent a timer created with NewTimer from firing after a call to Stop, check the
 return value and drain the channel. For example, assuming the program has not received
 from t.C already:
 if !t.Stop() {
 	<-t.C
 }
 This cannot be done concurrent to other receives from the Timer's channel.
 For a timer created with AfterFunc(d, f), if t.Stop returns false, then the timer
 has already expired and the function f has been started in its own goroutine; Stop
 does not wait for f to complete before returning. If the caller needs to know
 whether f is completed, it must coordinate with f explicitly.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
