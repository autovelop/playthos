// +build deploy,!play

package engine

import "log"

func init() {
	log.Println("engine_deploy.go - deploy = true | play = false")
	play = false
	deploy = true
}
