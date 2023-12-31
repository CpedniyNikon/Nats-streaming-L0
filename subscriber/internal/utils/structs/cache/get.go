package cache

import (
	"fmt"
	"publisher/internal/models"
	_ "publisher/internal/models"
)

func (c *Cache) Get(orderId string) models.Order {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	if value, ok := c.cache[orderId]; ok {
		return value
	} else {
		fmt.Println("there is no such a record in table")
		return models.Order{}
	}
}
