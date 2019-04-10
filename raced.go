package main

import (
	"demo/collection"
	"fmt"
	"time"
)

func main() {

	col := collection.GetInstance()

	for t := range time.NewTicker(1 * time.Second).C {

		fmt.Println("---------------------------- tick ", t)
		city, _ := col.GetInstanceByKey("city")
		fmt.Println(city.GetList())

		airline, _ := col.GetInstanceByKey("airline")
		fmt.Println("---------------------------- tick ", t)
		fmt.Println(airline.GetList())
	}
}
