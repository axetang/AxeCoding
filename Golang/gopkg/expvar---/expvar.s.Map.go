/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: expvar
 **Element: expvar.Map
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 // Map is a string-to-Var map variable that satisfies the Var interface.
 type Map struct {
    // contains filtered or unexported fields
	m      sync.Map // map[string]Var
	keysMu sync.RWMutex
	keys   []string // sorted
 }
 func (v *Map) Add(key string, delta int64)
 func (v *Map) AddFloat(key string, delta float64)
 func (v *Map) Delete(key string)
 func (v *Map) Do(f func(KeyValue))
 func (v *Map) Get(key string) Var
 func (v *Map) Init() *Map
 func (v *Map) Set(key string, av Var)
 func (v *Map) String() string
 ------------------------------------------------------------------------------------
 **Description:
 Map is a string-to-Var map variable that satisfies the Var interface.
 Add adds delta to the *Int value stored under the given map key.
 AddFloat adds delta to the *Float value stored under the given map key.
 Deletes the given key from the map.
 Do calls f for each entry in the map. The map is locked during the iteration,
 but existing entries may be concurrently updated.
 Init removes all keys from the map.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. Map代表一个string到Var的映射变量，满足Var接口;
 1. Add方法
 2. AddFloat方法
 3. Delete方法
 4. Do方法
 5. Get方法
 6. Init方法
 7. Set方法
 8. String方法
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
