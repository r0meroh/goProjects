package main

type ScoreType int

const (
	Food ScoreType = iota
	Beverage
	Water
	Cheese
)

type NutritionalScore struct {
	Value     int
	Positive  int
	Negative  int
	ScoreType ScoreType
}

type EnergyKJ float64
type SugarGram float64
type SaturatedFattyAcids float64
type SodiumMilligram float64
type FruitsPercent float64
type Fibergram float64
type Proteingram float64

type NutrionalData struct {
	Energy              EnergyKJ
	Sugars              SugarGram
	SaturatedFattyAcids SaturatedFattyAcids
	Sodium              SodiumMilligram
	Fruits              FruitsPercent
	Fiber               Fibergram
	Protein             Proteingram
	IsWater             bool
}

func (e EnergyKJ)GetPoints(st ScoreType) int {

	energy := 0
	return energy
}

func (s SugarGram)GetPoints(st ScoreType) int {

	sg := 0
	return sg
}

func (sf SaturatedFattyAcids)GetPoints(st ScoreType) int {

	sfa := 0
	return sfa
}

func (sd SodiumMilligram)GetPoints(st ScoreType) int {

	sm := 0
	return sm
}

func (f FruitsPercent)GetPoints(st ScoreType) int {

	fp := 0
	return fp
}

func (f Fibergram)GetPoints(st ScoreType) int {

	fg := 0
	return fg
}

func (p Proteingram)GetPoints(st ScoreType) int {

	pg := 0
	return pg
}

func EnergyFromkcal(kcal float64) EnergyKJ{
	return EnergyKJ(kcal * 4.184)
}

func SodiumFromSalt(saltMg float64) SodiumMilligram{
	return SodiumMilligram(saltMg/2.5)
}
func GetNutritionalScore(n NutrionalData, st ScoreType) NutritionalScore {
	value := 0
	positive := 0
	negative := 0
	
	if st != Water {
		fruitPoints := n.Fruits.GetPoints(st)
		fiberPoints := n.Fiber.GetPoints(st)
		proteinPoints := n.Protein.GetPoints(st)

		positive = fruitPoints + fiberPoints + proteinPoints
		negative = n.Energy.GetPoints(st) + n.Sugars.GetPoints(st) + n.SaturatedFattyAcids.GetPoints(st) + n.Sodium.GetPoints(st)
	}
	
	return NutritionalScore{
		Value: value,
		Positive: positive,
		Negative: negative,
		ScoreType: st,
	}
}
