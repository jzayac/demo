package collection

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"demo/cache"
	"demo/collection/airline"
	"demo/collection/city"
)

type collection map[string]cache.Interface

var instance *collection
var once sync.Once

func (c collection) GetKeys() []string {
	keys := make([]string, 0, len(c))
	for key, _ := range c {
		keys = append(keys, key)
	}
	return keys
}

func (c collection) GetInstanceByKey(key string) (cache.Interface, error) {
	i, ok := c[key]
	if !ok {
		return nil, errors.New("key not fond")
	}
	return i, nil
}

func GetInstance() *collection {
	once.Do(func() {
		instance = newInstance()
	})

	return instance
}

func periodicallyUpdateCache(c cache.Interface) {

	sec := c.GetUpdateIterationInSeconds()
	fmt.Println(sec)

	for t := range time.NewTicker(sec * time.Second).C {

		fmt.Println("time tick ", t, c.GetName())

		err := c.UpdateCacheDataFromDb()
		fmt.Println(err)

	}
}

func newInstance() *collection {
	ins := &collection{}

	airlineInstance := airline.GetInstance()
	(*ins)[airlineInstance.GetName()] = airlineInstance

	cityInstance := city.GetInstance()
	(*ins)[cityInstance.GetName()] = cityInstance

	go periodicallyUpdateCache(cityInstance)

	return ins
}
