package main

func mapExample(){
	myMap := make(map[string]int)
	myMap["apple"] = 5
	myMap["banana"] = 3

	for key, value := range myMap {
		println(key, value)
	}
}