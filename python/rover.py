class Rover(object):

    def __init__(self, x, y, d, plateau):
        self.x_coordinates = x
        self.y_coordinates = y
        self.direction = d
        self.plateau = plateau

    def setCoordinatesX(self, value):
        self.x_coordinates = value

    def setCoordinatesY(self, value):
        self.y_coordinates = value

    def setDirection(self, value):
        self.direction = value

    def turnRight(self):
        self.direction = {
            'N': 'E',
            'E': 'S',
            'S': 'W',
            'W': 'N',
        }[self.direction]

    def turnLeft(self):
        self.direction = {
            'N': 'W',
            'W': 'S',
            'S': 'E',
            'E': 'N',
        }[self.direction]

    def run(self, commands):
        for command in commands:
            if command == "L":
                self.turnLeft();
            elif command == "R":
                self.turnRight()
            elif command == "M":
                self.move()

    def move(self):
        if self.direction == 'N' and self.plateau.maxY > self.y_coordinates:
            self.y_coordinates += 1
        elif self.direction == 'E' and self.plateau.maxX > self.y_coordinates:
            self.x_coordinates += 1
        elif self.direction == 'S' and self.plateau.maxY > self.y_coordinates:
            self.y_coordinates -= 1
        elif self.direction == 'W' and self.plateau.maxX > self.y_coordinates:
            self.x_coordinates -= 1
    @property
    def position(self):
        return str(self.x_coordinates) + " " + str(self.y_coordinates) + " " + self.direction
