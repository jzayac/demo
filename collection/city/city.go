package city

import (
	"sync"
	"time"

	"demo/cache"
	"demo/model"
)

type Attr struct {
	Value string `json:"value"`
	Index string `json:"index"`
}

type cityIata cache.Iata

type city struct {
	data         cityIata
	mi           cache.ModelInterface
	name         string
	timeDuration time.Duration
}

var instance *city
var once sync.Once

var mu sync.Mutex

func (c city) Len() int {
	return len(c.data)
}

func (c city) GetList() cache.Iata {
	return cache.Iata(c.data)
}

func (c city) GetName() string {
	return c.name
}

func (c city) GetUpdateIterationInSeconds() time.Duration {
	return c.timeDuration
}

func (c *city) UpdateCacheDataFromDb() error {
	list, err := c.mi.GetList()
	if err != nil {
		return err
	}
	mu.Lock()
	defer mu.Unlock()
	c.data = cityIata(list)

	return nil
}

/**
 *  fetch data by interface
 */
func newCity(mi cache.ModelInterface) (*city, error) {
	list, err := mi.GetList()
	if err != nil {
		// LOGGER
		return &city{
			data:         make(cityIata),
			mi:           mi,
			name:         "city",
			timeDuration: 5,
		}, err
	}

	data := cityIata(list)
	return &city{
		data:         data,
		mi:           mi,
		name:         "city",
		timeDuration: 5,
	}, nil
}

func GetInstance() *city {
	once.Do(func() {
		cm := model.NewCity()
		list, err := newCity(cm)
		if err != nil {
			instance = &city{}
		} else {
			instance = list
		}
	})
	return instance
}
