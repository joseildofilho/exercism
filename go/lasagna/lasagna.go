package lasagna

// TODO: define the 'OvenTime' constant

const OvenTime = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	remaining := OvenTime - actualMinutesInOven
	if(remaining < 0) {
		panic("Its burned")
	}
	return remaining
}

const LayerTime = 2

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	if(numberOfLayers < 0) {
		panic("Number Of Layers cannot be negative")
	}
	return numberOfLayers * LayerTime
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}