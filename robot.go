package main

import (
	"container/list"
	"fmt"
)

const size = 5

type point struct {
	r, c int
}

func findRoute(cur, goal point, grid [size][size]bool, path *list.List) bool {
	if cur == goal {
        path.PushBack(cur)
		return true
    } else if cur.r >= size || cur.c >= size {
        return false
	} else if cur.r+1 < size && !grid[cur.r+1][cur.c] && findRoute(point{cur.r + 1, cur.c}, goal, grid, path) || cur.c+1<size && !grid[cur.r][cur.c+1] && findRoute(point{cur.r, cur.c + 1}, goal, grid, path) {
		path.PushBack(cur)
		return true
	} else {
		return false
	}
}
func main() {
	var grid [size][size]bool
	grid[1][3] = true
    grid[3][1] = true
    grid[1][4] = true
	grid[1][2] = true
	grid[1][0] = true
	grid[1][1] = true
	var robot point = point{0, 0}
	var goal point = point{4, 4}
	var path *list.List = list.New()
	fmt.Println(findRoute(robot, goal, grid, path))
	e := path.Back()
	for e != nil {
		fmt.Println(e.Value.(point))
		e = e.Prev()
	}
}
