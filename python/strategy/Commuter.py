from TransportStrategy import TransportStrategy

class CarStrategy(TransportStrategy):
    def transport(self) -> None:
        print("Transporting by car")


class BikeStrategy(TransportStrategy):
    def transport(self) -> None:
        print("Transporting by bike")


class BusStrategy(TransportStrategy):
    def transport(self) -> None:
        print("Transporting by bus")
        
from typing import Optional


class Commuter:
    def __init__(self):
        self._strategy: Optional[TransportStrategy] = None

    def set_strategy(self, strategy: TransportStrategy) -> None:
        self._strategy = strategy

    def commute(self) -> None:
        if not self._strategy:
            print("No transport strategy set")
            return
        self._strategy.transport()

