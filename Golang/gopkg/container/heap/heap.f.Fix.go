/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: heap
 **Element: heap.Fix
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Fix(h Interface, i int)
 ------------------------------------------------------------------------------------
 **Description:
 Fix re-establishes the heap ordering after the element at index i has changed
 its value. Changing the value of the element at index i and then calling Fix is
 equivalent to, but less expensive than, calling Remove(h, i) followed by a Push
 of the new value. The complexity is O(log n) where n = h.Len().
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main

func main() {
}
