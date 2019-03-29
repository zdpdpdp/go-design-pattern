package _3_abstract_factory

import "testing"

func TestFactory(t *testing.T) {

	var factory AbstractCarFactory
	var car Car
	var suv Suv

	//可以看到,业务逻辑变更时, 只需要更换一下具体的工厂即可

	factory = ToyataFactory{}
	car = factory.CreateCar()
	car.Run()
	suv = factory.CreateSuv()
	suv.RunInMountain()

	factory = BYDFactory{}
	car = factory.CreateCar()
	car.Run()
	suv = factory.CreateSuv()
	suv.RunInMountain()

}
