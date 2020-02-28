package main

import (
	"log"

	"github.com/hibooboo2/hopwatch"
)

func main() {
	for i := 0; i < 100; i++ {
		line()
		log.Printf("Passing break %d", i)
		hopwatch.Break()
	}
}

func line() {
	hopwatch.Break()
	log.Println("In line")
	hopwatch.Break()
	hopwatch.Printf("Layers are objects on the map that consist of one or more separate items, but are manipulated as a single unit. Layers generally reflect collections of objects that you add on top of the map to designate a common association.")
	hopwatch.Break()

}
