package geometry

import (
	"fmt"
)

type Color = Vec3

func WriteColor(c Color) {
	r := c.X()
	g := c.Y()
	b := c.Z()

	rByte := int(255.999 * r)
	gByte := int(255.999 * g)
	bByte := int(255.999 * b)

	fmt.Printf("%v %v %v\n", rByte, gByte, bByte)
}
