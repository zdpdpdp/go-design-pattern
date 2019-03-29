package _4_builder

import "testing"

func TestBuilder(t *testing.T) {

	smallCarDirector := NewDirector(SmallCarBuilder{})
	smallCar := smallCarDirector.construct()
	smallCar.Run()

	suvCarDirector := NewDirector(SuvCarBuilder{})
	suvCar := suvCarDirector.construct()
	suvCar.Run()
}
