package main

import (
	"qrCodeGenerator/window"
)

func main() {

	err := window.ShowQRCodeWindow()
	if err != nil {
		return
	}

}
