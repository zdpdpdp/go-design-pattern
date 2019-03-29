package _3_abstract_factory

import "fmt"

type Car interface {
	Run()
}

type Suv interface {
	RunInMountain()
}

//AbstractCarFactory 抽象工厂接口
type AbstractCarFactory interface {
	CreateCar() Car
	CreateSuv() Suv
}

type ToyataFactory struct {
}

func (ToyataFactory) CreateCar() Car {
	return ToyataCar{}
}

func (ToyataFactory) CreateSuv() Suv {
	return ToyataSuv{}
}

type ToyataCar struct {
}

func (ToyataCar) Run() {
	fmt.Println("run with toyata car")
}

type ToyataSuv struct {
}

func (ToyataSuv) RunInMountain() {
	fmt.Println("run with to")
}

type BYDFactory struct {
}

func (BYDFactory) CreateCar() Car {
	return BYDCar{}
}

func (BYDFactory) CreateSuv() Suv {
	return BYDSuv{}
}

type BYDCar struct {
}

func (BYDCar) Run() {
	fmt.Println("i am byd car")
}

type BYDSuv struct {
}

func (BYDSuv) RunInMountain() {
	fmt.Println("i am byd suv")
}
