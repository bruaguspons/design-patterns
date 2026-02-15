from abc import abstractmethod


class TransportStrategy():

    @abstractmethod
    def transport(self) -> None:
        pass
