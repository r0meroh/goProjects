package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to nutriscore")

	ns := GetNutritionalScore(NutrionalData{
		Energy:              EnergyFromkcal(0),
		Sugars:              SugarGram(10),
		SaturatedFattyAcids: SaturatedFattyAcids(33),
		Sodium:              SodiumMilligram(100),
		Fruits:              FruitsPercent(22),
		Fiber:               Fibergram(11),
		Protein:             Proteingram(30),
	}, Food)

	fmt.Printf("your score is: %v\n", ns.Value)
	fmt.Printf("Nutriscore: %s\n", ns.GetNutriScore())
}
