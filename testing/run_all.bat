SET GOCACHE=off
SET GODEBUG=gctrace=1
REM go test -v -tags="render glfw" glfw_test.go
go test -v -tags="render glfw opengl" glfw_opengl_test.go
go tool pprof --text cpuprofile
go tool pprof --text memprofile

