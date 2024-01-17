package main

import (
	"fmt"
	"hash/fnv"
)

type data struct {
	key   string
	value interface{}
}

type hashTable struct {
	bucket [][]data
	size   int
}

func (ht *hashTable) createHash(key string) int {
	hash := fnv.New32()
	hash.Write([]byte(key))
	hashcode := hash.Sum32()
	return int(hashcode) % ht.size
}

func (ht *hashTable) insert(key string, value int) {
	index := ht.createHash(key)
	ht.bucket[index] = append(ht.bucket[index], data{key: key, value: value})
}

func (ht *hashTable) get(key string) {
	index := ht.createHash(key)
	fmt.Println(ht.bucket[index], "**")
}

func (ht *hashTable) delete(key string) {
	index := ht.createHash(key)
	ht.bucket[index] = ht.bucket[index][1:]
}

func main() {
	ht := hashTable{bucket: make([][]data, 6), size: 6}
	ht.insert("vajid", 3)
	ht.get("vajid")
	ht.delete("vajid")
	fmt.Println(ht.bucket)
}
