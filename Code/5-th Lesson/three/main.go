package main

import (
	"fmt"
	"time"
)

type hogString string
type hogInt int

type Person struct {
	Name string
	Age  int
}

func main() {
	unnamed := struct {
		Name, LastName, YearOld string
	}{
		Name:     "NoName",
		LastName: "NoLastName",
		//YearOld: fmt.Sprintf("%s", time.Now())
		YearOld: time.Now().String(),
	}
	fmt.Println(unnamed)

	//create pointer to struct
	pVoAr := &Person{"VoAr", 10}
	fmt.Println(pVoAr)

	//create without field names
	ArsenVoar := Person{"ArsenVoar", 25}
	fmt.Println(ArsenVoar)
	//field accesing through the pointer
	pArsenVoar := &ArsenVoar
	fmt.Println((*pArsenVoar).Age)
	fmt.Println((pArsenVoar).Age)

	//create witn named field
	Voar := Person{
		Name: "Voar",
		Age:  20,
	}
	fmt.Println(Voar)

	//fields accesing
	var Arsen Person
	Arsen.Name = "Arsen"
	Arsen.Age = 60
	fmt.Println(Arsen)

	//create var struct...1

	fmt.Printf("%T %#v \n", Arsen, Arsen)

	//create var struct...2
	Arsen = Person{}
	fmt.Printf("%T %#v \n", Arsen, Arsen)

	// ...2

	lapInt := hogInt(5)
	fmt.Printf("%T %#v \n", lapInt, lapInt)
	fmt.Println(int(lapInt))

	var lapString hogString
	fmt.Printf("%T %#v \n", lapString, lapString)
	lapString = "Hi everybody"
	fmt.Printf("%T %#v \n", lapString, lapString) // - custom type...1
}
