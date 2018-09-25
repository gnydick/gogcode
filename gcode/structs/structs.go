package structs

type MotionPositioning int

const (
	relativeMotion MotionPositioning = 91
	absoluteMotion MotionPositioning = 90
)

type ExtruderPositioning int

const (
	relativeExtruding ExtruderPositioning = 83
	absoluteExtruding ExtruderPositioning = 82
)

type ToolSet struct {
	tools  map[int]Tool
	active int
}

const (
	off     int = 0
	active  int = 1
	standby int = 2
)

type Position struct {
	X float32
	Y float32
	Z float32
}

type Tool struct {
	Number      int
	ActiveTemp  int
	StandbyTemp int
	Positioning ExtruderPositioning
	Position    Position
}

type GcodeState struct {
	ToolSet           ToolSet
	HasX              bool
	HasY              bool
	HasZ              bool
	MotionPositioning MotionPositioning
	OtherParams       map[byte]bool
}
