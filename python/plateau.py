class Plateau(object):
    def __init__(self, width, height, min_width=0, min_height=0):
        self.minY = min_height
        self.minX = min_width
        self.maxX = width
        self.maxY = height
