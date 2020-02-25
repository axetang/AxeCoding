/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: expvar
 **Element: expvar.Int
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func NewInt(name string) *Int
 ------------------------------------------------------------------------------------
 **Description:
 Int is a 64-bit integer variable that satisfies the Var interface.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Int代表一个64位整数变量，满足Var接口；
 2. Int只有一个非导出成员i int64。
*************************************************************************************/
package main

import (
	"expvar"
	"fmt"
	"log"
	"net/http"
)

func kvFunc(kv expvar.KeyValue) {
	fmt.Println(kv.Key, kv.Value)
}

func main() {
	inerInt := int64(10)
	pubInt := expvar.NewInt("Int")
	pubInt.Set(inerInt)
	pubInt.Add(2)

	inerFloat := 1.2
	pubFloat := expvar.NewFloat("Float")
	pubFloat.Set(inerFloat)
	pubFloat.Add(0.1)

	inerString := "hello"
	pubString := expvar.NewString(inerString)
	pubString.Set(inerString)

	pubMap := expvar.NewMap("Map").Init()
	pubMap.Set("Int", pubInt)
	pubMap.Set("Float", pubFloat)
	pubMap.Set("String", pubString)
	fmt.Println("-----Print KeyValue----")
	pubMap.Do(kvFunc)
	fmt.Println("-----end----")

	pubMap.Add("Int", 1)
	pubMap.Add("NewInt", 123)
	pubMap.AddFloat("Float", 0.5)
	pubMap.AddFloat("NewFloat", 0.9)
	fmt.Println("-----Print KeyValue----")
	pubMap.Do(kvFunc)
	fmt.Println("-----end----")

	expvar.Do(kvFunc)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
