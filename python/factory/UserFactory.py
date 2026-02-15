from dataclasses import dataclass
from typing import Literal

class User():
    name: str
    age: int


@dataclass
class AdminUser(User):
    name: str
    age: int


@dataclass
class GuestUser(User):
    name: str
    age: int


@dataclass
class RegularUser(User):
    name: str
    age: int


UserType = Literal["admin", "guest", "regular"]

class UserFactory:

    @staticmethod
    def create_user(user_type: UserType, payload: dict) -> User:
        name = payload["name"]
        age = payload["age"]

        if user_type == "admin":
            return AdminUser(name, age)
        elif user_type == "guest":
            return GuestUser(name, age)
        elif user_type == "regular":
            return RegularUser(name, age)
        else:
            raise ValueError("Invalid user type")

