package main

import (
	"flag"
	"fmt"
	. "github.com/gnydick/gogcode/pkg/utils"
	"strings"
)

func main() {
	speed := flag.Int("s", 200, "mm/sec desired (not actual)")
	lines := flag.Int("n", 100, "number of random strokes")
	radius := flag.Float64("r", 1, "radius for movements")

	jerk := flag.String("j", "750:2000", "jerk, printing:lapping")
	flag.Parse()

	fmt.Println("G90;\nG1 X155 Y200;")

	jerkVals := strings.Split(*jerk, ":")

	fmt.Println(fmt.Sprintf("M566 X%s Y%s;", jerkVals[1], jerkVals[1]))

	for x := 1; x <= *lines; x++ {
		GenRandMove(155, 200, radius, speed)
	}

	fmt.Println(fmt.Sprintf("M566 X%s Y%s;", jerkVals[0], jerkVals[0]))
}
