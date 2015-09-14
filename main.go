package main

import (
	"fmt"
	"os"
	"image"
	"github.com/rogpeppe/misc/svg"
)

func example() (val string, err error){
	val = "test string";

	return
}

func main() {

	file, err := os.Open("img/test.svg");

	size := image.Point{300,300};

	if err != nil {
		panic(err);
	}

	fmt.Println(svg.Render(file, size));
}
