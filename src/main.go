package main

import (
	"fmt"
	"github.com/umahmood/haversine"

	prmt "github.com/gitchander/permutation"
)

type CoordArray []haversine.Coord

func (ps CoordArray) Len() int      { return len(ps) }
func (ps CoordArray) Swap(i, j int) { ps[i], ps[j] = ps[j], ps[i] }

var testList []haversine.Coord = []haversine.Coord{
	haversine.Coord{Lat: 10.1, Lon: 20.2},
	haversine.Coord{Lat: 10, Lon: 1003.3},
	haversine.Coord{Lat: 10, Lon: 12.},
}

func pathLength(path CoordArray) (mi, km float64) {
	var totalDistanceKm float64 = 0
	var totalDistanceMi float64 = 0

	for i := 1; i < len(path); i++ {
		km, mi = haversine.Distance(path[i-1], path[i])
		totalDistanceKm += km
		totalDistanceMi += mi
	}
	return totalDistanceKm, totalDistanceMi
}

func bestRoute(path) (mi, km float64, path CoordArray) {
	p := pmrt.New(path)

	var bestDistanceKm float64
	var bestDistanceMi float64
	var bestPath CoordArray

	for p.next() {
		mi, km := pathLength(path)
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
	return bestDistanceMi, bestDistanceKm, path
}

func main() {
	fmt.Println(testList)
}
