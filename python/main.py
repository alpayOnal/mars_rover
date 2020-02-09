from plateau import Plateau
from rover import Rover


def main():
    plateau = Plateau(5, 5)

    rover = Rover(1, 2, "N", plateau)
    rover.run("LMLMLMLMM")
    print(rover.position)
    rover = Rover(3, 3, "E", plateau)

    rover.run("MMRMMRMRRM")
    print(rover.position)


if __name__ == "__main__":
    main()
