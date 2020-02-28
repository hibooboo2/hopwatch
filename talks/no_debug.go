package main

import (
	"math"

	"github.com/hibooboo2/hopwatch"
)

func main() {
	liveOfPi()
}

func liveOfPi() {
	hopwatch.Dump(math.Pi).Break()
}
