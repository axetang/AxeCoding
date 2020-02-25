/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: hash
 **Element: hash.Hash64
 **Type: interface
 ------------------------------------------------------------------------------------
 **Definition:
 type Hash64 interface {
     Hash
     Sum64() uint64
 }
 ------------------------------------------------------------------------------------
 **Description:
 Hash64 is the common interface implemented by all 64-bit hash functions.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. Hash64是一个被所有64位hash函数实现的公共接口。
 2. Sum64方法用于计算Hash64校验和。
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
