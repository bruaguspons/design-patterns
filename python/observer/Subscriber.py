from abc import abstractmethod


class Subscriber():
    @abstractmethod
    def update(self, data: str) -> None:
        pass
