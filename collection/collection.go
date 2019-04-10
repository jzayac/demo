package collection

import (
	"errors"
	"os"
	"sync"
	"time"

	"github.com/go-kit/kit/log"

	"demo/cache"
	cc "demo/collection/cache"
	"demo/model"
)

type collection map[string]cache.Interface

var logger log.Logger

var instance *collection
var once sync.Once

var mu sync.Mutex

func (c collection) GetKeys() []string {
	keys := make([]string, 0, len(c))
	for key, _ := range c {
		keys = append(keys, key)
	}
	return keys
}

func (c collection) GetInstanceByKey(key string) (cache.Interface, error) {
	mu.Lock()
	defer mu.Unlock()
	i, ok := c[key]
	if !ok {
		return nil, errors.New("key not fond")
	}
	return i, nil
}

func GetInstance() *collection {
	once.Do(func() {
		logger = log.NewLogfmtLogger(os.Stderr)
		instance = newInstance()
	})

	return instance
}

func periodicallyUpdateCache(c cache.Interface) {
	sec := c.GetUpdateIterationInSeconds()

	for t := range time.NewTicker(sec * time.Second).C {
		mu.Lock()
		logger.Log("time_tick", t, "cache_name", c.GetName())

		err := c.UpdateCacheDataFromDb()

		if err != nil {
			logger.Log("message", "cache update fail: ", "err", err, "cache_name", c.GetName())
		}
		mu.Unlock()
	}
}

func newInstance() *collection {
	ins := &collection{}

	am := model.NewAirline()
	airline, _ := cc.NewCache("airline", 2, am)
	(*ins)[airline.GetName()] = airline

	go periodicallyUpdateCache(airline)

	cm := model.NewCity()
	city, _ := cc.NewCache("city", 2, cm)
	(*ins)[city.GetName()] = city

	go periodicallyUpdateCache(city)

	// TODO: exmaple for new cache but with same airline source
	// sm := model.NewAirline()
	// some, _ := cc.NewCache("some", 2, sm)
	// (*ins)[some.GetName()] = some

	// go periodicallyUpdateCache(some)

	logger.Log("msg", "collection initialized")
	return ins
}
