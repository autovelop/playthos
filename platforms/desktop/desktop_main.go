// +build deploy desktop

package desktop

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
	engine.NewIntegrant(&Desktop{})
	fmt.Println("> Desktop: Ready")
}

type Desktop struct {
	engine.Integrant
	assets   map[string][]byte
	isDeploy bool
}

func (l *Desktop) InitIntegrant() {
	l.assets = make(map[string][]byte, 0)
}

func (l *Desktop) AddIntegrant(engine.IntegrantRoutine) {}

func (l *Desktop) Destroy() {}

func (l *Desktop) IsDeploy() {
	l.isDeploy = true
}

func (l *Desktop) Asset(p string) []byte {
	if l.isDeploy {
		return nil
	}
	return l.assets[p]
}

func (l *Desktop) LoadAsset(p string) {
	if l.isDeploy {
		return
	}
	wd, err := os.Getwd()
	if err != nil {
		log.Println("> Engine: Unable to get working directory for Desktop platform")
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
			log.Println("> Engine: Invalid path to load asset for Desktop platform")
			log.Println("          PLATFORM: desktop")
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
		log.Println("> Engine: Unable to open asset file for Desktop platform. Could be in use.")
		log.Println("          PLATFORM: desktop")
		log.Printf("          PATH: %v\n", p)
		log.Printf("          CWD: %v\n", wd)
		log.Fatalf("          Error: %v", err)
	}

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("unable to read destination file or folder: %v", p)
	}
	file.Close()

	err = os.Chdir(wd)
	if err != nil {
		log.Fatalf("unable to navigate to parent from destination folder", err)
	}

	l.assets[p] = buf
}
