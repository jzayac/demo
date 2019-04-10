package cache

import (
	"errors"
	"reflect"
	"testing"
)

type modelInterfaceTest struct {
	data map[string]string
	err  error
}

func (m modelInterfaceTest) GetList() (map[string]string, error) {
	return m.data, m.err
}

func TestCacheLen(t *testing.T) {
	c := cache{}

	size := c.Len()
	expect := 0

	if size != expect {
		t.Errorf("Expect Len to be %d but get %d", expect, size)
	}
}

func TestCacheGetList(t *testing.T) {
	cases := []struct {
		inst      *cache
		expectLen int
	}{
		{
			inst: &cache{
				data: cacheIata{
					"1": "1",
				},
				mi: nil,
			},
			expectLen: 1,
		},
		{
			inst: &cache{
				data: cacheIata{
					"1": "1",
					"2": "2",
				},
				mi: nil,
			},
			expectLen: 2,
		},
	}

	for _, c := range cases {
		list := c.inst.GetList()
		size := len(list)
		if size != c.expectLen {
			t.Errorf("Expect Len to be %d but get %d", c.expectLen, size)
		}
	}
}

func TestGetListEqual(t *testing.T) {
	cases := []struct {
		inst      *cache
		expectLen int
	}{
		{
			inst: &cache{
				data: cacheIata{
					"1": "1",
				},
				mi: nil,
			},
			expectLen: 1,
		},
		{
			inst: &cache{
				data: cacheIata{
					"1": "1",
					"2": "2",
				},
				mi: nil,
			},
			expectLen: 2,
		},
	}

	for _, c := range cases {
		list := c.inst.GetList()

		if !reflect.DeepEqual(cacheIata(list), c.inst.data) {
			t.Errorf("Expected to be equal %+v :: %+v", list, c.inst.data)
		}
	}
}

func TestGetName(t *testing.T) {
	cases := []string{
		"test", "city", "1234",
	}

	for _, value := range cases {
		c := &cache{
			name: value,
		}

		name := c.GetName()
		if name != value {
			t.Error("Expected name to be ", value)
		}
	}
}

func TestnewCache(t *testing.T) {
	cases := []modelInterfaceTest{
		modelInterfaceTest{
			data: map[string]string{
				"1": "1",
				"2": "2",
			},
			err: nil,
		},
		modelInterfaceTest{
			data: map[string]string{},
			err:  errors.New("test error"),
		},
	}

	for _, c := range cases {
		cache, err := NewCache("test", 3, c)
		if !reflect.DeepEqual(cache.data, cacheIata(c.data)) {
			t.Errorf("Expected data to be equal %+v :: %+v", cache.data, c.data)
		}

		if !reflect.DeepEqual(err, c.err) {
			t.Errorf("Expected error to be equal %+v :: %+v", err, c.err)
		}

	}
}
