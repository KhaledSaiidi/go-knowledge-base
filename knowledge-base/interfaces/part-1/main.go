package main

import "math/rand"

func main() {
	team := make([]Player, 11)
	for i := 0; i < len(team)-2; i++ {
		team[i] = FootballPlayer{
			stamina: rand.Intn(10),
			power:   rand.Intn(10),
		}
	}
	team[len(team)-2] = CR7{
		stamina: 10,
		power:   10,
		SUI:     10,
	}
	team[len(team)-1] = Messi{
		stamina: 10,
		power:   10,
		SUI:     10,
	}
	for i := 0; i < len(team); i++ {
		team[i].KickBall()
	}
}
