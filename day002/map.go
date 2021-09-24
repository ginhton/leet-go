package main

import (
	"fmt"
)

func main() {
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Roma"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "xindeli"

	for country := range countryCapitalMap {
		fmt.Println(country, "'s capital is ", countryCapitalMap[country])
	}

	capital, ok := countryCapitalMap["American"]
	if ok {
		fmt.Println("American's capital is ", capital)
	} else {
		fmt.Println("Not found")
	}

	delete(countryCapitalMap, "France")
}
