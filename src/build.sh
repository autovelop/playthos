#!/bin/bash

# Build Linux
go build -v -o ../builds/Linux/demo desktop/desktop.go

# Build Android
# gomobile build -target=android -o=../builds/Android/demo.apk mobile

# Install Android
# adb install -rg ../builds/Android/demo.apk
# adb logcat -c
# adb logcat -s "GoLog"

# gomobile install -target=android -o=../builds/Android/demo.apk mobile

#go build -v -o ../builds/Linux/ desktop/desktop.go
#gomobile build -target=android prototype
#gomobile build -target=ios prototype
