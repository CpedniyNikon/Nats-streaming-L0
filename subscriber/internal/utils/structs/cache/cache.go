package cache

import "publisher/internal/models"

type Cache struct {
	cache map[string]models.Order
}

func NewCache() *Cache {
	return &Cache{
		cache: map[string]models.Order{},
	}
}
