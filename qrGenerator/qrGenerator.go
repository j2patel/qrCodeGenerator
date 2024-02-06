package qrGenerator

import (
	"fmt"
	"github.com/yeqown/go-qrcode/v2"
)

func GenerateQRCode(text string) (*qrcode.QRCode, error) {
	qrCode, err := qrcode.New(text)
	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
		return nil, nil
	}

	return qrCode, nil
}
