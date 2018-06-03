SET GOCACHE=off
REM SET GODEBUG=gctrace=1,cgocheck=1
REM go test -v -tags="render glfw" glfw_test.go
go test -v -tags="render glfw opengl" glfw_opengl_test.go
REM go tool pprof --text cpuprofile
REM go tool pprof --text memprofile

