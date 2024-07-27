package geometry

import (
	"fmt"
	"math"
)

type Vec3 struct {
	e [3]float64
}

func NewVec3(e ...float64) Vec3 {
	if len(e) == 0 {
		return Vec3{e: [3]float64{0, 0, 0}}
	}

	if len(e) == 3 {
		return Vec3{e: [3]float64{e[0], e[1], e[2]}}
	}

	panic("invalid vec3 size")
}

func (v *Vec3) X() float64 {
	return v.e[0]
}

func (v *Vec3) Y() float64 {
	return v.e[1]
}

func (v *Vec3) Z() float64 {
	return v.e[2]
}

func (v *Vec3) Neg() Vec3 {
	return Vec3{e: [3]float64{-v.e[0], -v.e[1], -v.e[2]}}
}

func (v *Vec3) At(i int) float64 {
	return v.e[i]
}

func (v *Vec3) SetAt(i int, value float64) {
	v.e[i] = value
}

func (v *Vec3) Add(u Vec3) Vec3 {
	v.e[0] += u.e[0]
	v.e[1] += u.e[1]
	v.e[2] += u.e[2]

	return *v
}

func (v *Vec3) Mul(t float64) Vec3 {
	v.e[0] *= t
	v.e[1] *= t
	v.e[2] *= t

	return *v
}

func (v *Vec3) Div(t float64) Vec3 {
	v.Mul(1 / t)

	return *v
}

func (v *Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v *Vec3) LengthSquared() float64 {
	return v.e[0]*v.e[0] + v.e[1]*v.e[1] + v.e[2]*v.e[2]
}

func (v *Vec3) String() string {
	return fmt.Sprintf("(%f, %f, %f)", v.e[0], v.e[1], v.e[2])
}

func AddVec3(u, v Vec3) Vec3 {
	return Vec3{
		e: [3]float64{
			u.e[0] + v.e[0],
			u.e[1] + v.e[1],
			u.e[2] + v.e[2],
		},
	}
}

func SubVec3(u, v Vec3) Vec3 {
	return Vec3{
		e: [3]float64{
			u.e[0] - v.e[0],
			u.e[1] - v.e[1],
			u.e[2] - v.e[2],
		},
	}
}

func MulVec3(u, v Vec3) Vec3 {
	return Vec3{
		e: [3]float64{
			u.e[0] * v.e[0],
			u.e[1] * v.e[1],
			u.e[2] * v.e[2],
		},
	}
}

func MulScalarVec3(t float64, v Vec3) Vec3 {
	return Vec3{
		e: [3]float64{
			t * v.e[0],
			t * v.e[1],
			t * v.e[2],
		},
	}
}

func MulVec3Scalar(v Vec3, t float64) Vec3 {
	return Vec3{
		e: [3]float64{
			v.e[0] * t,
			t * v.e[1] * t,
			t * v.e[2] * t,
		},
	}
}

func DivScalarVec3(v Vec3, t float64) Vec3 {
	return MulScalarVec3(1/t, v)
}

func Dot(u, v Vec3) float64 {
	return u.e[0]*v.e[0] + u.e[1]*v.e[1] + u.e[2]*v.e[2]
}

func Cross(u, v Vec3) Vec3 {
	return Vec3{
		e: [3]float64{
			u.e[1]*v.e[2] - u.e[2]*v.e[1],
			u.e[2]*v.e[0] - u.e[0]*v.e[2],
			u.e[0]*v.e[1] - u.e[1]*v.e[0],
		},
	}
}

func UnitVector(v Vec3) Vec3 {
	return DivScalarVec3(v, v.Length())
}
