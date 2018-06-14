@ECHO OFF
SET GOCACHE=off
:: SET GODEBUG=gctrace=1,cgocheck=1
:: go test -v  headless/headless_test.go || GOTO :error
:: go test -v -tags="render glfw" glfw/glfw_test.go || GOTO :error
:: go test -v -tags="render glfw opengl" opengl/opengl_test.go
:: go test -v -tags="render glfw opengl" opengl/opengl_square_test.go 
go test -v -tags="render glfw opengl windows" opengl/opengl_texture_test.go
goto :success

:error
PAUSE
GOTO :EOF

:success
PAUSE
GOTO :EOF
