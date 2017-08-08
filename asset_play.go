// +build !deployed

package engine

import (
	"go/build"
	"io/ioutil"
	// "log"
	"os"
)

func LoadAsset(d string, f string) ([]byte, error) {
	dir, err := build.ImportDir(d, build.FindOnly)
	if err != nil {
		return nil, err
	}
	err = os.Chdir(dir.Dir)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// go back to root dir
	err = os.Chdir("../")
	if err != nil {
		return nil, err
	}

	return buf, nil
}
