package main

import (
	"flag"

	"github.com/gnydick/gogcode/utils"
)

func main() {
	//var state = State{}
	input := flag.String("input", "", "input file")
	output := flag.String("output", "", "output file")

	xgcode := flag.String("xgcode", "M572 S", "what command to use along the x-axis")
	//xRange := flag.String("xRange", "0:100", "starting and ending value `0:100`")
	//ygcode := flag.String("ygcode", "M566 E", "what command to use along the y-axis")
	//yRange := flag.String("yRange", "0:100", "starting and ending value `0:100`")
	//zgcode := flag.String("zgcode", "M204 P", "what command to use along the z-axis")
	//zRange := flag.String("zRange", "0:100", "starting and ending value `0:100`")
	flag.Parse()
	println(*input)
	println(*output)
	println(*xgcode)

	parsingChannel := make(chan string)

	go utils.Scan(&parsingChannel, input)

	for line := range parsingChannel {

	}

}
