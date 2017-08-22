// +build deployed,!deploy

package engine

func init() {
	// log.Println("engine_play.go - deploy = false | play = true")
	play = true
	deploy = false
}
