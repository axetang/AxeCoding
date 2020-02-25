/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: expvar
 **Element: expvar.Handler
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func Handler() http.Handler
 ------------------------------------------------------------------------------------
 **Description:
 Handler returns the expvar HTTP Handler.
 This is only needed to install the handler in a non-standard location.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. KeyValue代表Map中的一条记录:键值对.
*************************************************************************************/
package main

import (
	"expvar"
	"fmt"
	"net/http"
)

var visits = expvar.NewInt("visits")

func handler(w http.ResponseWriter, r *http.Request) {
	visits.Add(1)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
