package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew" // for debuggin purposes
	prmt "github.com/gitchander/permutation"
	"github.com/gorilla/mux"
	"github.com/umahmood/haversine"
	"log"
	"net/http"
)

//// haversine library and optimal path finding

// CoordArray will be exported when it is moved out to its own package
type CoordArray []haversine.Coord

func (ps CoordArray) Len() int      { return len(ps) }
func (ps CoordArray) Swap(i, j int) { ps[i], ps[j] = ps[j], ps[i] }

var testList []haversine.Coord = []haversine.Coord{
	haversine.Coord{Lat: 10.1, Lon: 20.2},
	haversine.Coord{Lat: 10, Lon: 1003.3},
	haversine.Coord{Lat: 10, Lon: 12.},
}

func PathLength(path CoordArray) (mi, km float64) {
	var totalDistanceKm float64 = 0
	var totalDistanceMi float64 = 0

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
		fmt.Println(mi, km, path)
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
	fmt.Println("optimial distance")
	fmt.Println(bestDistanceMi, bestDistanceKm, bestPath)
	return bestDistanceMi, bestDistanceKm, bestPath
}

// Router will need ot be exported when it is moved into its own package later on
func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/optimize_route", optimizeUserPath).Methods("POST")
	return router
}

// example json request
// {
// 	"length": 3,
// 	"points": [
// 		{"lat": 10.2, "lon": 90.2},
// 		{"lat": 920.2, "lon": 123.111},
// 		{"lat": 32.2, "lon": 153.22}
// 	]
// }

type userRequest struct {
	length int
	path   []haversine.Coord
}

// middleware - wrappers around optimization functions to read json and return json
func optimizeUserPath(w http.ResponseWriter, r *http.Request) {
	var temp userRequest

	err := json.NewDecoder(r.Body).Decode(&temp)
	if err != nil {
		// Panic(err)
	}

	r.ParseForm()
	spew.Sdump(r.Body)
	log.Printf("%+v", temp)

	bestRoute(temp.path)
}

func main() {
	fmt.Println(testList)
	bestRoute(testList)

	r := Router()
	fmt.Printf("Starting server on the port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
