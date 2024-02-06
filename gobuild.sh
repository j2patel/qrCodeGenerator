#!/bin/bash

mkdir -p ./build

# Build for Windows 64-bit
GOOS=windows GOARCH=amd64 go build -o ./build/qrCodeGenerator-win64.exe main.go

# Build for macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o ./build/qrCodeGenerator-macsilicon.exe main.go

# Build for macOS (Intel)
GOOS=darwin GOARCH=amd64 go build -o ./build/qrCodeGenerator-macintel main.go

# Build for Linux (amd64)
GOOS=linux GOARCH=amd64 go build -o ./build/qrCodeGenerator-amd64 main.go

# Build for Linux (ARM architecture, e.g., Raspberry Pi)
GOOS=linux GOARCH=arm go build -o ./build/qrCodeGenerator-arm main.go
