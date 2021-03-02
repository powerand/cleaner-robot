# Cleaner robot

GIVEN API methods:
* TurnLeft
* TurnRight
* Move
* Clean

TODO: Clean every cell in room, which has impassable areas

EXAMPLE of room:

* `X  X  X  X  X  X  X  X`
* `X  0  0  0  0  X  X  X`
* `X  0  0  0  X  0  X  X`
* `X  0  P  X  X  0  X  X`
* `X  0  0  0  0  0  0  X`
* `X  0  0  X  X  0  X  X`
* `X  0  0  0  X  0  X  X`
* `X  X  X  X  X  X  X  X`

SYMBOLS:
* X - impassable area
* 0 - cell in room, which need to clean
* P - start point
