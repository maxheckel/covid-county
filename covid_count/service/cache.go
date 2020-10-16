package service

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

type Cache struct {
	Driver *cache.Cache
}


func (c *Cache) Set(key string, val interface{}, minutes float64){
	duration, _ := time.ParseDuration(fmt.Sprintf("%gm", minutes))
	c.Driver.Set(key, val, duration)
}

func (c *Cache) Get(key string) (interface{}, bool){
	return c.Driver.Get(key)
}

