/*
Package audio adds the Audio listener integrant with the clip, sound, source, and listener components.
*/
package audio

import (
	"fmt"
	"github.com/autovelop/playthos"
)

func init() {
	engine.RegisterPackage(&engine.Package{"audio", []string{}, []string{"audio"}, []string{"generic"}})
	fmt.Println("> Audio: Initializing")
}
