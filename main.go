package main

import "fmt"

type packageType string

var stardard packageType = "STANDARD"
var special packageType = "SPECIAL"
var rejected packageType = "REJECTED"
var invalid packageType = "INVALID"

func isHeavy(mass int) bool {
	return mass >= 20
}

func isBulky(width, height, length int) bool {
	return (width*height*length) >= 1_000_000 ||
		width >= 150 ||
		height >= 150 ||
		length >= 150
}

func sort(width, height, length, mass int) packageType {
	if width < 0 || height < 0 || length < 0 || mass < 0 {
		return invalid
	}
	if isBulky(width, height, length) && isHeavy(mass) {
		return rejected
	} else if isBulky(width, height, length) || isHeavy(mass) {
		return special
	}
	return stardard
}

func main() {
	fmt.Println("This is a package sorting program")
}
