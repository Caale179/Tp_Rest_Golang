@echo off
@echo on
go clean -v ./...
go build -o . -v ./...
@echo off
if %ERRORLEVEL% GEQ 1 echo !!!!! ERROR !!!!!