package main

func main() {
	commuter := &Commuter{}

	commuter.SetStrategy(CarStrategy{})
	commuter.Commute()

	commuter.SetStrategy(BikeStrategy{})
	commuter.Commute()
}
