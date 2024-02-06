package window

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/yeqown/go-qrcode/writer/standard"
	"io"
	"os"
	"qrCodeGenerator/qrGenerator"
)

func ShowQRCodeWindow() error {
	qrApp := app.New()
	qrWin := qrApp.NewWindow("QR Code Generator")
	qrWin.Resize(fyne.Size{Height: 1000, Width: 1280})

	inputEntry := widget.NewMultiLineEntry()
	inputEntry.SetPlaceHolder("Enter text")

	generateButton := widget.NewButtonWithIcon("Generate QR Code", theme.ContentAddIcon(), func() {
		text := inputEntry.Text
		if text != "" {
			qrCode, err := qrGenerator.GenerateQRCode(text)
			if err != nil {
				dialog.ShowError(err, qrWin)
				return
			}

			w, err := standard.New("./assets/temp/qrcode.png")
			if err != nil {
				fmt.Printf("standard.New failed: %v", err)
			}

			// save file
			if err = qrCode.Save(w); err != nil {
				fmt.Printf("could not save image: %v", err)
			}

			qrImage := canvas.NewImageFromFile("./assets/temp/qrcode.png")
			qrImage.FillMode = canvas.ImageFillOriginal

			downloadButton := widget.NewButtonWithIcon("Download", theme.DocumentSaveIcon(), func() {
				saveDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
					if err == nil && writer != nil {
						// Copy the content of the generated QR file to the selected location
						qrContent, err := os.Open("./assets/temp/qrcode.png")
						if err != nil {
							dialog.ShowError(err, qrWin)
							return
						}
						defer func(qrContent *os.File) {
							err := qrContent.Close()
							if err != nil {

							}
						}(qrContent)

						_, err = io.Copy(writer, qrContent)
						if err != nil {
							dialog.ShowError(err, qrWin)
							return
						}
						err = writer.Close()
						if err != nil {
							return
						}
					}
				}, qrWin)
				saveDialog.SetFileName("qrcode.png")
				saveDialog.Resize(fyne.Size{Height: 720, Width: 1280})
				saveDialog.Show()
			})

			qrDialogContainer := container.New(layout.NewPaddedLayout(), qrImage, layout.NewSpacer(), widget.NewLabel(text))
			qrDialogContainer.Resize(fyne.Size{Height: 720, Width: 1280})

			dialog.NewCustom("Generated QR Code", "OK", container.NewVBox(
				qrDialogContainer,
				layout.NewSpacer(),
				downloadButton,
			), qrWin).Show()

		} else {
			dialog.ShowInformation("Error", "Please enter text", qrWin)
		}
	})

	content := container.NewVBox(
		inputEntry,
		layout.NewSpacer(),
		generateButton,
	)

	qrWin.SetContent(content)
	qrWin.ShowAndRun()

	return nil
}
