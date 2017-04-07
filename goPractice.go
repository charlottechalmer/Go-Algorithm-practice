package main

import "fmt"

func main() {
	aCar := car{}
	fmt.Println(aCar)
	aCar.mpg = 25
	aCar.numberOfDoors = 2
	fmt.Println(aCar)
	aCar.getMpg()
	bCar := car{vehicle{19, 4}, red}
	fmt.Println(bCar)
	bCar.getMpg()
	cCar := car{}
	cCar.color = black
	cCar.mpg = 34
	fmt.Println(cCar)
	cCar.getMpg()

	defaultCar := newCar()
	fmt.Println("default car: ", defaultCar)

	//polymorphic
	aTruck := truck{vehicle{15, 2}, black}
	mans := [...]manufacturer{cCar, aTruck}
	for i := range mans {
		fmt.Println("iteration=", i, mans[i].getVIN())
	}
}

type vehicle struct {
	mpg           int
	numberOfDoors int
}

type Color string

const (
	blue  Color = "blue"
	red   Color = "red"
	black Color = "black"
)

type car struct {
	vehicle
	color Color
}

type truck struct {
	vehicle
	color Color
}

func (v *vehicle) getMpg() {
	println("This vehicle gets", v.mpg, "mpg")
}

func (t truck) getVIN() string {
	return "truck VIN"
}

func (c car) getVIN() string {
	return "car VIN"
}

type manufacturer interface {
	getVIN() string
}

func newCar() *car {
	result := car{}
	result.mpg = 20
	result.numberOfDoors = 4
	result.color = black
	return &result
}
