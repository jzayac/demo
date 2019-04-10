package cache

import "time"

type Iata map[string]string

type ModelInterface interface {
	GetList() (map[string]string, error)
}

type Interface interface {
	Len() int
	// Update(interface{})
	UpdateCacheDataFromDb() error
	GetList() Iata
	GetName() string
	GetUpdateIterationInSeconds() time.Duration
}
