package main

import "fmt"

type Computer interface {
	Start()
	RunGame()
	StopGame()
	RunComputation()
	StopComputation()
	Shutdown()
	GetName() string
}

type UniversalComputer struct {
	Name string
	GPU  string
	CPU  string
}

type Laptop struct {
	Name string
	GPU  string
	CPU  string
}

func (uc *UniversalComputer) Start() {
	fmt.Printf("%s стартует\n", uc.Name)
}

func (uc *UniversalComputer) Shutdown() {
	fmt.Printf("%s выключается\n", uc.Name)
}

func (uc *UniversalComputer) GetName() string {
	return uc.Name
}

func (uc *UniversalComputer) RunGame() {
	fmt.Printf("%s запускает игру на %s\n", uc.Name, uc.GPU)
}

func (uc *UniversalComputer) StopGame() {
	fmt.Printf("%s останавливает игру\n", uc.Name)
}

func (uc *UniversalComputer) RunComputation() {
	fmt.Printf("%s начинает вычислительную задачу на %s\n", uc.Name, uc.CPU)
}

func (uc *UniversalComputer) StopComputation() {
	fmt.Printf("%s завершил вычислительную задачу\n", uc.Name)
}

func (l *Laptop) Start() {
	fmt.Printf("%s стартует\n", l.Name)
}

func (l *Laptop) Shutdown() {
	fmt.Printf("%s выключается\n", l.Name)
}

func (l *Laptop) GetName() string {
	return l.Name
}

func (l *Laptop) RunGame() {
}

func (l *Laptop) StopGame() {
}

func (l *Laptop) RunComputation() {
}

func (l *Laptop) StopComputation() {
}

func main() {
	uniComp := &UniversalComputer{
		Name: "Computer",
		GPU:  "NVIDIA RTX 3060",
		CPU:  "Intel Core i5-11400",
	}
	laptp := &Laptop{
		Name: "Laptop",
		GPU:  "RTX3050",
		CPU:  "I5-12500H",
	}

	computers := []Computer{uniComp, laptp}
	for _, comp := range computers {
		comp.Start()
		comp.RunGame()
		comp.StopGame()
		comp.RunComputation()
		comp.StopComputation()
		comp.Shutdown()
		fmt.Println()
	}
}
