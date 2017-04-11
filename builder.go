// !build builder

package main

import (
	"./engine"
	"log"
	"os"
	"os/exec"
)

func main() {
	var (
		out []byte
		err error
	)
	log.Println("builder")
	log.Println(engine.GetTags())
	if len(engine.GetTags()) > 0 {
		if out, err = exec.Command("go", "build", "-tags", engine.GetTags(), "-v", "-o", "game").Output(); err != nil {
			log.Println("Failed:", err)
			os.Exit(1)
		} else {
			log.Println("Success:", string(out))
		}
	} else {
		log.Println("Failed:", "Empty game")
	}
}
