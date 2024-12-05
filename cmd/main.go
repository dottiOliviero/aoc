package main

import (
	historianhysteria "advent_of_code/2024/01_Historian_Hysteria"
)

func main() {
	// camelcard.Run()
	// r := camelcard.RunWithJoker("")
	// r := seeds.Compute()
	// r := seeds.ComputeRange()
	// r := lenslibrary.Compute(lenslibrary.INPUT)
	// r := lenslibrary.ComputePartTwo(lenslibrary.INPUT)
	// r := historianhysteria.Run(historianhysteria.LeftInput, historianhysteria.RightInput)
	r := historianhysteria.RunSimilarityScore(historianhysteria.LeftInput, historianhysteria.RightInput)
	println(r)
}
