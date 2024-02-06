if (-not(Test-Path -Path ".\build" -PathType Container))
{
    New-Item -ItemType Directory -Path ".\build"
}

# Build for Windows 64-bit
$env:GOOS = "windows"
$env:GOARCH = "amd64"
go build -o ./build/qrCodeGenerator-win64.exe main.go

# Build for macOS (Apple Silicon)
$env:GOOS = "darwin"
$env:GOARCH = "arm64"
go build -o ./build/qrCodeGenerator-macsilicon.exe main.go

# Build for macOS (Intel)
$env:GOOS = "darwin"
$env:GOARCH = "amd64"
go build -o ./build/qrCodeGenerator-macintel.exe main.go

# Build for Linux (amd64)
$env:GOOS = "linux"
$env:GOARCH = "amd64"
go build -o ./build/qrCodeGenerator-amd64.exe main.go

# Build for Linux (ARM architecture, e.g., Raspberry Pi)
$env:GOOS = "linux"
$env:GOARCH = "arm"
go build -o ./build/qrCodeGenerator-arm.exe main.go
