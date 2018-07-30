/*
    an insurance company wants to send a floor plan (an n-by-m array of "rooms"), which is pre-seeded with a set of guards (denoted by -1 ), and locked rooms (denoted by -2) and unreachable rooms (denoted by -3).  the interface should return the floor plan with each "room" (a location in the x-y plane) coded with a value > 0 with the distance to the nearest guard.


thoughts:

    - insurance is a relatively undisrupted market and is likely not basing on guard distance; to innovate in the insurance market, this is likely the wrong entry point.  additionally, most museums have a number of technical components which augment guards such as motion detection, movement (of object) sensors, laser grid, etc.  most novel attacks against museums ignore guard locations by moving the guards around or distracting them to increase guard distance.  but conceptually, a motion detector is a lot like a guard that never has to take breaks, so the analysis can be useful either way.  since insurance is fairly undisrupted, the data set for this likely already exists inside of a company like barclay's or another large insurance company in london/new york (the two major global financial hubs), and is probably adjusted based on lived experience of the firms.
    - this would be a poor market entry choice for a net-new business as software isn't the defining characteristic but rather brand trust and corporate history (rightly or wrongly)
    - reminds me of Dijkstra’s and other shortest path algorithm mechanisms, which i have no prior experience in; those use existing costs, in this case, we're calculating costs
    - would want the guards and locked rooms to be specified independently so it's just plug-and-chug and reduces the amount of loops across the floor plan as floor size scales up
*/

package main

import "errors"
import "fmt"
import "math"

type coordinate struct {
    x int
    y int
}

/*
    iterate through the 2-d int slice to find all the guard locations, and store them somewhere
    iterate through entire array, push all elements onto a queue
    pop elements off the queue and calcuate distance from any guard to the square
    store the lowest value
    return error if locked room through which guard cannot pass (implement later; mvp)

    - need a calculate distance function to calculate the distance from one square to another square

    example of why having "special" rooms be notified separately is useful:
*/
/*func analyzeAndScore(plan [][]int) error {
    guards := make(map[coordinate]bool)
    // first, iterate across all the locations and remember the locations of the guards
    for i := 0; i < len(plan); i++ {
        for j := 0; j < len(plan[i]); j++ {
            if plan[i][j] == 0 {
                var guard coordinate = coordinate{i, j}
                guards[guard] = true
            } else if plan[i][j] == -1 {
                return errors.New("Locked rooms not supported yet")
            }
        }
    }
    ...
    return nil
}*/

// iterate through the rooms in the floor plan and update the distances; guards stores the location of the guards
// note that since this is a slice, the underlying data structure is modified; this is sub-optimal and would need to later be improved to ensure that the data is either locked (eg with a mutex) or a serialized version is passed to the function.
// the room is treated as an x-y coordinate system; buildings must be rectangular; locked rooms not supported yet
func analyzeAndScore(plan [][]int, guards map[coordinate]bool) error {
    var ySize int
    for x := 0; x < len(plan); x++ {
        if ySize == 0 {
            ySize = len(plan[x])
        } else if len(plan[x]) != ySize {
            return errors.New("room is not rectangular") // XXX move to type or support rectangular rooms
        }
        for y := 0; y < len(plan[x]); y++ {
            var distance int = math.MaxInt32
            for k, _ := range guards {
                if x == k.x && y == y.x {
                    // it's a guard
                    distance = -1
                    break
                }
                t := math.Abs(k.x - x) + math.Abs(k.y - y)
                if t < distance {
                    distance = t
                }
            }
            plan[x][y] = distance
        }
    }
    return nil
}

func main() {
    // build some example cases, supply to function
}
