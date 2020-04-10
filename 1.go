// go run 1.go
package main

import (
	"fmt"
	"strings"
)

// methods are just an alternative way for writing a function, that
// requires a struct as an input parameter. Let's take a
// look at the normal way first. In 2.go we'll look at the 'methods' way.

type Avenger struct {
	ID            int
	SuperHeroName string
	RealName      string
	WeightInKG    int
}

func UpperCase(MyHero Avenger) string {
	return strings.ToUpper(MyHero.SuperHeroName)
}

// Another function, which this time requires an input
// parameter, 'metricweight'
func GetWeight(MyHero Avenger, metricweight bool) float64 {
	if metricweight {
		return float64(MyHero.WeightInKG)
	} else {
		WeightInPounds := float64(MyHero.WeightInKG) * 2.205
		return WeightInPounds
	}
}

func main() {

	firstAvenger := Avenger{
		ID:            1,
		SuperHeroName: "ironman",
		RealName:      "Tony Stark",
		WeightInKG:    75,
	}

	fmt.Println(UpperCase(firstAvenger)) // IRONMAN

	fmt.Println(GetWeight(firstAvenger, true))  // 75
	fmt.Println(GetWeight(firstAvenger, false)) // 165.375
}
