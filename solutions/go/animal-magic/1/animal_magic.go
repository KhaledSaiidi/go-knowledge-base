package chance
import "math/rand"
// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
return rand.Intn(19) + 1
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	randomInt := float64(rand.Intn(12))
    randomFloat := rand.Float64()
    return randomInt + randomFloat
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
    rand.Shuffle(len(animals), func(i, j int) {
		animals[i], animals[j] = animals[j], animals[i]
	})
    return animals
}
