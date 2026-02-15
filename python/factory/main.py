from UserFactory import UserFactory

if __name__ == "__main__":
    admin = UserFactory.create_user("admin", {"name": "John", "age": 30})
    guest = UserFactory.create_user("guest", {"name": "Jane", "age": 25})
    regular = UserFactory.create_user("regular", {"name": "Bob", "age": 40})

    print(admin)
    print(guest)
    print(regular)
