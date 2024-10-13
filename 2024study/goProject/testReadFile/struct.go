package main

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

//type KeyValue struct {
//	Key   string
//	Value string
//}

// ByKey for sorting by key.
type ByKey []KeyValue

// Len for sorting by key.
func (a ByKey) Len() int           { return len(a) }
func (a ByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByKey) Less(i, j int) bool { return a[i].Key < a[j].Key }

func ihash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32() & 0x7fffffff)
}

func Map(filename string, contents string) []KeyValue {
	// function to detect word separators.
	ff := func(r rune) bool { return !unicode.IsLetter(r) }

	// split contents into an array of words.
	// strings.FieldsFunc(,)将满足函数ff的内容返回，组成切片
	words := strings.FieldsFunc(contents, ff)

	var kva []KeyValue
	//kv is a struct
	for _, w := range words {
		kv := KeyValue{w, "1"}
		kva = append(kva, kv)
	}
	return kva
}
func Reduce(key string, values []string) string {
	// return the number of occurrences of this word.
	return strconv.Itoa(len(values))
}

func test(kv *KeyValue) *KeyValue {
	kv.Value = "change"
	return kv
}

func main() {
	fileName := "D:\\GitRepository\\2024study\\goProject\\testReadFile\\pg-being_ernest.txt"
	toWrite := "D:\\GitRepository\\2024study\\goProject\\testReadFile\\pg-being_ernest_json.txt"
	file, err := os.Open(fileName)
	toWritefile, _ := os.Create(toWrite)
	//dir, err := os.Getwd()
	//dir = D:\GitRepository\2024study\goProject
	if err != nil {
		fmt.Println("failed to open : ", err)
	}
	content, err1 := io.ReadAll(file)
	if err1 != nil {
		fmt.Println("failed to read content")
	}
	keyValues := Map(fileName, string(content))
	encoder := json.NewEncoder(toWritefile)
	for _, kv := range keyValues {
		err := encoder.Encode(&kv)
		//fmt.Println(err)
		if err == nil {
			continue
		}
	}
	//fmt.Println("12" < "13")

}
