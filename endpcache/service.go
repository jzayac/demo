package endpcache

import (
	"errors"

	"demo/cache"
	"demo/collection"
)

type Service interface {
	GetList(string) (*cache.Iata, error)
}

type service struct{}

func (sv service) GetList(name string) (*cache.Iata, error) {
	data, err := collection.GetInstance().GetInstanceByKey(name)
	if err != nil {
		return nil, ErrInvalidKey
	}
	list := data.GetList()
	return &list, nil
}

func NewService() Service {
	return &service{}
}

var ErrInvalidKey = errors.New("cache not found")
