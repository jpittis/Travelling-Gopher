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

/* Creates a random location with given maximums. */
func randLocation(xMax int, yMax int) location {
    x := rand.Intn(xMax * 2) + 1 - xMax 
    y := rand.Intn(yMax * 2) + 1 - yMax
    return location{x: x, y: y}
}

/* Creates list of locations with given maximums. */
func randTrip(numLocations int, xMax int, yMax int) []location {
    trip := make([]location, numLocations)
    for i := 0; i < numLocations; i++ {
        trip[i] = randLocation(xMax, yMax)
    }
    return trip
}

/* Return distance between two given locations. */
func distanceBetween(a location, b location) float64 {
    return math.Sqrt(math.Pow(float64(a.x - b.x), float64(2)) + math.Pow(float64(a.y - b.y), float64(2)))
}

/* Return distance of round trip. */
func totalDistance(trip []location) float64 {
    var total float64
    for i := 0; i < len(trip) - 1; i++ {
        total += distanceBetween(trip[i], trip[i + 1])
    }
    total += distanceBetween(trip[len(trip) - 2], trip[len(trip) - 1])
    return total
}

/* Generates a random trip and finds the distance. */
func main() {
    trip := randTrip(10, 100, 100)
    fmt.Printf("%f\n", totalDistance(trip))
}