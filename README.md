astar
=====

A small package that implements the A\* (A star) pathfinding algorithm.

Example
-------

	import (
		"fmt"
		"github.com/sperre/astar"
	)

	func main() {
		// If your map is something more than 2D matrix then you might want to modify adjacentNodes
		map_data := astar.GetMapFromImage(test.png)

		// Returns a list of nodes from start to stop avoiding all, obstacles if possible
		startx, starty :=  0,  0
		stopx, stopy   := 99, 99
		shortest_path := astar.Astar(map_data, startx, starty, stopx, stopy)
	}

Origin
------

This library was developed for Eurobot-NTNU 2012,
based on code from: https://github.com/humanfromearth/gopathfinding
