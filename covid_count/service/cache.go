package service

import (
	"github.com/patrickmn/go-cache"
)

type Cache struct {
	Driver *cache.Cache
}


func (c *Cache) Set(key string, val interface{}, minutes float64){
	c.Driver.Set(key, val, cache.DefaultExpiration)
}

func (c *Cache) Get(key string) (interface{}, bool){
	return c.Driver.Get(key)
}

