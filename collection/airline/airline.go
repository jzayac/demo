package airline

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

// type airline cache.Iata

type airlineIata cache.Iata

type airline struct {
	data airlineIata
	mi   cache.ModelInterface
}

var instance *airline
var once sync.Once

var mu sync.Mutex

func (c airline) Len() int {
	return len(c.data)
}

func (c airline) GetList() cache.Iata {
	return cache.Iata(c.data)
}

func (c airline) GetName() string {
	return "airline"
}

func (c airline) GetUpdateIterationInSeconds() time.Duration {
	return 2
}

func (c *airline) UpdateCacheDataFromDb() error {
	list, err := c.mi.GetList()
	if err != nil {
		return err
	}
	mu.Lock()
	defer mu.Unlock()
	c.data = airlineIata(list)

	return nil
}

func newAirline(mi cache.ModelInterface) (*airline, error) {
	list, err := mi.GetList()
	if err != nil {
		// LOGGER
		return &airline{
			data: make(airlineIata),
			mi:   mi,
		}, err
	}

	data := airlineIata(list)
	return &airline{
		data: data,
		mi:   mi,
	}, nil
}

func GetInstance() *airline {
	once.Do(func() {
		am := model.NewAirline()
		list, err := newAirline(am)
		if err != nil {
			instance = &airline{}
		} else {
			instance = list
		}
	})
	return instance
}
