#/bin/bash

rm -f mental-math-trainer-windows.exe
GOOS=windows GOARCH=amd64 go build -o mental-math-trainer-windows.exe

rm -f mental-math-trainer-linux
GOOS=linux GOARCH=amd64 go build -o mental-math-trainer-linux

