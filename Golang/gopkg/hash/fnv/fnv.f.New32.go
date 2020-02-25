/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: fnv
 **Element: fnv.New32
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func New32() hash.Hash32
 ------------------------------------------------------------------------------------
 **Description:
 New32 returns a new 32-bit FNV-1 hash.Hash. Its Sum method will lay the value out
 in big-endian byte order.
 ------------------------------------------------------------------------------------
 **要点总结:
*************************************************************************************/
package main


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