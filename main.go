package main

import (
	"fmt"
	"strings"
)

/*
	Heres the syntax to define a method uses the 'func' key word, the
	same keyword we use when creating a function. The syntax is a bit different this time thoough.
	This time we have to defined the input parameter "(nameOfHero Avenger)" first, followed by the method name
    method is something you apply to a particular custom datatype.
*/
func (MyHero Avenger) UpperCase() string {

	// Notice how we drill down the custom object's (struct) field using the dot notaton:
	return strings.ToUpper(MyHero.SuperHeroName)
}

/*
Another function, which this time requires an input parameter, 'metricweight'
*/
func (MyHero Avenger) GetWeight(metricweight bool) float64 {

	if metricweight {
		return float64(MyHero.WeightInKG)
	} else {
		WeightInPounds := float64(MyHero.WeightInKG) * 2.205
		return WeightInPounds

	}
}

// have to define struct at same level scope as method.
type Avenger struct {
	ID            int
	SuperHeroName string
	RealName      string
	WeightInKG    int
}

func main() {

	firstAvenger := Avenger{
		ID:            1,
		SuperHeroName: "ironman",
		RealName:      "Tony Stark",
		WeightInKG:    75,
	}

	/*
		Notice the variable.methodname() syntax for applying a method to a variable.
		we use a dot notation to apply the UpperCase method to the firstAvenger custom object.
	*/
	fmt.Println(firstAvenger.UpperCase())

	fmt.Println(firstAvenger.GetWeight(true))
	fmt.Println(firstAvenger.GetWeight(false))
}
