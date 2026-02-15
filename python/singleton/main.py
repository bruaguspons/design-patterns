from logging import Logger

if __name__ == "__main__":
    logger1 = Logger()
    logger2 = Logger()

    logger1.log("This is a log message.")

    # Verificamos que es la misma instancia
    print(logger1 is logger2)  # True
