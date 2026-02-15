package main

import "fmt"

type CarStrategy struct{}

func (c CarStrategy) Transport() {
	fmt.Println("Transporting by car")
}

type BikeStrategy struct{}

func (b BikeStrategy) Transport() {
	fmt.Println("Transporting by bike")
}

type BusStrategy struct{}

func (b BusStrategy) Transport() {
	fmt.Println("Transporting by bus")
}

type Commuter struct {
	strategy TransportStrategy
}

func (c *Commuter) SetStrategy(strategy TransportStrategy) {
	c.strategy = strategy
}

func (c *Commuter) Commute() {
	if c.strategy == nil {
		fmt.Println("No transport strategy set")
		return
	}
	c.strategy.Transport()
}
