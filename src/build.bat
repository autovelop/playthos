:: Build Win
go build -v -o ../builds/Windows/demo.exe desktop/desktop.go

:: Build Android
gomobile build -target=android -o=../builds/Android/demo.apk mobile

:: Debug Android (device plugged in)
:: gomobile install -target=android -o=../builds/Android/demo.apk mobile
:: adb logcat -c
:: adb logcat -s "GoLog"