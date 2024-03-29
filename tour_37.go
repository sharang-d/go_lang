package main

import "code.google.com/p/go-tour/pic"


func Pic(dx, dy int) [][]uint8 {
    result := make([][]uint8, dy)
    for i := range result {
        result[i] = make([]uint8, dx)
        for j := range result[i]{
			result[i][j] = uint8(i * j)
        }
    }
    return result
}

func main() {
    pic.Show(Pic)
}
