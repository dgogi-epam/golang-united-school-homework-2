package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum int) float64 {
	const sideTriangle = 3
	const sideSquare = 4
	const sideCircle = 0
	switch sidesNum {
	case sideCircle:
		return math.Pi * math.Pow(sideLen, 2)
	case sideTriangle:
		return math.Pow(sideLen, 2) * math.Sin(math.Pi/3) / 2
	case sideSquare:
		return math.Pow(sideLen, 2)
	default:
		return 0
	}
}
