package main

import "testing"

func TestComposition(t *testing.T) {
	swimmer := CompositeSwimmerA{
		MySwim: Swim,
	}
	swimmer.MyAthlete.Train()
	swimmer.MySwim()
	localSwim := Swim
	swimmer2 := CompositeSwimmerA{
		MySwim: localSwim,
	}
	swimmer2.MyAthlete.Train()
	swimmer2.MySwim()
}
func TestCompositionAnimal(t *testing.T) {
	fish := Shark{Swim: Swim}
	fish.Swim()
	fish.Eat()
}
func TestCompositeB(t *testing.T) {
	swimmer := CompositeSwimmerB{
		&Athlete{},
		&SwimmerImpl{},
	}
	swimmer.Swim()
	swimmer.Train()
}
