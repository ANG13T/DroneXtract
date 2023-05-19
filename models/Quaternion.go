package models

import (
	"math"
)

type Quaternion struct {
	Scalar float64
	X      float64
	Y      float64
	Z      float64
}

func NewQuaternion(x, y, z, scalar float64) *Quaternion {
	return &Quaternion{
		X:      x,
		Y:      y,
		Z:      z,
		Scalar: scalar,
	}
}

func (q *Quaternion) ToEuler() []float64 {
	sqw := math.Pow(q.Scalar, 2)
	sqx := math.Pow(q.X, 2)
	sqy := math.Pow(q.Y, 2)
	sqz := math.Pow(q.Z, 2)
	yaw := 0.0
	roll := 0.0
	pitch := 0.0
	retv := make([]float64, 3)
	unit := sqx + sqy + sqz + sqw

	test := q.Scalar*q.X + q.Y*q.Z
	if unit == 0 {
		return []float64{pitch, roll, yaw}
	} else if test > 0.499*unit {
		yaw = 2.0 * math.Atan2(q.Y, q.Scalar)
		pitch = 1.570796326794897
		roll = 0.0
	} else if test < -0.499*unit {
		yaw = -2.0 * math.Atan2(q.Y, q.Scalar)
		pitch = -1.570796326794897
		roll = 0.0
	} else {
		yaw = math.Atan2(2.0*(q.Scalar*q.Z-q.X*q.Y), 1.0-2.0*(sqz+sqx))
		roll = math.Asin(2.0*test/unit)
		pitch = math.Atan2(2.0*(q.Scalar*q.Y-q.X*q.Z), 1.0-2.0*(sqy+sqx))
	}

	retv[0] = pitch
	retv[1] = roll
	retv[2] = yaw
	return retv
}

func (q *Quaternion) Conjugate() *Quaternion {
	return &Quaternion{
		X:      -q.X,
		Y:      -q.Y,
		Z:      -q.Z,
		Scalar: q.Scalar,
	}
}

func (q *Quaternion) Times(b *Quaternion) *Quaternion {
	y0 := q.Scalar*b.Scalar - q.X*b.X - q.Y*b.Y - q.Z*b.Z
	y1 := q.Scalar*b.X + q.X*b.Scalar + q.Y*b.Z - q.Z*b.Y
	y2 := q.Scalar*b.Y - q.X*b.Z + q.Y*b.Scalar + q.Z*b.X
	y3 := q.Scalar*b.Z + q.X*b.Y - q.Y*b.X + q.Z*b.Scalar
	return &Quaternion{
		X:      y1,
		Y:      y2,
		Z:      y3,
		Scalar: y0,
	}
}
