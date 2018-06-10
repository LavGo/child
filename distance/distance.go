package distance

import (
	"math"
)

// 赤道半径
const Ea float64 = 6378137

// 极半径
const Eb float64 = 6356725

type Position struct {
	Longitude float64 // 经度
	Latitude  float64 // 纬度
}

type RectArea struct {
	MinLongitude float64
	MaxLongitude float64
	MinLatitude  float64
	MaxLatitude  float64
}

/**目标点
start ： 起始点
distance ： 距离
angle ： 弧度
 */
func PositionConvert(start *Position, distance float64, angle float64) *Position {

	dx := distance * 1000 * math.Sin((angle*math.Pi)/180.0)
	dy := distance * 1000 * math.Cos((angle*math.Pi)/180.0)

	ec := Eb + (Ea-Eb)*(90-start.Latitude)/90.0
	ed := ec * math.Cos((start.Latitude*math.Pi)/180.0)

	return &Position{
		Latitude:  (dy/ec + start.Latitude*math.Pi/180.0) * 180.0 / math.Pi,
		Longitude: (dx/ed + start.Longitude*math.Pi/180.0) * 180.0 / math.Pi,
	}
}

/**范围矩形
center ： 中心点
distance ：范围
 */
func TargetArea(center *Position, distance float64) *RectArea {
	up := PositionConvert(center, distance, 0.0)
	down := PositionConvert(center, distance, 180.0)

	right := PositionConvert(center, distance, 90.0)
	left := PositionConvert(center, distance, 270.0)
	return &RectArea{
		MaxLatitude:  up.Latitude,
		MinLatitude:  down.Latitude,
		MaxLongitude: right.Longitude,
		MinLongitude: left.Longitude,
	}
}
