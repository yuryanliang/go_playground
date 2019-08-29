/** mocking_test.go **/

package main

import (
	"strconv"
	"testing"
)

// MockTest helps implement our customized GetVolume and GetSurfaceArea
// needed to database the original implementation in mocking.go
type MockTest struct {
	Elem int
}

// GetVolume returns the volume of a cube
func (config *MockTest) GetVolume() int {
	return config.Elem * config.Elem * config.Elem
}

// GetSurfaceArea returns the surface area of a cube
func (config *MockTest) GetSurfaceArea() int {
	return config.Elem * config.Elem * 6
}

// TestGetVolumeAndArea tests the functionality of GetVolumeAndArea
func TestGetVolumeAndArea(t *testing.T) {
	for key, val := range map[MockTest][]int{
		MockTest{Elem: 5}: []int{125, 150},
		MockTest{Elem: 2}: []int{8, 24},
		MockTest{Elem: 1}: []int{1, 6},
	} {
		t.Run("Cube side length "+strconv.Itoa(key.Elem), func(t *testing.T) {
			volume, area := GetVolumeAndArea(&key)

			if volume != val[0] {
				t.Errorf("Expected volume to be equal to %d but was equal to %d ",
					volume, val[0])
				t.FailNow()
			}

			if area != val[1] {
				t.Errorf("Expected area to be equal to %d but was equal to %d ",
					area, val[1])
				t.FailNow()
			}
		})
	}
}
