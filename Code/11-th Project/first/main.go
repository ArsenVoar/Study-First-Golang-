package main

//maps

import "fmt"

type Person struct {
	Id   int64
	Name string
}

func main() {
	var defaultMap map[int64]string

	fmt.Printf("Type: %T Value:%#v\n", defaultMap, defaultMap)
	fmt.Printf("Len: %d\n", len(defaultMap))

	mapByMake := make(map[string]string, 3)
	fmt.Printf("Type: %T Value:%#v\n", mapByMake, mapByMake)

	mapByLiteral := map[string]int{"Mortis": 29, "Ben": 20}

	fmt.Printf("Type: %T Value:%#v\n", mapByLiteral, mapByLiteral)
	fmt.Printf("Len: %d\n", len(mapByLiteral))

	mapByMake["First"] = "Vito"
	fmt.Printf("Type: %T Value:%#v\n", mapByMake, mapByMake)

	//get map default value
	fmt.Println(mapByLiteral["Second"])

	one, two := mapByLiteral["Second"]
	fmt.Printf("Value: %d isExist: %t\n", one, two)

	//unique values
	users := []Person{
		{
			Id:   1,
			Name: "SomebodyOne",
		},
		{
			Id:   8,
			Name: "Somebodytwo",
		},
		{
			Id:   8,
			Name: "Somebodythree",
		},
	}
	uniquePersons := make(map[int64]struct{}, len(users))

	for _, user := range users {
		if _, ok := uniquePersons[user.Id]; !ok {
			uniquePersons[user.Id] = struct{}{}
		}
	}

	fmt.Printf("Type: %T Value: %#v\n", uniquePersons, uniquePersons)
}
