// +build deploy windows

package windows

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
	engine.NewIntegrant(&Windows{})
	fmt.Println("> Windows: Ready")
}

type Windows struct {
	engine.Integrant
	assets map[string][]byte
}

func (l *Windows) InitIntegrant() {
	l.assets = make(map[string][]byte, 0)
}

func (l *Windows) AddIntegrant(engine.IntegrantRoutine) {}

func (l *Windows) Destroy() {}

func (l *Windows) Asset(p string) []byte {
	return l.assets[p]
}

func (l *Windows) LoadAsset(p string) {
	wd, err := os.Getwd()
	if err != nil {
		log.Println("> Engine: Unable to get working directory for Windows platform")
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
			log.Println("> Engine: Invalid path to load asset for Windows platform")
			log.Println("          PLATFORM: windows")
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
		log.Println("> Engine: Unable to open asset file for Windows platform. Could be in use.")
		log.Println("          PLATFORM: windows")
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
}
