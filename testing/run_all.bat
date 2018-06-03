SET GOCACHE=off
REM SET GODEBUG=gctrace=1,cgocheck=1
go test -v -tags="render glfw" glfw/glfw_test.go
go test -v -tags="render glfw opengl" opengl/opengl_test.go
