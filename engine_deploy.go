// +build deploy,!play

package engine

func init() {
	play = false
	deploy = true
}
