/** mocking.go **/

package main

import "fmt"

type (
	// Values defines interface to be used in mocking
	Values interface {
		GetVolume() int
		GetSurfaceArea() int
	}

	// Measurements defines the measurements used to calculate
	// the surface area and volume of a cube or a cuboid.
	Measurements struct {
		Length int
		Width  int
		Height int
	}
)

// GetVolume calculates and returns the volume of a cube or a cuboid
func (config *Measurements) GetVolume() int {
	return config.Height * config.Length * config.Width
}

// GetSurfaceArea calculates and returns the surface area of a cube or cuboid
func (config *Measurements) GetSurfaceArea() int {
	return (config.Length * config.Width * 2) +
		(config.Length * config.Width * 2) +
		(config.Height * config.Width * 2)
}

// GetVolumeAndArea fetches and returns the volume and the area
func GetVolumeAndArea(val Values) (int, int) {
	return val.GetVolume(), val.GetSurfaceArea()
}

// program execution starts here
func main() {
	var (
		area, volume int

		data = &Measurements{
			Length: 1,
			Height: 5,
			Width:  3,
		}
	)

	volume, area = GetVolumeAndArea(data)

	fmt.Printf("Volume : %d , Area : %d \n", volume, area)
}
