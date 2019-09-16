package main

import (
	"fmt"
	"strings"
)

// Heres the syntax to define a method
// it uses the 'func' key word. 
// method is something you apply to a particulart custom datatype. 
func (nameOfHero Avenger) UpperCase() string {

	return strings.ToUpper(nameOfHero.SuperHeroName)
}

// have to define struct at same level scope as method.
type Avenger struct {
	ID            int
	SuperHeroName string
	RealName      string
}

func main() {

	firstAvenger := Avenger{
		ID:            1,
		SuperHeroName: "ironman",
		RealName:      "Tony Stark",
	}

	/* 
	Notice the variable.methodname() syntax for applying a method to a variable. 
	*/
	fmt.Println(firstAvenger.UpperCase())

	//superheroUpperCase := superHero.UpperCase()
	//fmt.Println(superheroUpperCase)
}

