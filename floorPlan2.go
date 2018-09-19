// find the distance from any room to the nearest guard, considering locked rooms
package main

import (
	"container/list"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const GUARD = -2
const LOCKED = -1
const UNREACHABLE = 0

// maxes are 1-indexed
type grid struct {
	// [row][colum]
	dist       [][]int
	cMax, rMax int
}
type room struct {
	r, c int
}

// create a new grid of arbitrary size
func newGrid(rMax, cMax int) *grid {
	var g *grid = new(grid)
	g.dist = make([][]int, rMax)
	for r := 0; r < rMax; r++ {
		g.dist[r] = make([]int, cMax)
	}
	g.cMax = cMax
	g.rMax = rMax
	return g
}

// is a room in a grid?
func (g *grid) contains(r room) bool {
	if r.r >= 0 && r.r < g.rMax && r.c >= 0 && r.c < g.cMax {
		return true
	}
	return false
}

// is a room valid
func (g *grid) validRoom(r room) {
	if !g.contains(r) {
		panic(fmt.Sprintf("room{%d, %d} is not a valid room\n", r.r, r.c))
	}
}

// mark a room as a guard
func (g *grid) placeGuard(r room) {
	g.validRoom(r)
	g.dist[r.r][r.c] = GUARD
}

// is a room guarded?
func (g *grid) guarded(r room) bool {
	g.validRoom(r)
	if g.dist[r.r][r.c] == GUARD {
		return true
	}
	return false
}

//lock a room
func (g *grid) lockRoom(r room) {
	g.validRoom(r)
	g.dist[r.r][r.c] = LOCKED
}

//is a room locked?
func (g *grid) locked(r room) bool {
	g.validRoom(r)
	if g.dist[r.r][r.c] == LOCKED {
		return true
	}
	return false
}

// find all the rooms marked as guards in the distance map
func (g *grid) findGuards() []room {
	var guards []room = make([]room, 0)
	for r := 0; r < g.rMax; r++ {
		for c := 0; c < g.cMax; c++ {
			if g.guarded(room{r, c}) {
				guards = append(guards, room{r, c})
			}
		}
	}
	return guards
}
func (g *grid) String() string {
	var sb strings.Builder
	for r := 0; r < g.rMax; r++ {
		for c := 0; c < g.cMax; c++ {
			if g.locked(room{r, c}) {
				sb.WriteRune(rune('L'))
			} else if g.guarded(room{r, c}) {
				sb.WriteRune(rune('G'))
			} else {
				sb.WriteString(strconv.Itoa(g.dist[r][c]))
			}
			if c != g.cMax-1 {
				sb.WriteRune(rune(' '))
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
func (r room) proposeNeighbors() []room {
	return []room{room{r.r - 1, r.c}, room{r.r, r.c - 1}, room{r.r + 1, r.c}, room{r.r, r.c + 1}}
}

// returns the neighbors (adjacent rooms) of a room, iff in the grid and not locked or a guard
func (g *grid) neighbors(r room) []room {
	var n []room = make([]room, 0)
	var c []room = r.proposeNeighbors()
	for i := range c {
		var cr room = c[i]
		if g.contains(cr) && !g.guarded(cr) && !g.locked(cr) {
			n = append(n, cr)
		}
	}
	return n
}

// for each guard, visit all the rooms and mark their distance.  track which rooms have been visited
// for any given guard, to prevent multiple-visitation
func (g *grid) visitRoomsViaGuards(guards []room) {
	for gI := range guards {
		guard := guards[gI]
		neighbors := g.neighbors(guard)
		var toVisit *list.List = list.New() // using a linked list as a queue
		for i := range neighbors {
			toVisit.PushBack(neighbors[i])
		}
		var visited map[room]bool = make(map[room]bool) // using a hashmap as a filter on the queue
		g.calcDistances(guard, toVisit, visited)
	}
}

// iterate through the queue and calculate the distance for each room, skipping visited rooms
func (g *grid) calcDistances(guard room, toVisit *list.List, visited map[room]bool) {
	e := toVisit.Front()
	for e != nil {
		var r room = e.Value.(room)
		if visited[r] {
			e = e.Next()
			continue
		}
		g.calcDistance(guard, r, visited)
		neighbors := g.neighbors(r)
		for i := range neighbors {
			if !visited[neighbors[i]] {
				toVisit.PushBack(neighbors[i])
			}
		}
		e = e.Next()
	}
}

// calculate the distance from a guard to a room by inspecting the rooms neighbors.  skip visited rooms
func (g *grid) calcDistance(guard room, r room, visited map[room]bool) {
	if abs(guard.r-r.r)+abs(guard.c-r.c) == 1 { //they are neighbors
		g.dist[r.r][r.c] = 1
	} else { // peek at the rooms neighbors and pick the lowest non-zero value > 0
		var low int = int(math.MaxInt32)
		var neighbors []room = g.neighbors(r)
		for i := range neighbors {
			n := neighbors[i]
			if g.dist[n.r][n.c] > 0 && g.dist[n.r][n.c] < low {
				low = g.dist[n.r][n.c]
			}
		}
		// ensure only to pick a new lowest if the room is vistable by a guard
		if g.dist[r.r][r.c] == 0 {
			g.dist[r.r][r.c] = low + 1
		} else if low < g.dist[r.r][r.c] {
			g.dist[r.r][r.c] = low + 1
		}
	}
	visited[r] = true
}
func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func main() {
	fmt.Println("hello")
	g := newGrid(10, 10)
	g.placeGuard(room{3, 3})
	g.placeGuard(room{2, 4})
	g.placeGuard(room{1, 1})
	g.placeGuard(room{6, 6})
	g.lockRoom(room{1, 2})
	g.lockRoom(room{4, 4})
	g.lockRoom(room{2, 3})
	g.lockRoom(room{8, 0})
	g.lockRoom(room{8, 1})
	g.lockRoom(room{8, 2})
	g.lockRoom(room{8, 3})
	g.lockRoom(room{8, 4})
	g.lockRoom(room{8, 5})
	g.lockRoom(room{8, 6})
	g.lockRoom(room{8, 7})
	g.lockRoom(room{8, 8})
	g.lockRoom(room{8, 9})
	fmt.Println(g)
	g.visitRoomsViaGuards(g.findGuards())
	fmt.Println(g)
}
