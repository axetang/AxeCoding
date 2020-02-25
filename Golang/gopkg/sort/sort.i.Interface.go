/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: sort
 **Element: sort.Interface
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type Interface interface {
     // Len is the number of elements in the collection.
     Len() int
     // Less reports whether the element with
     // index i should sort before the element with index j.
     Less(i, j int) bool
     // Swap swaps the elements with indexes i and j.
     Swap(i, j int)
 }
 ------------------------------------------------------------------------------------
 **Description:
 A type, typically a collection, that satisfies sort.Interface can be sorted by the
 routines in this package. The methods require that the elements of the collection
 be enumerated by an integer index.
 ------------------------------------------------------------------------------------
 **要点总结:
 1.一个满足sort.Interface接口的（集合）类型可以被本包的函数进行排序。方法要求集合中的元素可以被
 整数索引;
 2. Interface接口的三个方法：Len，Less, Swap。
*************************************************************************************/
package main

func main() {
}
