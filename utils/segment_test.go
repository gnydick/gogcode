package utils

import (
	"fmt"
	"testing"
)

func TestCalcWidth(*testing.T) {
	hotend := NewHotEnd(1.75, .4)
	segment := Segment{
		SegmentLength:      10,
		FeedRate:           3600,
		FilamentFeedLength: 0.36343319022096000000,
		LayerHeight:        .2,
	}

	ext := NewExtruder(hotend)
	fmt.Println(ext.CrossSection(segment, 2))
	fmt.Println(ext.SegmentWidth(segment, 2))
}
