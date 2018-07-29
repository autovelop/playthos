// +build !deploy

package engine

import (
	"fmt"
	"os"
)

func initDeploy(n string) {
	fmt.Printf("> Engine: Unable to deploy with this executable\n")
	os.Exit(0)
}
