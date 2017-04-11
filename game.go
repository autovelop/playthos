// +build !builder

package main

import (
	"./engine"
	"log"
)

func main() {
	log.Println("game")
	log.Println(engine.GetTags())
}
