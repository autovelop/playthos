#!/bin/bash
export GOCACHE=off
# go test -v  headless/headless_test.go
# go test -v -tags="render glfw" glfw/glfw_test.go
# go test -v -tags="render glfw opengl web deploy" opengl/opengl_texture_test.go
go run -v -tags="render glfw opengl deploy" opengl/opengl_texture.go
# go test -v -tags="render glfw opengl" opengl/opengl_test.go
