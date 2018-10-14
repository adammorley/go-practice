package main

import (
	"errors"
	"fmt"
	"hash"
	"hash/fnv"
	"log"
)

const defaultBuckets = 8

type hashedKey uint32
type kvPair struct {
	key, value string
}
type hashMap struct {
	buckets      [][]kvPair
	numBuckets   uint32
	hashFunction hash.Hash32
}

func NewHashMap() *hashMap {
	return NewHashMapMoreBuckets(defaultBuckets)
}
func NewHashMapMoreBuckets(n uint32) *hashMap {
	if n < 1 {
		log.Fatal("too few buckets")
	}
	h := new(hashMap)
	h.numBuckets = n
	h.buckets = make([][]kvPair, n)
	var i uint32
	for ; i < n; i++ {
		h.buckets[i] = make([]kvPair, 0)
	}
	h.hashFunction = fnv.New32()
	return h
}
func (h *hashMap) hashKey(key string) hashedKey {
	h.hashFunction.Write([]byte(key)) // XXX utf-8 support not functional
	r := hashedKey(h.hashFunction.Sum32())
	h.hashFunction.Reset()
	return r
}
func (h *hashMap) pickBucket(hk hashedKey) uint32 {
	return uint32(hk) % h.numBuckets
}
func (h *hashMap) getKvPair(key string) (kvPair, error) {
	hk := h.hashKey(key)
	b := h.pickBucket(hk)
	for _, kv := range h.buckets[b] {
		if kv.key == key {
			return kv, nil
		}
	}
	return kvPair{"", ""}, errors.New("not found")
}
func (h *hashMap) get(key string) (string, error) {
	kv, e := h.getKvPair(key)
	if e != nil {
		return "", e
	}
	return kv.value, nil
}
func (h *hashMap) put(key, value string) {
	kv, e := h.getKvPair(key)
	if e != nil {
		b := h.pickBucket(h.hashKey(key))
		h.buckets[b] = append(h.buckets[b], kvPair{key: key, value: value})
	} else { // found, replace
		if kv.value != value { // value doesn't match, update
			b := h.pickBucket(h.hashKey(key))
			for i, t := range h.buckets[b] {
				if t.key == key {
					h.buckets[b][i].value = value
				}
			}
		} // else do nothing
	}

}

func main() {
	h := NewHashMap()
	h.put("cats", "5551212")
	h.put("dogs", "5551213")
	r, e := h.get("cats")
	fmt.Println(h.buckets)
	if e == nil {
		log.Print("cats is ", r)
	} else {
		log.Fatal(e)
	}
	r, e = h.get("zebras")
	if e == nil {
		log.Fatal("we should no find zebras")
	} else {
		log.Print("did not find zebras, good")
	}
}
