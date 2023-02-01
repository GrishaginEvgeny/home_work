package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mutex    sync.RWMutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type queueItem struct {
	Key   Key
	Value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	if _, isInMap := cache.items[key]; isInMap {
		cache.queue.MoveToFront(cache.items[key])
		cache.queue.Front().Value = &queueItem{
			Key:   key,
			Value: value,
		}
		cache.items[key].Value.(*queueItem).Value = value
		return true
	}
	if cache.queue.Len() == cache.capacity {
		delete(cache.items, cache.queue.Back().Value.(*queueItem).Key)
		cache.queue.Remove(cache.queue.Back())
	}
	cache.queue.PushFront(&queueItem{
		Key:   key,
		Value: value,
	})
	cache.items[key] = cache.queue.Front()
	return false
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	if _, isInMap := cache.items[key]; !isInMap {
		return nil, false
	}
	cache.queue.MoveToFront(cache.items[key])
	cache.items[key] = cache.queue.Front()
	return cache.items[key].Value.(*queueItem).Value, true
}

func (cache *lruCache) Clear() {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	cache.queue = NewList()
	cache.items = make(map[Key]*ListItem, cache.capacity)
}
