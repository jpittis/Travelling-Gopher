/* Travelling Gopher
Solving the travelling salesman problem
with an evolutionary algorithm.

Jake Pittis, December 2014

Let's see how creative the algorithm will get.
Maybe it will revolutionize the world! */

package main

import (
    "fmt"
    "math/rand"
    "math"
)

/* Holds the x and y coordinates of a location. */
type location struct {
    x int
    y int
}

type trip struct {
    locations []location
    totalDistance float64
}

/* Creates a random location with given maximums. */
func randLocation(xMax int, yMax int) location {
    x := rand.Intn(xMax * 2) + 1 - xMax 
    y := rand.Intn(yMax * 2) + 1 - yMax
    return location{x: x, y: y}
}

/* Creates list of locations with given maximums. */
func randLocations(numLocations int, xMax int, yMax int) []location {
    trip := make([]location, numLocations)
    for i := 0; i < numLocations; i++ {
        trip[i] = randLocation(xMax, yMax)
    }
    return trip
}

/* Return distance between two given locations.
TODO: sqrt may not be needed */
func distanceBetween(a location, b location) float64 {
    return math.Sqrt(math.Pow(float64(a.x - b.x), float64(2)) + math.Pow(float64(a.y - b.y), float64(2)))
}

/* Return distance of round trip. */
func totalDistance(locations []location) float64 {
    var total float64
    length := len(locations)
    for i := 0; i < length - 1; i++ {
        total += distanceBetween(locations[i], locations[i + 1])
    }
    total += distanceBetween(locations[length - 2], locations[length - 1])
    return total
}

/* Creates a trip structures by combining given locations
and calculating total distance. */
func newTrip(locations []location) trip {
    return trip{locations: locations, totalDistance: totalDistance(locations)}
}

/* Generates a new random trip and shuffles it into given number of combinations. */
func newGeneration(numTrips int, numLocations int, xMax int, yMax int) []trip {
    locations := randLocations(numLocations, xMax, yMax)
    generation := make([]trip, numTrips)
    for i := 0; i < numTrips; i++ {
        generation[i] = newTrip(shuffleLocations(locations))
    }
    return generation
}

/* Returns a randomly shuffled set of locations. */
func shuffleLocations(locations []location) []location {
    order := rand.Perm(len(locations))
    result := make([]location, len(locations))
    for i, j := range order {
        result[j] = locations[i]
    }
    return result
}

/* Given a generation of trips, create n choose 2 children.
TODO: mutation in this step? */
func makeChildren(generation []trip, numChildren int) []trip {
    length := len(generation)
    children := make([]trip, 0, length)
    for i := 0; i < length; i++ {
        for j := i; j <length; j++ {
            if i != j {
                child := makeChild(generation[i].locations, generation[j].locations)
                // room for mutation?
                children = append(children, newTrip(child))
            }
        }
    }
    return children
}

/* Returns a child of the two given locations. */
func makeChild(a []location, b []location) []location {
    length := len(a)
    child := make([]location, length)

    start := rand.Intn(length)
    end := rand.Intn(length)
    /* Index from start to end will be from parent a. The rest is from parent b. */
    for i := 0; i < length; i++ {
        if i >= start && i <= end {
            child[i] = a[i]
        } else {
            child[i] = b[i]  
        }
    }
    return child
}

/* Returns the given number of smallest trips.
TODO: base case of len(generation) == 0 will crash */
func getSmallest(generation []trip, numTrips int) []trip {
    smallest := make([]int, 0, numTrips)
    for i := 0; i < numTrips; i++ {
        index := 0
        for j := 1; j < len(generation); j++ {
            if generation[j].totalDistance < generation[index].totalDistance {
                index = j
            }
        }
        smallest = append(smallest, index)
    }
    return getGenerationIndexes(generation, smallest)
}

/* Returns new slice with values at indexes. */
func getGenerationIndexes(generation []trip, indexes []int) []trip {
    length := len(indexes)
    result := make([]trip, length)
    for i := 0; i < length; i++ {
        result[i] = generation[indexes[i]]
    }
    return result
}

/* Print out a generation. */
func printGeneration(generation []trip) {
    for i := 0; i < len(generation); i++ {
        fmt.Printf("%s\n", generation[i])
    }
}

/* Create a random generations of trips. Run the evolutionary loop.g */
func main() {
    generation := newGeneration(10, 4, 10, 10)
    printGeneration(generation)
    fmt.Print("----------\n")
    //generation = makeChildren(generation, 3)
    generation = getSmallest(generation, 3)
    printGeneration(generation)
}