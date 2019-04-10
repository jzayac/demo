package model

type airlineModel struct{}

func NewAirline() airlineModel {
	ins := airlineModel{}
	return ins
}

func (c airlineModel) GetList() (map[string]string, error) {
	var err error

	db, errDb := dbConnect()

	if errDb != nil {
		return nil, err
	}
	defer db.Close()

	al := []*Iata{}

	err = db.Find(&al).Error

	data := make(map[string]string)
	if err != nil {
		return data, err
	}

	for _, val := range al {
		data[val.Index] = val.Value
	}

	return data, err
}
