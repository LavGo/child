package distance

import (
	"testing"
	"fmt"
)

var start = &Position{
	Latitude:  22.52931,
	Longitude: 113.95305,
}

func TestPositionConvert(t *testing.T) {

	end := PositionConvert(start, 100, 0)
	fmt.Println(end)
}

func TestTargetArea(t *testing.T) {
	area:=TargetArea(start,100)
	fmt.Println(area)
}
