// +build deploy linux

package linux

import (
	"fmt"
	"github.com/autovelop/playthos"
	"go/build"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func init() {
	engine.NewIntegrant(&Linux{})
	fmt.Println("> Linux: Ready")
}

type Linux struct {
	engine.Integrant
	assets   map[string][]byte
	isDeploy bool
}

func (l *Linux) InitIntegrant() {
	l.assets = make(map[string][]byte, 0)
}

func (l *Linux) AddIntegrant(engine.IntegrantRoutine) {}

func (l *Linux) Destroy() {}

func (l *Linux) IsDeploy() {
	l.isDeploy = true
}

func (l *Linux) Asset(p string) []byte {
	if l.isDeploy {
		return nil
	}
	return l.assets[p]
}

func (l *Linux) LoadAsset(p string) {
	if l.isDeploy {
		return
	}
	wd, err := os.Getwd()
	if err != nil {
		log.Println("> Engine: Unable to get working directory for Linux platform")
	}

	splits := strings.Split(p, "/")

	d := splits[0]
	f := d
	if d != p {
		if len(splits) > 1 {
			f = splits[1]
		}

		dir, err := build.ImportDir(d, build.FindOnly)
		if err != nil {
			log.Println("> Engine: Invalid path to load asset for Linux platform")
			log.Println("          PLATFORM: linux")
			log.Printf("          DIRECTORY: %v\n", d)
			log.Printf("          PATH: %v\n", p)
			log.Fatalf("          CWD: %v", wd)
		}

		err = os.Chdir(dir.Dir)
		if err != nil {
			wd, _ := os.Getwd()
			log.Fatalf("unable to navigate to destination folder %v from %v", dir.Dir, wd)
		}
	}

	file, err := os.Open(f)
	if err != nil {
		log.Println("> Engine: Unable to open asset file for Linux platform. Could be in use.")
		log.Println("          PLATFORM: linux")
		log.Printf("          PATH: %v\n", p)
		log.Printf("          CWD: %v\n", wd)
		log.Fatalf("          Error: %v", err)
	}

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("unable to read destination file or folder: %v", p)
	}
	file.Close()

	// l.Decode(buf)

	// go back to root dir
	err = os.Chdir(wd)
	if err != nil {
		log.Fatalf("unable to navigate to parent from destination folder", err)
	}

	l.assets[p] = buf
	// splits := strings.Split(p, "/")
	// d := splits[0]
	// f := splits[1]

	// dir, err := build.ImportDir(d, build.FindOnly)
	// if err != nil {
	// 	log.Fatalf("destination file or folder doesn't exist: %v", p)
	// 	return
	// }

	// err = os.Chdir(dir.Dir)
	// if err != nil {
	// 	wd, _ := os.Getwd()
	// 	log.Fatalf("unable to navigate to destination folder %v from %v", dir.Dir, wd)
	// 	return
	// }

	// file, err := os.Open(f)
	// if err != nil {
	// 	log.Fatalf("unable to open destination file: %v", f)
	// 	return
	// }
	// defer file.Close()

	// buf, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	log.Fatalf("unable to read destination file or folder: %v", p)
	// 	return
	// }

	// // l.Decode(buf)

	// // go back to root dir
	// err = os.Chdir("../")
	// if err != nil {
	// 	return
	// }

	// l.assets[p] = buf
}
