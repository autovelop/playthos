#!/bin/bash

# Build Linux
go build -v -o ../builds/Linux/demo desktop/desktop.go

# Build Android
# gomobile build -target=android -o=../builds/Android/demo.apk mobile

# Install Android
# gomobile install -target=android -o=../builds/Android/demo.apk mobile
# adb install -rg ../builds/Android/demo.apk
# adb logcat -c
# adb logcat -s "GoLog"

#go build -v -o ../builds/Linux/ desktop/desktop.go
#gomobile build -target=android prototype
#gomobile build -target=ios prototype
