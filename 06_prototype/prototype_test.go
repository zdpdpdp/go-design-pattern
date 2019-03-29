package _6_prototype

import "testing"

type SmallCar struct {
	speed int
}

type SuvCar struct {
	height int
}

func (s *SmallCar) Clone() Cloneable {
	return &SmallCar{speed: s.speed}
}

func (s *SuvCar) Clone() Cloneable {
	return &SuvCar{height: s.height}
}

var manager Manager

func init() {
	manager = NewManager()
	var sc SmallCar
	sc.speed = 110
	manager.Set("small_car", &sc)

	var su SuvCar
	su.height = 250
	manager.Set("suv_car", &su)
}

func TestManager(t *testing.T) {
	cloneSmallCar := manager.Get("small_car").Clone()

	newSmallCar := cloneSmallCar.(*SmallCar)
	if newSmallCar.speed != 110 {
		t.Fatal("wrong new small car")
	}

	cloneSuvCar := manager.Get("suv_car").Clone()

	newSuvCar := cloneSuvCar.(*SuvCar)
	if newSuvCar.height != 250 {
		t.Fatal("wrong new small car")
	}

}
