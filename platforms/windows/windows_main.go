// +build deploy win

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
	splits := strings.Split(p, "/")
	d := splits[0]
	f := splits[1]

	dir, err := build.ImportDir(d, build.FindOnly)
	if err != nil {
		log.Fatalf("destination file or folder doesn't exist: %v", p)
		return
	}

	err = os.Chdir(dir.Dir)
	if err != nil {
		wd, _ := os.Getwd()
		log.Fatalf("unable to navigate to destination folder %v from %v", dir.Dir, wd)
		return
	}

	file, err := os.Open(f)
	if err != nil {
		log.Fatalf("unable to open detination file: %v", f)
		return
	}
	defer file.Close()

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("unable to read detination file or folder: %v", p)
		return
	}

	// l.Decode(buf)

	// go back to root dir
	err = os.Chdir("../")
	if err != nil {
		return
	}

	l.assets[p] = buf
}
