package main

import (
	"fmt"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/rogpeppe/misc/svg"
	"image"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	file, err := os.Open("img/test.svg")

	check(err)

	defer file.Close()

	size := image.Point{1000, 1000}

	dest, err := svg.Render(file, size)

	check(err)

	err = draw2dimg.SaveToPngFile("/home/test.png", dest)

	check(err)

	fmt.Println("done")
}
