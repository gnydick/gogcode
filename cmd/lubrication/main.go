package main

import (
	flag "flag"
	"fmt"
	. "github.com/gnydick/gogcode/pkg/utils"
)

func main() {

	speed := flag.Float64("speed", 200, "mm/sec desired (not actual)")
	axis := flag.String("axis", "x", "which axis to lube")
	start := flag.Float64("start", 20, "start position on axis to lube")
	end := flag.Float64("end", 20, "end position on axis to lube")
	strokes := flag.Int("strokes", 100, "number of strokes before proceeding to next position")
	intervals := flag.Float64("intervals", 10, "how many intervals at which to lube")
	strokeLen := flag.Float64("strokeLen", .6, "length of stroke")
	flag.Parse()

	totalDistance := *end - *start
	interval := totalDistance / *intervals

	for i := 0.0; i < *intervals; i++ {
		pos := *start + i*interval
		tfs := make([]Transform, 1)
		for j := 1; j <= *strokes; j++ {
			ax := getAxis(axis)
			tfs[0] = NewTransform(ax, pos)
			fmt.Println(Move(tfs, *speed, fmt.Sprintf("stroke %d", j)))
			tfs[0] = NewTransform(ax, pos+*strokeLen)
			fmt.Println(Move(tfs, *speed))
			tfs[0] = NewTransform(ax, pos)
			fmt.Println(Move(tfs, *speed))
			fmt.Println()
		}
	}
}

func getAxis(axis *string) Axis {
	switch *axis {
	case "X":
	case "x":
	default:
		return X
	case "Y":
	case "y":
		return Y
	case "Z":
	case "z":
		return Z
	case "E0":
	case "e0":
	case "e":
		return E0
	case "E1":
	case "e1":
		return E1
	}
	return X
}
