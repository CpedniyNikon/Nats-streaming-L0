package cache

import (
	"publisher/internal/models"
	"sync"
)

type Cache struct {
	mutex sync.RWMutex
	cache map[string]models.Order
}

func NewCache() *Cache {
	return &Cache{
		mutex: sync.RWMutex{},
		cache: map[string]models.Order{},
	}
}
