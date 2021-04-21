@echo off

SET GOROOT=C:\Users\gmaga\sdk\go1.16.3
SET GOPATH=C:\Users\gmaga\go

mkdir build 2> NUL

:: *** Build Windows ***
echo * Building Windows executable
mkdir build\windows 2> NUL
SET GOOS=windows
SET GOARCH=amd64
del build\windows\filecleanup.exe
%GOROOT%\bin\go.exe build -o build\windows\filecleanup.exe src\filecleanup.go

:: *** Build Linux ***
echo * Building Linux executable
mkdir build\linux 2> NUL
SET GOOS=linux
SET GOARCH=amd64
del build\linux\filecleanup
%GOROOT%\bin\go.exe build -o build\linux\filecleanup src\filecleanup.go
