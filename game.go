// +build !builder

package main

import (
	"github.com/autovelop/playthos"
	"log"
)

func main() {
	log.Println("game")
	log.Println(engine.GetTags())
}
