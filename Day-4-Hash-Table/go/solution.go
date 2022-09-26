package main

import "fmt"

type DataItem struct {
	key   int
	value int
}

type HashTable struct {
	size int
	data []DataItem
}

func (h *HashTable) Hash(key int) int {
	return key % h.size
}

func (h *HashTable) Insert(key int, value int) {
	dataItem := DataItem{key, value}
	index := h.Hash(key)
	for h.data[index].key != 0 {
		index++
		index %= h.size
	}
	h.data[index] = dataItem
}

func (h *HashTable) Search(key int) int {
	index := h.Hash(key)
	for h.data[index].key != 0 {
		if h.data[index].key == key {
			return h.data[index].value
		}
		index++
		index %= h.size
	}
	return -1
}

func (h *HashTable) Delete(key int) int {
	index := h.Hash(key)
	for h.data[index].key != 0 {
		if h.data[index].key == key {
			data := h.data[index]
			h.data[index] = DataItem{}
			return data.value
		}
		index++
		index %= h.size
	}
	return -1
}

func main() {
	hashTable := HashTable{3, []DataItem{{}, {}, {}}}
	hashTable.Insert(1, 1)
	hashTable.Insert(2, 2)
	hashTable.Insert(3, 3)
	fmt.Println(hashTable.Search(1))
	fmt.Println(hashTable.Search(2))
	fmt.Println(hashTable.Search(3))
	fmt.Println(hashTable.Delete(1))
	fmt.Println(hashTable.Delete(2))
	fmt.Println(hashTable.Delete(3))
	fmt.Println(hashTable.Delete(4))
}
