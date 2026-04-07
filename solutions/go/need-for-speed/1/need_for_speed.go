package speed

// TODO: define the 'Car' type struct
type Car struct {
    speed int
    batteryDrain int
    battery int
    distance int
}
// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	c := Car {
        speed: speed,
        batteryDrain: batteryDrain,
        battery: 100,
        distance: 0,
    }
    return c
}

// TODO: define the 'Track' type struct
type Track struct {
    distance int
}
// NewTrack creates a new track
func NewTrack(distance int) Track {
	t := Track {
        distance: distance,
    }
    return t
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if car.batteryDrain > car.battery {
        return car
    }
    newCar := Car {
        speed: car.speed,
        batteryDrain: car.batteryDrain,
        battery: car.battery - car.batteryDrain,
        distance: car.distance + car.speed,
    }
    return newCar
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	carCanMove :=  (car.battery /  car.batteryDrain) * car.speed
	return car.distance + carCanMove >= track.distance
}
