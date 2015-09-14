package main

import "fmt"
import "os"
import "github.com/rogpeppe/misc/svg"
import "image"

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
