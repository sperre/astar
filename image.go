package astar

import (
	"os"
	"image"
	//"fmt"
)

import _ "image/png"

func openImage(filename string) (image.Image) {
	f, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer f.Close()
	img, _, _ := image.Decode(f)
	return img
}

func parseImage(img image.Image) MapData {
	max := uint32(65536-1) // 2^16-1

	bounds := img.Bounds()
	map_data := NewMapData(bounds.Max.X, bounds.Max.Y)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()

			if(r == max && g == max && b == max && a == max) {
				map_data[x][bounds.Max.Y-1-y] = LAND
				//fmt.Printf(".")
			} else {
				map_data[x][bounds.Max.Y-1-y] = WALL
				//fmt.Printf("#")
			}
		}
		//fmt.Println("")
	}
	return map_data
}

func GetMapFromImage(filename string) MapData {
	img := openImage(filename)
	if(img == nil) {
		return nil
	}
	return parseImage(img)
}
