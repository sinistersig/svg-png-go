package main

import (
	"fmt"
	"os"
	"image"
	"github.com/rogpeppe/misc/svg"
	"github.com/llgcode/draw2d/draw2dimg"
)

func example() (val string, err error){
	val = "test string";

	return
}

func main() {

	file, err := os.Open("img/test.svg");

	size := image.Point{1000,1000};

	if err != nil {
		panic(err);
	}


	dest, err := svg.Render(file, size);

	if err != nil {
		panic(err);
	}

	draw2dimg.SaveToPngFile("/home/test.png", dest);

	file.Close();

	fmt.Println("done");
}
