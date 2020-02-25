/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: rand
 **Element: rand.Source
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type Source interface {
    Int63() int64
    Seed(seed int64)
 }
 ------------------------------------------------------------------------------------
 **Description:
 A Source represents a source of uniformly-distributed pseudo-random int64 values in
 the range [0, 1<<63).
 ------------------------------------------------------------------------------------
 **要点总结:
 0. Source接口表示一个一致分布伪随机int64的值，范围是[0,1<<63]；包含两个成员方法Int63()和Seed();
 1.
*************************************************************************************/
package main

func main() {
}
