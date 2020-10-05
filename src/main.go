package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"time"

	// "github.com/davecgh/go-spew/spew" // for debuggin purposes
	"log"
	"net/http"
	"net/http/httputil"

	prmt "github.com/gitchander/permutation"
	"github.com/gorilla/mux"
	"github.com/umahmood/haversine"
)

//// haversine library and optimal path finding

// CoordArray will be exported when it is moved out to its own package
type CoordArray []haversine.Coord

// Len to fufil interface
func (ps CoordArray) Len() int { return len(ps) }

// Swap to fufil interface
func (ps CoordArray) Swap(i, j int) { ps[i], ps[j] = ps[j], ps[i] }

// PathLength is meant to find length between two paths
func PathLength(path CoordArray) (mi, km float64) {
	var totalDistanceKm float64 // 0 is the default zero value for float64
	var totalDistanceMi float64

	for i := 1; i < len(path); i++ {
		km, mi = haversine.Distance(path[i-1], path[i])
		totalDistanceKm += km
		totalDistanceMi += mi
	}
	return totalDistanceKm, totalDistanceMi
}

// TODO
// all permutations is different from all unique paths
// optimize for speed
func bestRoute(path CoordArray) (float64, float64, CoordArray) {
	p := prmt.New(path)

	var bestDistanceKm float64
	var bestDistanceMi float64
	var bestPath CoordArray

	for p.Next() {
		mi, km := PathLength(path)
		// add logging to ensure its working
		if bestPath == nil {
			// fist pass - set everything
			bestDistanceKm = km
			bestDistanceMi = mi
			bestPath = path
		}
		if km < bestDistanceKm {
			// store new distance
			bestDistanceKm = km
			bestDistanceMi = mi
			bestPath = path
		}
	}
	return bestDistanceMi, bestDistanceKm, bestPath
}

// DefinedCoordiante name captialized because we expect to move this out to a seperate golang package
type DefinedCoordiante struct {
	id         int
	coordinate haversine.Coord
}

// DefinedCoordinateArray also to be exported
type DefinedCoordinateArray []DefinedCoordiante

// Len to fufil interface
func (ps DefinedCoordinateArray) Len() int { return len(ps) }

// Swap to fufil interface
func (ps DefinedCoordinateArray) Swap(i, j int) { ps[i], ps[j] = ps[j], ps[i] }

type distance struct {
	mi float64
	km float64
}

// subpath used to identify two points uniquely
type subPath struct {
	startID, endID int
}

type costTable map[subPath]distance

func bestRouteCached(points DefinedCoordinateArray) (float64, float64, DefinedCoordinateArray) {
	costTable := generateCostTable(points)

	p := prmt.New(points) // might have to ensure an interface exists for this

	var bestDistanceKm float64
	var bestDistanceMi float64
	var bestPath DefinedCoordinateArray

	for p.Next() {
		mi, km := findCachedPathLength(points, costTable)
		// add logging to ensure its working
		if bestPath == nil {
			// fist pass - set everything
			bestDistanceKm = km
			bestDistanceMi = mi
			bestPath = points
		}
		if km < bestDistanceKm {
			// store new distance
			bestDistanceKm = km
			bestDistanceMi = mi
			bestPath = points
		}
	}
	return bestDistanceMi, bestDistanceKm, bestPath
}

func findCachedDistance(firstID int, secondID int, costTable costTable) distance {
	return costTable[subPath{firstID, secondID}]
}

func findCachedPathLength(path DefinedCoordinateArray, costTable costTable) (mi, km float64) {
	var totalDistanceKm float64 // 0 is the default zero value for float64
	var totalDistanceMi float64

	for i := 1; i < len(path); i++ {
		distance := costTable[subPath{path[i-1].id, path[i].id}]
		totalDistanceKm += distance.km
		totalDistanceMi += distance.mi
	}
	return totalDistanceKm, totalDistanceMi
}

func generateCostTable(points DefinedCoordinateArray) costTable {
	// create map
	// iterate over all combinations to memoize them
	// generate a table memoizing any costs we could make going forward
	costTable := make(costTable)
	for i := 0; i < (len(points) - 1); i++ {
		for k := i + 1; k < len(points); k++ {
			// cost route between the two and fill both entries on the costTable
			mi, km := haversine.Distance(points[i].coordinate, points[k].coordinate)

			leftToRight := subPath{points[i].id, points[k].id}
			rightToLeft := subPath{points[k].id, points[i].id}
			calculatedDistance := distance{mi, km}

			costTable[leftToRight] = calculatedDistance
			costTable[rightToLeft] = calculatedDistance
		}
	}
	return costTable
}

// TODO: implement a simulated annealing implementation now
// start with a random solution A and generate a related canditate solution B
// if B is better than A we want to use B and restart. If A is better than B then we want to explore
// use a temperature function to decide if we should explore candidate solutions in the region of B

// wiki pseudo code

// Let s = s0
// For k = 0 through kmax (exclusive):
// 	T ← temperature( (k+1)/kmax )
// 	Pick a random neighbour, snew ← neighbour(s)
// 	If P(E(s), E(snew), T) ≥ random(0, 1):
// 		s ← snew
// Output: the final state s

