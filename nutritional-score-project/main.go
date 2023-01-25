package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to nutriscore")

	ns := GetNutritionalScore(NutrionalData{
		Energy:              EnergyFromkcal(0),
		Sugars:              SugarGram(10),
		SaturatedFattyAcids: SaturatedFattyAcids(),
		Sodium:              SodiumMilligram(),
		Fruits:              FruitsPercent(),
		Fiber:               Fibergram(),
		Protein:             Proteingram(),
	}, Food)

	fmt.Printf("your score is %v",)
}
