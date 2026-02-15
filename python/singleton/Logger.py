class Logger:
    _instance = None  # instancia única

    def __new__(cls):
        if cls._instance is None:
            cls._instance = super(Logger, cls).__new__(cls)
        return cls._instance

    def log(self, message: str) -> None:
        print(f"[LOG]: {message}")
