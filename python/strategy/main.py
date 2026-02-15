from Commuter import Commuter, CarStrategy, BikeStrategy

if __name__ == "__main__":
    commuter = Commuter()

    commuter.set_strategy(CarStrategy())
    commuter.commute()

    commuter.set_strategy(BikeStrategy())
    commuter.commute()
