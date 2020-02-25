/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: hash
 **Element: hash.Hash
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type Hash interface {
     // Write (via the embedded io.Writer interface) adds more data to the running hash.
     // It never returns an error.
     io.Writer

     // Sum appends the current hash to b and returns the resulting slice.
     // It does not change the underlying hash state.
     Sum(b []byte) []byte

     // Reset resets the Hash to its initial state.
     Reset()

     // Size returns the number of bytes Sum will return.
     Size() int

     // BlockSize returns the hash's underlying block size.
     // The Write method must be able to accept any amount
     // of data, but it may operate more efficiently if all writes
     // are a multiple of the block size.
     BlockSize() int
 }
 ------------------------------------------------------------------------------------
 **Description:
 Hash is the common interface implemented by all hash functions.
 Hash implementations in the standard library (e.g. hash/crc32 and crypto/sha256)
 implement the encoding.BinaryMarshaler and encoding.BinaryUnmarshaler interfaces.
 Marshaling a hash implementation allows its internal state to be saved and used
 for additional processing later, without having to re-write the data previously
 written to the hash. The hash state may contain portions of the input in its
 original form, which users are expected to handle for any possible security
 implications.
 Compatibility: Any future changes to hash or crypto packages will endeavor to
 maintain compatibility with state encoded using previous versions. That is, any
 released versions of the packages should be able to decode data written with any
 previously released version, subject to issues such as security fixes. See the Go
 compatibility document for background: https://golang.org/doc/go1compat
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Hash是一个被所有hash函数实现的公共接口;
 2. Hash接口定义如下
 type Hash interface {
    // 通过嵌入的匿名io.Writer接口的Write方法向hash中添加更多数据，永远不返回错误
    io.Writer
    // 返回添加b到当前的hash值后的新切片，不会改变底层的hash状态
    Sum(b []byte) []byte
    // 重设hash为无数据输入的状态
    Reset()
    // 返回Sum会返回的切片的长度
    Size() int
    // 返回hash底层的块大小；Write方法可以接受任何大小的数据，
    // 但提供的数据是块大小的倍数时效率更高
    BlockSize() int
 }
 3. Sum方法作用和使用方法还不清楚，特别对于fnv.New128和fnv.New128a函数返回的是hash.Hash接口
 类型，如何计算Sum还不清楚，后续更新。
*************************************************************************************/
package main

import (
	"errors"
	"flag"
	"fmt"
	"hash"
	"hash/adler32"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"
	"io"
	"os"
)

var ErrInvalidHashType = errors.New("hashtool: invalid hash type")

func handleError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stdout, "err=", err.Error())
	}
}

func hashFile(name string, h interface{}, r io.Reader) error {
	w, ok := h.(hash.Hash)
	if ok {
		_, err := io.Copy(w, r)
		if err != nil {
			return err
		}

		switch hashType := h.(type) {
		case hash.Hash32:
			{
				result := hashType.Sum32()
				fmt.Fprintf(os.Stdout, "%s = 0x%x\n", name, result)
			}
		case hash.Hash64:
			{
				result := hashType.Sum64()
				fmt.Fprintf(os.Stdout, "%s = 0x%x\n", name, result)
			}
		default:
			{
				return ErrInvalidHashType
			}
		}
	}
	return nil
}

func main() {
	fmt.Println("Start")
	var hashString string
	var filePath string
	flag.StringVar(&hashString, "hash", "all", "hash type")
	flag.StringVar(&filePath, "filepath", "", "the target file path")
	fmt.Println("Before Parse")
	flag.Parse()
	fmt.Println("After Parse, filepath is ", filePath)
	if len(filePath) == 0 {
		flag.PrintDefaults()
		return
	}
	fmt.Fprintf(os.Stdout, "%s\n", filePath)
	f, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println("OpenFile")
		handleError(err)
		return
	}
	defer f.Close()

	hashMap := make(map[string]interface{})
	hashMap["adler32"] = adler32.New()
	hashMap["crc32"] = crc32.NewIEEE()
	hashMap["crc64"] = crc64.New(crc64.MakeTable(crc64.ISO))
	hashMap["fnv32"] = fnv.New32()
	hashMap["fnv32a"] = fnv.New32a()
	hashMap["fnv64"] = fnv.New64()
	hashMap["fnv64a"] = fnv.New64a()

	if hashString == "all" {
		for k, h := range hashMap {
			f.Seek(0, os.SEEK_SET)
			err := hashFile(k, h, f)
			if err != nil {
				handleError(err)
				break
			}
		}
	} else {
		h, ok := hashMap[hashString]
		if !ok {
			handleError(ErrInvalidHashType)
		} else {
			hashFile(hashString, h, f)
		}
	}

}
