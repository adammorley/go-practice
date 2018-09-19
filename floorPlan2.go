package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const GUARD = -2
const LOCKED = -1

// row column; maxes are 1-indexed
type grid struct {
	dist       [][]int
	cMax, rMax int
}
type room struct {
	r, c int
}

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
func (g *grid) placeGuard(r room) {
	g.dist[r.r][r.c] = GUARD
}
func (g *grid) guarded(r room) bool {
	if g.dist[r.r][r.c] == GUARD {
		return true
	}
	return false
}
func (g *grid) lockRoom(r room) {
	g.dist[r.r][r.c] = LOCKED
}
func (g *grid) locked(r room) bool {
	if g.dist[r.r][r.c] == LOCKED {
		return true
	}
	return false
}
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
func (g *grid) contains(r room) bool {
	if r.r >= 0 && r.r < g.rMax && r.c >= 0 && r.c < g.cMax {
		return true
	}
	return false
}
func proposeNeighbors(r room) []room {
	return []room{room{r.r - 1, r.c}, room{r.r, r.c - 1}, room{r.r + 1, r.c}, room{r.r, r.c + 1}}
}

// returns the neighbors (adjacent rooms) of a room, iff in the grid and not locked or a guard
func (g *grid) neighbors(r room) []room {
	var n []room = make([]room, 0)
	var c []room = proposeNeighbors(r)
	for i := range c {
		if g.contains(c[i]) && !g.guarded(c[i]) && !g.locked(c[i]) {
			n = append(n, c[i])
		}
	}
	return n
}
func (g *grid) visitRoomsViaGuards(guards []room) {
	for gI := range guards {
		// track whether this room has been visited yet
		var visited map[room]bool = make(map[room]bool)
		guard := guards[gI]
		g.calcDistances(guard, g.neighbors(guard), visited)
	}
}
func (g *grid) calcDistances(guard room, rooms []room, visited map[room]bool) {
	if len(rooms) != 0 {
		for r := range rooms {
			g.calcDistance(guard, rooms[r], visited)
		}

	}
}
func (g *grid) calcDistance(guard room, r room, visited map[room]bool) {
	fmt.Println("examining room", r, "from guard location", guard)
	if visited[r] {
		return
	}
	var neighbors []room = g.neighbors(r)
	if abs(guard.r-r.r)+abs(guard.c-r.c) == 1 { //they are neighbors
		g.dist[r.r][r.c] = 1
	} else { // peek at the rooms neighbors and pick the lowest non-zero value > 0
		var low int = int(math.MaxInt32)
		for i := range neighbors {
			n := neighbors[i]
			if g.dist[n.r][n.c] > 0 && g.dist[n.r][n.c] < low {
				low = g.dist[n.r][n.c]
			}
		}
        if g.dist[r.r][r.c] == 0 {
            g.dist[r.r][r.c] = low + 1
        } else if low < g.dist[r.r][r.c] {
            g.dist[r.r][r.c] = low + 1
        }
	}
	visited[r] = true
	g.calcDistances(guard, neighbors, visited)
}
func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func main() {
	fmt.Println("hello")
	g := newGrid(5, 5)
	g.placeGuard(room{3, 3})
	g.placeGuard(room{2, 4})
	g.placeGuard(room{1, 1})
	g.lockRoom(room{1, 2})
	g.lockRoom(room{4, 4})
	g.lockRoom(room{2, 3})
	fmt.Println(g)
	g.visitRoomsViaGuards(g.findGuards())
	fmt.Println(g)
}
