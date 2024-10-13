package main

import (
	"fmt"
	"sort"
)

type KeyValue struct {
	Key   string
	Value string
}

type SortedKey []KeyValue

// Len 重写len,swap,less才能排序
func (k SortedKey) Len() int           { return len(k) }
func (k SortedKey) Swap(i, j int)      { k[i], k[j] = k[j], k[i] }
func (k SortedKey) Less(i, j int) bool { return k[i].Key < k[j].Key }

func test04() *KeyValue {
	kv := KeyValue{}
	return &kv
}

func main() {
	kv1 := KeyValue{Key: "abc",
		Value: "1",
	}
	kv2 := KeyValue{"bcd", "1"}
	kv3 := KeyValue{"abc", "1"}
	sk := SortedKey{kv1, kv2}
	sk = append(sk, kv3)

	fmt.Println(sk)
	sort.Sort(sk)
	fmt.Println(sk)

}
