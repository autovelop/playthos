package audio

import (
	"fmt"
	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPackage(&engine.Package{"audio", []string{}, []string{"audio"}, []string{"generic"}})
	fmt.Println("> Audio: Initializing")
}
