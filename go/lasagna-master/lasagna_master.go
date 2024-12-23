package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, preparationTime int) int {
	if preparationTime == 0 {
		preparationTime = 2
	}
	return preparationTime * len(layers)
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0
	for _, ingredient := range layers {
		switch ingredient {
		case "sauce":
			sauce += 0.2
		case "noodles":
			noodles += 50
		}
	}
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(potionsAmount []float64, potions int) []float64 {
	if potions == 0 || len(potionsAmount) == 0 {
		return nil
	}

	even := potions%2 == 0

	recalculatedAmounts := make([]float64, len(potionsAmount))

	if !even {
		for i, amount := range potionsAmount {
			recalculatedAmounts[i] = amount / 2
		}
	}

	for i, amount := range potionsAmount {
		recalculatedAmounts[i] = amount * (float64(potions) / 2)
	}

	return recalculatedAmounts
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
