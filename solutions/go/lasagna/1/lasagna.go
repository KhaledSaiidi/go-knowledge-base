package lasagna

const OvenTime = 40

func RemainingOvenTime(actualMinutesInOven int) int {
    remainingOvenTime := OvenTime - actualMinutesInOven
    return remainingOvenTime
}

func PreparationTime(numberOfLayers int) int {
    preparationTime := numberOfLayers * 2
    return preparationTime
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
    elapsedTime := PreparationTime(numberOfLayers) + actualMinutesInOven
    return elapsedTime
}
func main() {
    ElapsedTime(3, 20)
}
