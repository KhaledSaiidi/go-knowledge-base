package lasagnamaster
import "fmt"
// TODO: define the 'PreparationTime()' function

func PreparationTime(layers []string, averageTime int) int {
    if averageTime <= 0 {
        averageTime = 2
    }
	return len(layers) * averageTime
    
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64){
	noodles := 0
	sauce := 0
    for i := range len(layers) {
        if layers[i] == "noodles" {
            noodles++
        } else if layers[i] == "sauce" {
            sauce++
        }
    }
	neededNoodles := noodles * 50
    neededSauce := float64(sauce) * 0.2
    fmt.Printf("For %d noodles layers and %d sauce layers, you will need %d gram of noodles and %2f liters of sauce",noodles,sauce, neededNoodles, neededSauce)
    return neededNoodles, neededSauce
}




// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(frindIngredients, ownIngredients []string){
	secretIngredient := frindIngredients[len(frindIngredients) - 1]
    ownIngredients[len(ownIngredients) -1] = secretIngredient
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(neededAmount []float64, numberPortions int) []float64 {
	scaled := make([]float64, len(neededAmount))
	for i, v := range neededAmount {
		scaled[i] = v / 2 * float64(numberPortions)
	}
	return scaled
}    
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
