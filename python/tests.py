import unittest

from plateau import Plateau
from rover import Rover


class TestRoverS(unittest.TestCase):

    def setUp(self):
        self.plateau = Plateau(5, 5)
        self.rover = Rover(1, 5, 'S', self.plateau)

    def testRoverTurnLeft(self):
        self.rover.turnLeft()
        self.assertEqual(self.rover.direction, "E")

    def testRoverTurnRight(self):
        self.rover.turnRight()
        self.assertEqual(self.rover.direction, "W")

    def testRoverMoveEast(self):
        self.rover = Rover(1, 2, 'E', self.plateau)
        self.rover.move()
        self.assertEqual(self.rover.position, "2 2 E")

    def testRoverMoveWest(self):
        self.rover = Rover(1, 2, 'W', self.plateau)
        self.rover.move()
        self.assertEqual(self.rover.position, "0 2 W")

    def testRoverMoveNorth(self):
        self.rover = Rover(1, 2, 'N', self.plateau)
        self.rover.move()
        self.assertEqual(self.rover.position, "1 3 N")

    def testRoverMoveSouth(self):
        self.rover = Rover(1, 2, 'S', self.plateau)
        self.rover.move()
        self.assertEqual(self.rover.position, "1 1 S")

    def testRoverCase1(self):
        self.rover = Rover(1, 2, 'N', self.plateau)
        self.rover.run("LMLMLMLMM")

        self.assertEqual(self.rover.position, "1 3 N")

    def testRoverCase2(self):
        self.rover = Rover(3, 3, 'E', self.plateau)
        self.rover.run("MMRMMRMRRM")

        self.assertEqual(self.rover.position, "5 1 E")


if __name__ == "__main__":
    unittest.main()
