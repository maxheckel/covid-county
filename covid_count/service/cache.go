package service

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

type Cache interface{
	Set(key string, val interface{}, minutes float64)
	Get(key string) (interface{}, bool)
}

func NewCacheService(driver *cache.Cache) Cache{
	return &cacheService{
		Driver: driver,
	}
}

type cacheService struct {
	Driver *cache.Cache
}


func (c *cacheService) Set(key string, val interface{}, minutes float64){
	duration, _ := time.ParseDuration(fmt.Sprintf("%gm", minutes))
	c.Driver.Set(key, val, duration)
}

func (c *cacheService) Get(key string) (interface{}, bool){
	return c.Driver.Get(key)
}

