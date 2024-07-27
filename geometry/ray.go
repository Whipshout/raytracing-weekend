package geometry

import "math"

type Ray struct {
	orig Point3
	dir  Vec3
}

func NewRay(origin Point3, direction Vec3) Ray {
	return Ray{
		orig: origin,
		dir:  direction,
	}
}

func (r Ray) Origin() Point3 {
	return r.orig
}

func (r Ray) Direction() Vec3 {
	return r.dir
}

func (r Ray) At(t float64) Point3 {
	return r.orig.Add(r.dir.Mul(t))
}

func RayColor(r Ray) Color {
	t := HitSphere(NewVec3(0.0, 0.0, -1.0), 0.5, r)

	if t > 0.0 {
		n := UnitVector(SubVec3(r.At(t), NewVec3(0.0, 0.0, -1.0)))
		return MulScalarVec3(0.5, NewVec3(n.X()+1, n.Y()+1, n.Z()+1))
	}

	unitDirection := UnitVector(r.Direction())
	a := 0.5 * (unitDirection.Y() + 1.0)

	firstPartAddition := MulScalarVec3(1.0-a, NewVec3(1.0, 1.0, 1.0))
	secondPartAddition := MulScalarVec3(a, NewVec3(0.5, 0.7, 1.0))

	return AddVec3(firstPartAddition, secondPartAddition)
}

func HitSphere(center Point3, radius float64, r Ray) float64 {
	oc := SubVec3(center, r.Origin())
	a := r.Direction()
	aa := a.LengthSquared()
	h := Dot(r.Direction(), oc)
	c := oc.LengthSquared() - radius*radius
	discriminant := h*h - aa*c

	if discriminant < 0 {
		return -1.0
	} else {
		return (h - math.Sqrt(discriminant)) / aa
	}
}
