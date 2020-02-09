# Mars Rovers

Solution of Mars Rover problem 

### Requirements

    Docker
    Docker Compose


### Installation
    docker-compose up -d
    docker-compose exec app python /app/main.py

    1 3 N
    5 1 E

### Tests
docker-compose exec app python /app/tests.py -v

    testRoverCase1 (__main__.TestRoverS) ... ok
    testRoverMoveEast (__main__.TestRoverS) ... ok
    testRoverMoveNorth (__main__.TestRoverS) ... ok
    testRoverMoveSouth (__main__.TestRoverS) ... ok
    testRoverMoveWest (__main__.TestRoverS) ... ok
    testRoverTurnLeft (__main__.TestRoverS) ... ok
    testRoverTurnRight (__main__.TestRoverS) ... ok
    
    ----------------------------------------------------------------------
    Ran 7 tests in 0.003s
    
    OK




