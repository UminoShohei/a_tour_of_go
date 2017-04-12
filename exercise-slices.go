package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dy)
	for i := range img {
		img[i] = make([]uint8, dx)
	}
	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			img[x][y] = uint8(x) * uint8(y)
		}
	}
	return img
}

func main() {
	pic.Show(Pic)
}
