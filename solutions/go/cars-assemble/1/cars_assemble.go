package cars
import "fmt"

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	tensCar := int(carsCount/10)
	restCars := carsCount % 10
    carCost := uint(tensCar * 95000 + restCars * 10000)

    fmt.Printf("tensCar: %d\n", tensCar)
	fmt.Printf("restCars: %d\n", restCars)
	fmt.Printf("carCost: %d\n", carCost)
    
    return carCost
}
