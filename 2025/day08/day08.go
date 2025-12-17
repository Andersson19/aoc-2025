package day08

import (
	"cmp"
	"slices"
	"strings"

	"github.com/Andersson19/aoc-2025/internal/util"
)

type Circuit struct {
	connections []Point
}

type Point struct {
	X int
	Y int
	Z int
}

type Edge struct {
	A Point
	B Point
	distance int64
}



func PartOne(lines []string, extras ...any) any {

	connections := len(lines)
	if len(extras) > 0 {
		connections = extras[0].(int)
	}

	points := getBoxPoints(lines)
	edges := getBoxEdges(points)

	// add edges to circuit
	var circuits []Circuit
	for c := 0; c < connections; c++ {
		circuits, _ = addToCircuit(edges[c], circuits)
	}

	// find length of all connections in each circuit
	circuitSizes := make([]int, len(circuits))
	for i := range circuits {
		circuitSizes[i] = len(circuits[i].connections)
	}

	// sort sizes by descending order
	slices.SortFunc(circuitSizes, func(a,b int) int {
		switch {
    	case a < b:
    	    return 1
    	case a > b:
    	    return -1
    	default:
        	return 0
    }
	})
	
	return circuitSizes[0] * circuitSizes[1] * circuitSizes[2]
}

func PartTwo(lines []string, extras ...any) any {
	points := getBoxPoints(lines)
	edges := getBoxEdges(points)

	// add edges to circuit
	var circuits []Circuit

	// find first edge that joins all boxes into one circuit
	edge := Edge{}
	c := 0
	for {
		circuits, edge = addToCircuit(edges[c], circuits)
		
		if len(circuits) == 1 && len(circuits[0].connections) == len(lines) {
			return edge.A.X * edge.B.X
		}
		c += 1
	}
}

func getBoxPoints(lines []string) []Point {
	points := make([]Point, len(lines))
	for i, line := range lines {
		nums := strings.Split(line, ",")
		x := util.Atoi(nums[0])
		y := util.Atoi(nums[1])
		z := util.Atoi(nums[2])
		points[i] = Point{X: x, Y: y, Z: z}
	}
	return points
}

func getBoxEdges(points []Point) []Edge {
	var edges []Edge
	for i, outer := range points[:len(points)-1] {
		for _, inner := range points[i+1:] {
			dx := int64(outer.X - inner.X)
			dy := int64(outer.Y - inner.Y)
			dz := int64(outer.Z - inner.Z)
			distance := dx*dx + dy*dy + dz*dz
			edges = append(edges, Edge{A: outer, B: inner, distance: distance})
		}
	}

	slices.SortFunc(edges, func(a,b Edge) int {
		return cmp.Compare(a.distance,b.distance)
	})
	return edges
}

func addToCircuit(edge Edge, circuits []Circuit) ([]Circuit, Edge) {
	pointA := edge.A
	pointB := edge.B

	if circuits == nil {
		circuits = append(circuits, Circuit{
			connections: []Point{pointA, pointB},
		})
		return circuits, edge
	}

	// check if pointA and pointB already are in circuits
	aInCircuit, aIndex := false, 0
	bInCircuit, bIndex := false, 0
	for i, circuit := range circuits {
		if slices.Contains(circuit.connections, pointA) {
			aInCircuit, aIndex = true, i
		}
		if slices.Contains(circuit.connections, pointB) {
			bInCircuit, bIndex = true, i
		}
	}

	// both exists in the same circuit
	if aInCircuit && bInCircuit && aIndex == bIndex {
		return circuits, edge
	}

	// two connections need to be merged into one
	if aInCircuit && bInCircuit && aIndex != bIndex {
		circuits[aIndex].connections = slices.Concat(circuits[aIndex].connections, circuits[bIndex].connections)
		circuits = slices.Delete(circuits, bIndex, bIndex + 1)
		return circuits, edge
	}

	// if only a exists
	if aInCircuit {
		circuits[aIndex].connections = append(circuits[aIndex].connections, pointB)
		return circuits, edge
	}
	// if only b exists
	if bInCircuit {
		circuits[bIndex].connections = append(circuits[bIndex].connections, pointA)
		return circuits, edge
	}

	// neither exists in any circuit
	circuits = append(circuits, Circuit{
		connections: []Point{pointA, pointB},
	})
	return circuits, edge
}