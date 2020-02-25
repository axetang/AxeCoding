/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: url
 **Element: url.Values
 **Type: type
 ------------------------------------------------------------------------------------
 **Definition:
 type Values map[string][]string
 func (v Values) Add(key, value string)
 func (v Values) Del(key string)
 func (v Values) Encode() string
 func (v Values) Get(key string) string
 func (v Values) Set(key, value string)
 ------------------------------------------------------------------------------------
 **Description:
 Values maps a string key to a list of values. It is typically used for query
 parameters and form values. Unlike in the http.Header map, the keys in a Values
 map are case-sensitive.
 Query is expected to be a list of key=value settings separated by ampersands or
 semicolons. A setting without an equals sign is interpreted as a key set to an empty
 value.
 Add adds the value to key. It appends to any existing values associated with key.
 Del deletes the values associated with key.
 Encode encodes the values into “URL encoded” form ("bar=baz&foo=quux") sorted by key.
 Get gets the first value associated with the given key. If there are no values
 associated with the key, Get returns the empty string. To access multiple values,
 use the map directly.
 Set sets the key to value. It replaces any existing values.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. Values将字符串键映射到字符串值切片的Map。它一般用于查询的参数和表单的属性。不同于
 http.Header这个字典类型，Values的键是大小写敏感的；
 1.Add方法把key-value对加入v Value；
 2. Del方法删除与key对应的所有values；
 3. Encode方法对v中的values进行编码，格式如("bar=baz&foo=quux"),编码时会以键进行排序;
 4. Get方法获得给定key对应的第一个value，如果找不到，Get方法返回空串。要访问多个值，直接使用var的map即可；
 5. Set方法给指定的key设置value，它会取代任何已经存在的values。
*************************************************************************************/
package main

import (
	"fmt"
	"net/url"
)

func main() {
	v := url.Values{}
	v.Set("name", "Ava")
	v.Add("friend", "Jess")
	v.Add("friend", "Sarah")
	v.Add("friend", "Zoe")
	// v.Encode() == "name=Ava&friend=Jess&friend=Sarah&friend=Zoe"
	fmt.Println("Encode():", v.Encode())
	fmt.Println(v.Get("name"))
	fmt.Println(v.Get("friend"))
	fmt.Println(v["friend"])
	v.Del("name")
	fmt.Println("del name", v)
	v.Set("friend", "axe")
	fmt.Println("set friend axe", v)
}
