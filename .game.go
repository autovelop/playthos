// +build !builder

package main

import (
	"github.com/autovelop/playthos/engine"
	"log"
)

func main() {
	log.Println("game")
	log.Println(engine.GetTags())
}
