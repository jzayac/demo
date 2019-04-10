package model

import (
	"demo/util/database"
)

type cityModel struct{}

func NewCity() cityModel {
	ins := cityModel{}
	return ins
}

func (c cityModel) GetList() (map[string]string, error) {
	client, err := database.RedisNewClient()

	if err != nil {
		return nil, err
	}

	citySlc, err := client.Keys("*").Result()
	if err != nil {
		return nil, err
	}

	data := make(map[string]string)
	// assuming that data in redis db is just only key/value -> airline index/airline name
	for _, key := range citySlc {
		value, err := client.Get(key).Result()
		if err != nil {
			// log error
			value = "missing value"
		}
		data[key] = value

	}

	return data, nil
}
