package main

import "fmt"

type Player interface {
	KickBall()
}

type CR7 struct {
	stamina int
	power   int
	SUI     int
}

func (f CR7) KickBall() {
	shot := f.stamina + f.power*f.SUI
	fmt.Println("CR7 Kick the ball", shot)
}

type Messi struct {
	stamina int
	power   int
	SUI     int
}

func (f Messi) KickBall() {
	shot := f.stamina + f.power*f.SUI
	fmt.Println("Messi Kick the ball", shot)
}

type FootballPlayer struct {
	stamina int
	power   int
}

func (f FootballPlayer) KickBall() {
	shot := f.stamina + f.power
	fmt.Println("Kick the ball", shot)
}
