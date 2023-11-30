package service

import "publisher/internal/utils/structs/cache"

type Service struct {
	cache *cache.Cache
}

func NewService(cache *cache.Cache) *Service {
	return &Service{cache: cache}
}
