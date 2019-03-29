package _4_builder

import "fmt"

type Car interface {
	Run()
}

type SmallCar struct {
}

func (SmallCar) Run() {
	fmt.Println("small car running")
}

type SuvCar struct {
}

func (SuvCar) Run() {
	fmt.Println("suv car running")
}

type CarBuilder interface {
	BuildTire()
	BuildEngine()
	BuildBody()
	GetResult() Car
}
type Director struct {
	builder CarBuilder
}

func NewDirector(builder CarBuilder) *Director {
	return &Director{builder}
}

func (d *Director) construct() Car {
	//创建引擎
	d.builder.BuildEngine()
	//创建轮胎
	d.builder.BuildTire()
	//创建车身
	d.builder.BuildBody()

	return d.builder.GetResult()
}

type SmallCarBuilder struct {
}

func (SmallCarBuilder) BuildTire() {
	fmt.Println("i build tire")
}

func (SmallCarBuilder) BuildEngine() {
	fmt.Println("i build engine")
}

func (SmallCarBuilder) BuildBody() {
	fmt.Println("i build body")
}

func (SmallCarBuilder) GetResult() Car {
	return &SmallCar{}
}

type SuvCarBuilder struct {
}

func (SuvCarBuilder) BuildTire() {
	fmt.Println("i build suv tire")
}

func (SuvCarBuilder) BuildEngine() {
	fmt.Println("i build suv engine")
}

func (SuvCarBuilder) BuildBody() {
	fmt.Println("i build suv body")
}

func (SuvCarBuilder) GetResult() Car {
	return &SuvCar{}
}
