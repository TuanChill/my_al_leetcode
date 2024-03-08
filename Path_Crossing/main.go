package main

import "fmt"

// 1496. Path Crossing
// https://leetcode.com/problems/path-crossing/description/

/*
Given a string path, where path[i] = 'N', 'S', 'E' or 'W',
each representing moving one unit north, south, east, or west, respectively.
You start at the origin (0, 0) on a 2D plane and walk on the path specified by path.

Return true if the path crosses itself at any point, that is, if at any time you are on a location you have previously visited. Return false otherwise.
*/

// Solution

/*

N - +0 -> x, +1 -> y
E - +1 -> x, +0 -> y
S - +0 -> x, -1 -> y
W - -1 -> x, +0 -> y

*/

type Point struct {
	x, y int
}

func (point *Point) cross(x, y int) {
	point.x += x
	point.y += y
}

func isPathCrossing(path string) bool {
	startPoint := Point{
		x: 0,
		y: 0,
	}

	visitedPoint := make(map[Point]bool)

	visitedPoint[startPoint] = true

	for _, v := range path {
		switch v {
		case 'N':
			startPoint.cross(0, 1)
		case 'S':
			startPoint.cross(0, -1)
		case 'E':
			startPoint.cross(1, 0)
		case 'W':
			startPoint.cross(-1, 0)
		}
		if _, exists := visitedPoint[startPoint]; exists {
			return true
		} else {
			visitedPoint[startPoint] = true
		}
	}

	return false
}

func main() {
	fmt.Println(isPathCrossing("NESWW"))
}
