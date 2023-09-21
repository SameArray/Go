package main

import "fmt"

type Strategy interface {
	Process() error
}

type FirstAlgorithm struct{}

func (f *FirstAlgorithm) Process() error {
	fmt.Println("first algorithm is in progress")
	// do something

	return nil
}

type SecondAlgorithm struct{}

func (s *SecondAlgorithm) Process() error {
	fmt.Println("second algorithm is in progress")
	// do something
	// do something else

	return nil
}

type ThirdAlgorithm struct{}

func (s *ThirdAlgorithm) Process() error {
	fmt.Println("third algorithm is in progress")
	// do something

	return nil
}

type Context struct {
	activeStrategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.activeStrategy = strategy
}

func (c *Context) Process() error {
	err := c.activeStrategy.Process()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	context := Context{}
	context.SetStrategy(&FirstAlgorithm{})

	err := context.Process()
	if err != nil {
		return
	}

	context.SetStrategy(&SecondAlgorithm{})
	err = context.Process()
	if err != nil {
		return
	}

	context.SetStrategy(&ThirdAlgorithm{})
	err = context.Process()
	if err != nil {
		return
	}

	//f := FirstAlgorithm{}
	//err := f.Process()
	//if err != nil {
	//	return
	//}
}
