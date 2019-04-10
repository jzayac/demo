package main

import (
	"demo/collection/city"
	// "demo/model"
	"fmt"
	"math/rand"
	"time"
	// "demo/lib/database"
)

type kokot map[string]string

// func main() {
// 	fmt.Println("test")

// 	if err := model.InitDatabaseModel(); err != nil {
// 		fmt.Println(err)
// 	}

// 	arilineModel := model.NewAirline()
// 	if list, err := arilineModel.GetList(); err == nil {

// 		for key, value := range list {
// 			fmt.Println(key, value)
// 		}
// 	}

// 	cityModel := model.NewCity()
// 	// list, _ := cityModel.GetList()
// 	// for _, data := range list {
// 	// 	fmt.Println(data.Index, data.Value)
// 	// }

// 	list, _ := cityModel.GetList()
// 	for key, value := range list {
// 		fmt.Println(key, value)
// 	}

// }

// func main() {

// 	_ = model.NewCity()
// 	// test := &kokot{}

// 	// (*test)["aaa"] = "bbbb"

// 	// fmt.Println(test)
// 	// test2 := map[string]string(*test)
// 	// fmt.Println(test2)

// 	// test := &kokot{}

// 	test2 := make(map[string]string)
// 	test2["aaa"] = "bbbb"

// 	fmt.Println(test2)
// 	test := kokot(test2)
// 	fmt.Println(test)
// }

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type kokotina struct{}

func (k kokotina) GetList() (map[string]string, error) {

	t := make(map[string]string)
	for i := 0; i < 10; i++ {

		key := RandStringRunes(4)
		value := RandStringRunes(4)
		t[key] = value

	}
	return t, nil
}

func main() {

	// TEST PERIDICAL UPDATE

	// model.GetInstance()
	ci := city.GetInstance()
	fmt.Println("_____________________")
	ls, _ := kokotina{}.GetList()
	fmt.Println(ls)

	c := make(chan bool)
	cm := kokotina{}

	fmt.Println(&ci)
	go func() {
		ci.TimeoutUpdate(cm)

	}()

	for t := range time.NewTicker(2 * time.Second).C {

		fmt.Println("---------------------------- tick ", t)
		fmt.Println(ci.GetList())
	}
	<-c
	fmt.Println("chuj")
}
