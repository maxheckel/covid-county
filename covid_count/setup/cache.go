package setup

import (
	"github.com/maxheckel/covid_county/covid_count/service"
	"github.com/patrickmn/go-cache"
)

func NewCache(c *cache.Cache) service.Cache{
	return service.NewCacheService(c)
}
