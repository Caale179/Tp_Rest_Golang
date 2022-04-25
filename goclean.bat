@echo off
@echo on
go clean -v ./...
@echo off
if %ERRORLEVEL% GEQ 1 echo !!!!! ERROR !!!!!