package main

import (
	mullitover "advent_of_code/2024/03_Mull_it_over"
)

func main() {
	// camelcard.Run()
	// r := camelcard.RunWithJoker("")
	// r := seeds.Compute()
	// r := seeds.ComputeRange()
	// r := lenslibrary.Compute(lenslibrary.INPUT)
	// r := lenslibrary.ComputePartTwo(lenslibrary.INPUT)
	// r := historianhysteria.Run(historianhysteria.LeftInput, historianhysteria.RightInput)
	// r := historianhysteria.Run(historianhysteria.LeftInput, historianhysteria.RightInput)
	// r := historianhysteria.RunSimilarityScore(historianhysteria.LeftInput, historianhysteria.RightInput)
	// r := rednosedreports.HowManyAreSafe(rednosedreports.INPUT)
	r := mullitover.RunWithOnOff(mullitover.INPUT)
	println(r)
}
