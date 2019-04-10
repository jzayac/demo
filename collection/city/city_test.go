package city

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

func TestCityLen(t *testing.T) {
	c := city{}

	size := c.Len()
	expect := 0

	if size != expect {
		t.Errorf("Expect Len to be %d but get %d", expect, size)
	}
}

func TestCityGetList(t *testing.T) {
	cases := []struct {
		inst      *city
		expectLen int
	}{
		{
			inst: &city{
				data: cityIata{
					"1": "1",
				},
				mi: nil,
			},
			expectLen: 1,
		},
		{
			inst: &city{
				data: cityIata{
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
		inst      *city
		expectLen int
	}{
		{
			inst: &city{
				data: cityIata{
					"1": "1",
				},
				mi: nil,
			},
			expectLen: 1,
		},
		{
			inst: &city{
				data: cityIata{
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

		if !reflect.DeepEqual(cityIata(list), c.inst.data) {
			t.Errorf("Expected to be equal %+v :: %+v", list, c.inst.data)
		}
	}
}

func TestGetName(t *testing.T) {
	c := &city{}

	name := c.GetName()

	if name != "city" {
		t.Error("Expected name to be city")
	}
}

func TestnewCity(t *testing.T) {
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
		city, err := newCity(c)
		if !reflect.DeepEqual(city.data, cityIata(c.data)) {
			t.Errorf("Expected data to be equal %+v :: %+v", city.data, c.data)
		}

		if !reflect.DeepEqual(err, c.err) {
			t.Errorf("Expected error to be equal %+v :: %+v", err, c.err)
		}

	}
}
