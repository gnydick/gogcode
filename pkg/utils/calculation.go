package utils

import "math"

// May be three dimensional, like in vase mode
type Segment struct {
	SegmentLength      float64
	FeedRate           float64
	FilamentFeedLength float64
	LayerHeight        float64
}

type HotEnd struct {
	filamentDiameter     float64
	nozzleDiameter       float64
	feedRatio            float64
	filamentCrossSection float64
	nozzleCrossSection   float64
}

func NewHotEnd(filamentDiameter float64, nozzleDiameter float64) HotEnd {
	return HotEnd{
		filamentDiameter:     filamentDiameter,
		nozzleDiameter:       nozzleDiameter,
		filamentCrossSection: math.Pow(filamentDiameter/2, 2) * math.Pi,
		nozzleCrossSection:   math.Pow(nozzleDiameter/2, 2) * math.Pi,
		feedRatio:            math.Pow(filamentDiameter/2, 2) / math.Pow(nozzleDiameter/2, 2),
	}
}

func NewExtruder(hotEnd HotEnd) Extruder {
	return Extruder{
		hotEnd: hotEnd,
	}
}

func (h HotEnd) FeedRatio() float64 {
	return h.feedRatio
}

type Extruder struct {
	hotEnd HotEnd
}

func (e Extruder) FullPrecisionSegmentVolume(segment Segment) (segmentVolume float64) {
	return e.hotEnd.filamentCrossSection * segment.FilamentFeedLength
}

func (e Extruder) FullPrecisionCrossSection(segment Segment) (crossSection float64) {
	return e.FullPrecisionSegmentVolume(segment) / segment.SegmentLength
}

func (e Extruder) FullPrecisionSegmentWidth(segment Segment) (extrusionWidth float64) {
	return (e.FullPrecisionCrossSection(segment)-math.Pi*math.Pow(segment.LayerHeight/2, 2))/segment.LayerHeight + segment.LayerHeight

}

func (e Extruder) SegmentVolume(segment Segment, decimals float64) (segmentVolume float64) {
	return round(e.FullPrecisionSegmentVolume(segment), decimals)
}

func (e Extruder) CrossSection(segment Segment, decimals float64) (crossSection float64) {
	return round(e.FullPrecisionCrossSection(segment), decimals)
}

func (e Extruder) SegmentWidth(segment Segment, decimals float64) (extrusionWidth float64) {
	return round(e.FullPrecisionSegmentWidth(segment), decimals)

}

func round(float float64, decimals float64) float64 {
	return math.Round(float*math.Pow(10, decimals)) / math.Pow(10, decimals)
}
