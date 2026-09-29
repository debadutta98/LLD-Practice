package singleton

import (
	"sync"
)

type configManager struct {
	data sync.Map
}

var (
	instance *configManager
	once     sync.Once
)

func GetConfigurationManager() *configManager {
	once.Do(func() {
		instance = &configManager{}
	})
	return instance
}

func (c *configManager) Get(key string) (any, bool) {
	return c.data.Load(key)
}

func (c *configManager) Set(key string, val any) {
	c.data.Store(key, val)
}
