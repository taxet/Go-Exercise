package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var res = make([][]uint8, dy)
	for i := range res {
		res[i] = make([]uint8, dx)
		for j := range res[i] {
			//res[i][j] = uint8(i * j)
			res[i][j] = uint8(dx ^ dy + (dx+dy)/2)
		}
	}
	return res
}
func main() {
	pic.Show(Pic)
}
