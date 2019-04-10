package cache

import (
	"sync"
	"time"

	cm "demo/cache"
)

type Attr struct {
	Value string `json:"value"`
	Index string `json:"index"`
}

type cacheIata cm.Iata

type cache struct {
	data         cacheIata
	mi           cm.ModelInterface
	name         string
	timeDuration time.Duration
	mu           sync.Mutex
}

var instance *cache

func (c cache) Len() int {
	return len(c.data)
}

func (c cache) GetList() cm.Iata {
	return cm.Iata(c.data)
}

func (c cache) GetName() string {
	return c.name
}

func (c cache) GetUpdateIterationInSeconds() time.Duration {
	return c.timeDuration
}

func (c *cache) UpdateCacheDataFromDb() error {
	list, err := c.mi.GetList()
	if err != nil {
		return err
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = cacheIata(list)

	return nil
}

/**
 *  fetch data by interface
 */
func NewCache(name string, timeDuration time.Duration, mi cm.ModelInterface) (*cache, error) {
	list, err := mi.GetList()
	if err != nil {
		return &cache{
			data:         make(cacheIata),
			mi:           mi,
			name:         name,
			timeDuration: timeDuration,
		}, err
	}

	data := cacheIata(list)
	return &cache{
		data:         data,
		mi:           mi,
		name:         name,
		timeDuration: timeDuration,
	}, nil
}
