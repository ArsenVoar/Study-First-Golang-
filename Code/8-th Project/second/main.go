package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Build struct {
	Building
}

type Building struct {
	Builder
	Name string
}

type Builder interface {
	Build()
}

type WoodBuilder struct {
	Person
}

type BrickBuilder struct {
	Person
}

type woodenBuilding struct {
	Building
}

type BrickBuilding struct {
	Building
}

func (wb WoodBuilder) Build() {
	fmt.Println("Wood Built")
}

func (bb BrickBuilder) Build() {
	fmt.Println("Brick Built")
}

func main() {
	//explanation()
	usecase()
}

func usecase() {
	woodenBuilding := Building{
		Builder: WoodBuilder{Person{
			Name: "Garry",
			Age:  38,
		}},
		Name: "woodenBuilding",
	}
	woodenBuilding.Build()

	brickBuilding := Building{
		Builder: BrickBuilder{Person{
			Name: "Luck",
			Age:  29,
		},
		},
		Name: "Brick Build",
	}
	brickBuilding.Build()
}

/*type Builder interface {
	Build()
}

type Person struct {
	Name string
	Age  int
}

type WorkExperience struct {
	Name string
	Age  int
}

func (p Person) printName() {
	fmt.Println(p.Name)
}

type WoodBuilder struct {
	Person        // - if WoodBuilder doesn't have Name
	Name   string // if have
}

func (w WoodBuilder) printName() {
	fmt.Println(w.Name)
} // its the --- 2 --- when WoodBuilder have printName. And name will be Pit to do name Jack we need to go to 3
*/

/*func explanation() {
	builder := WoodBuilder{Person{Name: "Jack", Age: 40}, "Pit"} // if have
	fmt.Printf("Type: %T Value: %#v\n", builder, builder)

	//shorthands
	fmt.Println(builder.Person.Age)
	fmt.Println(builder.Age)

	//shadowing
	//fmt.Println(builder.Name)  // -- 1  // in one will be Jack, but if we add print.Name to WoodBuilder will be 2
	fmt.Println(builder.Name)
	fmt.Println(builder.Person.Name) // its --- 3 ---

	builder.printName()
	builder.Person.printName() // and its --- 3 --- too

	/*builder := WoodBuilder{Person{Name: "Jack", Age: 40}}
	fmt.Printf("Type: %T Value: %#v\n", builder, builder)

	//shorthands
	fmt.Println(builder.Person.Age)
	fmt.Println(builder.Age)

	//shadowing
	fmt.Println(builder.Name)

	builder.printName() // - if WoodBuilder doesn't have Name
} */
