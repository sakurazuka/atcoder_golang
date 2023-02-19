package main

import (
	"atcoder_golang/beginners"
	"atcoder_golang/sample"
	"os"
)

func main() {
	switch os.Args[1] {
	case "sample":
		switch os.Args[2] {
		case "HelloWorld":
			sample.HelloWorld("hoge")
		}
	case "beginners":
		switch os.Args[2] {
		case "WelcomeToAtcoder":
			beginners.WelcomeToAtcoder(os.Stdin)
		case "Product":
			beginners.Product(os.Stdin)
		case "PlacingMarbles":
			beginners.PlacingMarbles(os.Stdin)
		case "ShiftOnly":
			beginners.ShiftOnly(os.Stdin)
		}
	}
}
