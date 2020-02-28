package main

import (
	"log"

	"github.com/emicklei/hopwatch"
)

func main() {
	for i := 0; i < 100; i++ {
		line()
		log.Println("Passing break")
		hopwatch.Break()
		log.Println("Passed break")
	}
}

func line() {
	log.Println("In line")
	hopwatch.Printf("Layers are objects on the map that consist of one or more separate items, but are manipulated as a single unit. Layers generally reflect collections of objects that you add on top of the map to designate a common association.")
}
