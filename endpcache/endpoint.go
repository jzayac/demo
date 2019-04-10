package endpcache

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"demo/cache"
)

type getListResponse struct {
	Data *cache.Iata `json:"data"`
	Err  error       `json:"err,omitempty"` // errors don't define JSON marshaling
}

func (r getListResponse) error() error { return r.Err }

func makeGetListEndpoint(svc Service, name string) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		v, err := svc.GetList(name)
		if err != nil {
			return getListResponse{Data: v, Err: err}, nil
		}
		return getListResponse{v, nil}, nil
	}
}
