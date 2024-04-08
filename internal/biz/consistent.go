package biz

import (
	"fmt"
	"hash/crc32"
	"master/internal/conf"
	"sort"
	"strconv"
	"sync"
)

type ConsistentHash struct {
	keys     []uint32
	replicas int
	hash     func(data []byte) uint32
	sync.RWMutex
	hashMap map[uint32]int64
}

func NewConsistentHash(conf *conf.Hash) *ConsistentHash {
	c := &ConsistentHash{
		replicas: int(conf.Replicas),
		hashMap:  make(map[uint32]int64),
	}
	switch conf.HashFn {
	case "crc32":
		c.hash = crc32.ChecksumIEEE
	default:
		c.hash = crc32.ChecksumIEEE
	}
	return c
}

func (c *ConsistentHash) Add(nodeID ...int64) {
	c.Lock()
	for _, key := range nodeID {
		for i := 0; i < c.replicas; i++ {
			hash := c.hash([]byte(strconv.FormatInt(key, 16) + strconv.Itoa(i)))
			c.hashMap[hash] = key
			c.keys = append(c.keys, hash)
		}
	}
	c.Unlock()
	sort.Slice(c.keys, func(i, j int) bool {
		return c.keys[i] < c.keys[j]
	})
}

func (c *ConsistentHash) GetN(key int64, n int) ([]int64, error) {
	c.RLock()
	defer c.RUnlock()

	if len(c.keys) == 0 {
		return nil, fmt.Errorf("empty nodes")
	}
	hash := c.hash([]byte(strconv.FormatInt(key, 16)))
	idx := sort.Search(len(c.keys), func(i int) bool {
		return c.keys[i] > hash
	})
	if idx == len(c.keys) {
		idx = 0
	}
	ids := make([]int64, 0, n)
	for i := 0; i < n; i++ {
		if idx == len(c.keys) {
			idx = 0
		}
		ids = append(ids, c.hashMap[c.keys[idx]])
		idx++
	}
	return ids, nil
}

func (c *ConsistentHash) Remove(nodeID ...int64) {
	c.Lock()
	for _, key := range nodeID {
		for i := 0; i < c.replicas; i++ {
			hash := c.hash([]byte(strconv.FormatInt(key, 16) + strconv.Itoa(i)))
			delete(c.hashMap, hash)
			for i, k := range c.keys {
				if k == hash {
					c.keys = append(c.keys[:i], c.keys[i+1:]...)
				}
			}
		}
	}
	c.Unlock()
}

func (c *ConsistentHash) Keys() []int64 {
	c.RLock()
	defer c.RUnlock()

	var keys []int64
	for _, key := range c.hashMap {
		keys = append(keys, key)
	}
	return keys
}
