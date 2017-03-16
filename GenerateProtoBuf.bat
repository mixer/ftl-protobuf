
@echo off
echo Setting up...
set gopath = %USERPROFILE%\go
set gobin = %gopath%\bin
set path = %path%;%gopath%;%gobin%;
set old_goos = %goos%

echo Ensuring protoc-gen-go is installed...
go get -u github.com/golang/protobuf/protoc-gen-go

echo Generating...
protobuf-3.2.0\bin\protoc.exe -I="ftl\protocols" --cpp_out=ftl\generated-cpp --go_out=ftl\generated-go ftl\protocols\handshake.proto

set goos = %goos%

echo Done!