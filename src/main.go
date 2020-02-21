package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"time"
	// "github.com/davecgh/go-spew/spew" // for debuggin purposes
	"log"
	"net/http"

	prmt "github.com/gitchander/permutation"
	"github.com/gorilla/mux"
	"github.com/umahmood/haversine"
)

//// haversine library and optimal path finding

// CoordArray will be exported when it is moved out to its own package
type CoordArray []haversine.Coord

func (ps CoordArray) Len() int      { return len(ps) }
func (ps CoordArray) Swap(i, j int) { ps[i], ps[j] = ps[j], ps[i] }

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

// server constance
const (
	STATIC_DIR = "/views/"
	PORT       = "8000"
)

// Router will need ot be exported when it is moved into its own package later on
func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/optimize_route", optimizeUserPath).Methods("POST")

	router.
		PathPrefix(STATIC_DIR).
		Handler(http.StripPrefix(STATIC_DIR, http.FileServer(http.Dir("."+STATIC_DIR))))

	return router
}

type userRequest struct {
	Length int
	Path   []haversine.Coord
}

// middleware - wrappers around optimization functions to read json and return json
func optimizeUserPath(w http.ResponseWriter, r *http.Request) {
	var temp userRequest

	json.NewDecoder(r.Body).Decode(&temp)
	log.Printf("%+v", temp)

	beforeCalculation := time.Now()
	mi, km, path := bestRoute(temp.Path)
	afterCalculation := time.Now()
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
