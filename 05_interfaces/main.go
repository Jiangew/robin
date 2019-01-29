package main

import (
	"fmt"
)

// Defining Interfaces
type Guitarist interface {
	// PlayGuitar prints out "Playing Guitar"
	PlayGuitar()
}

type BaseGuitarist struct {
	Name string
}

type AcousticGuitarist struct {
	Name string
}

func (b BaseGuitarist) PlayGuitar() {
	fmt.Printf("%s plays the Base Guitar\n", b.Name)
}

func (b AcousticGuitarist) PlayGuitar() {
	fmt.Printf("%s plays the Acoustic Guitar\n", b.Name)
}

func main() {
	fmt.Println("------ Interfaces ------")

	var player BaseGuitarist
	player.Name = "Jiangew"
	player.PlayGuitar()

	var player2 AcousticGuitarist
	player2.Name = "JamesiWorks"
	player2.PlayGuitar()

	var guitarists []Guitarist
	guitarists = append(guitarists, player)
	guitarists = append(guitarists, player2)
	fmt.Println(guitarists)

	fmt.Println("------ Interfaces End ------")
}
