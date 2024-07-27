package main

import (
	"fmt"
	"log"
	"os"
	"raytracing_weekend/geometry"
)

func main() {
	logger := log.New(os.Stderr, "", 0)

	// Image
	aspectRatio := 16.0 / 9.0
	imageWidth := 400

	// Calculate the image height, and ensure that it's at least 1
	imageHeight := int(float64(imageWidth) / aspectRatio)
	if imageHeight < 1 {
		imageHeight = 1
	}

	// Camera
	focalLength := 1.0
	viewportHeight := 2.0
	viewportWidth := viewportHeight * float64(imageWidth/imageHeight)
	cameraCenter := geometry.NewVec3()

	// Calculate the vectors across the horizontal and down the vertical viewport edges
	viewportU := geometry.NewVec3(viewportWidth, 0, 0)
	viewportV := geometry.NewVec3(0, -viewportHeight, 0)

	// Calculate the horizontal and vertical delta vectors from pixel to pixel
	pixelDeltaU := geometry.DivScalarVec3(viewportU, float64(imageWidth))
	pixelDeltaV := geometry.DivScalarVec3(viewportV, float64(imageHeight))

	// Calculate the location of the upper left pixel
	viewportUpperLeft := geometry.SubVec3(cameraCenter, geometry.NewVec3(0, 0, focalLength))
	viewportUpperLeft = geometry.SubVec3(viewportUpperLeft, geometry.DivScalarVec3(viewportU, 2.0))
	viewportUpperLeft = geometry.SubVec3(viewportUpperLeft, geometry.DivScalarVec3(viewportV, 2.0))
	pixel00Loc := geometry.AddVec3(viewportUpperLeft, geometry.MulScalarVec3(0.5, geometry.AddVec3(pixelDeltaU, pixelDeltaV)))

	// Render
	fmt.Printf("P3\n%v %v\n255\n", imageWidth, imageHeight)

	for j := 0; j < imageHeight; j++ {
		logger.Printf("Scanlines remaining: %v \n", imageHeight-j)
		for i := 0; i < imageWidth; i++ {
			deltaU := geometry.MulScalarVec3(float64(i), pixelDeltaU)
			deltaV := geometry.MulScalarVec3(float64(j), pixelDeltaV)

			pixelCenter := geometry.AddVec3(pixel00Loc, deltaU)
			pixelCenter = geometry.AddVec3(pixelCenter, deltaV)

			rayDirection := geometry.SubVec3(pixelCenter, cameraCenter)

			ray := geometry.NewRay(cameraCenter, rayDirection)

			pixelColor := geometry.RayColor(ray)

			geometry.WriteColor(pixelColor)
		}
	}
}
