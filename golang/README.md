# Mars Rovers

Solution of Mars Rover problem

### Requirements

    Docker
    Docker Compose


### Installation
docker-compose up
   
    app_1  | 2020/02/08 11:08:00 Running build command!
    app_1  | 2020/02/08 11:08:00 Build ok.
    app_1  | 2020/02/08 11:08:00 Restarting the given command.
    app_1  | 2020/02/08 11:08:00 stdout: 1 1 S
    app_1  | 2020/02/08 11:08:00 stdout: 5 1 E

### Tests
    docker-compose up -d
    docker-compose exec app /bin/bash

Sample Result:
    
    === RUN   TestRoverTurnLeft
    --- PASS: TestRoverTurnLeft (0.00s)
    === RUN   TestRoverTurnRight
    --- PASS: TestRoverTurnRight (0.00s)
    === RUN   TestRoverMoveEast
    --- PASS: TestRoverMoveEast (0.00s)
    === RUN   TestRoverMovWest
    --- PASS: TestRoverMovWest (0.00s)
    === RUN   TestRoverMoveNorth
    --- PASS: TestRoverMoveNorth (0.00s)
    === RUN   TestRoverMoveSouth
    --- PASS: TestRoverMoveSouth (0.00s)
    === RUN   TestRoverCase1
    --- PASS: TestRoverCase1 (0.00s)
    === RUN   TestRoverCase2
    --- PASS: TestRoverCase2 (0.00s)
    PASS
    ok  	mars_rover	0.007s




