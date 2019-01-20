package main

import (
	"fmt"
	"github.com/gnydick/gogcode/gcode/utils"
)

func main() {
	hotend := utils.NewHotEnd(1.75, .4)
	segment := utils.Segment{
		SegmentLength:      10,
		FeedRate:           3600,
		FilamentFeedLength: 0.36343319022096000000,
		LayerHeight:        .2,
	}

	ext := utils.NewExtruder(hotend)
	fmt.Println(ext.CrossSection(segment, 2))
	fmt.Println(ext.SegmentWidth(segment, 2))
}
