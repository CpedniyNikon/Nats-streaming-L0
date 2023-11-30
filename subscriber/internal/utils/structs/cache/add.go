package cache

import (
	"fmt"
	"publisher/internal/models"
)

func (c *Cache) Add(order models.Order) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if _, ok := c.cache[order.OrderUid]; ok {
		fmt.Println("there is already such a record in table")
	}
	c.cache[order.OrderUid] = order
}