// Tsp stands for traveling salesman and the struct encompases all of the recordkeeping required that will persist
// through itterations.
type Tsp struct {
	totalPoints          int
	bestRoute            DefinedCoordinateArray
	bestRouteDistance    float64
	currentRoute         DefinedCoordinateArray
	currentRouteDistance float64
	temp                 float64
	distanceCache        map[subPath]distance
}

func (tsp *Tsp) degradeTemp() {
	tsp.temp = tsp.temp * .9
}

// based on the current temperature find a relates path
// we expect a higher temperature to cause a neighboring path with more differences when
// compared to a lower temp
func (tsp *Tsp) neighboringPath() DefinedCoordinateArray {
	neighboringPath := tsp.currentRoute
	upperBound := tsp.totalPoints - 1
	var pointA, pointB int // initialized with a 0 value
	for {
		// pick random points until we have two different cities selected
		pointA = rand.Intn(upperBound)
		pointB = rand.Intn(upperBound)

		if pointA != pointB {
			break
		}
	}
	// swap points in the base route to get a neighboring route
	neighboringPath[pointA], neighboringPath[pointB] = neighboringPath[pointB], neighboringPath[pointA]

	return neighboringPath
}

// calulateSubPathLength takes a subPathPair which is defined as two points that follow one another
// if the id of the subpath is not already memoized it is calculated out and stored in the map
func (tsp *Tsp) calculateSubPathLength(pointA, pointB DefinedCoordiante) distance {
	pairID := subPath{pointA.id, pointB.id}
	if dist, ok := tsp.distanceCache[pairID]; ok {
		return dist
	} else {
		// we need to calculate it
		mi, km := haversine.Distance(pointA.coordinate, pointB.coordinate)
		dist = distance{mi, km}
		// memozie it
		altPairID := subPath{pointB.id, pointA.id}
		tsp.distanceCache[pairID] = dist
		tsp.distanceCache[altPairID] = dist
		// return it
		return dist
	}
}

func (tsp *Tsp) calculatePathLength(path DefinedCoordinateArray) distance {
	totalDist := distance{0, 0}
	for i := 1; i < len(path); i++ {
		nextPoint := path[i]
		currPoint := path[i-1]
		stepDist := tsp.calculateSubPathLength(currPoint, nextPoint)
		totalDist.mi += stepDist.mi
		totalDist.km += stepDist.km
	}
	return totalDist
}

func (tsp *Tsp) acceptanceProbability(newDist float64) float64 {
	return math.Exp((tsp.currentRouteDistance - newDist) / temperature)
}

func (tsp *Tsp) acceptNewPath(proposedPathDist float64, proposedPath DefinedCoordinateArray) {
	if proposedPathDist <= tsp.currentRouteDistance {
		tsp.currentRoute = proposedPath
		tsp.currentRouteDistance = proposedPathDist
		if proposedPathDist < tsp.bestRouteDistance {
			tsp.bestRoute = proposedPath
			tsp.bestRouteDistance = proposedPathDist
		}
	} else {
		// acceptanceProbability = tsp.acceptanceProbability(proposedPathDist)
		// randomVal = rand.float64()
		// if acceptanceProbability > randomVal {
		// 	tsp.currentRoute = proposedPath
		// 	tsp.currentRouteDistance = proposedPathDist
		}
	}
}

// server constants
const (
	PORT = "8000"
)

// Router will need ot be exported when it is moved into its own package later on
func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/optimize_route/", optimizeUserPath).Methods("POST")
	router.HandleFunc("/pulse", pulse).Methods("GET", "POST")
	// router.HandleFunc("/", whatTheFuck)

	return router
}

type userRequest struct {
	Length int
	Path   []haversine.Coord
}

func pulse(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("call to pulse triggered!")
	fmt.Fprint(w, "{'message': 'alive and accessible'}")
}

// middleware - wrappers around optimization functions to read json and return json
func optimizeUserPath(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Optimizing user path!\n")
	var temp userRequest
	log.Printf("data send in")

	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))

	// add error handeling
	json.NewDecoder(r.Body).Decode(&temp)
	log.Printf("%+v", temp)

	beforeCalculation := time.Now()
	mi, km, path := bestRoute(temp.Path)
	afterCalculation := time.Now()

	// agressive print out strategy
	fmt.Fprintf(w, "{\"mi\": %f, \"km\": %f, \"time\": \"%s\", \"path\": [", mi, km, afterCalculation.Sub(beforeCalculation))
	fmt.Fprintf(w, "{\"Lat\": %f, \"Lon\": %f}", path[0].Lat, path[0].Lon)
	for i := 1; i < len(path); i++ {
		fmt.Fprintf(w, ", {\"Lat\": %f, \"Lon\": %f}", path[i].Lat, path[i].Lon)
	}

	fmt.Fprintf(w, "]}")
}

func main() {
	r := Router()
	fmt.Printf("Starting server on the port %s...\n", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, r))
}
