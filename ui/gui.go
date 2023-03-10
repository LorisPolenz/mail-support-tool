package ui

import (
	"fmt"
	"image/color"
	"time"

	"dev.lopo.mail-support/helper"
	"dev.lopo.mail-support/mail"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type FormContent struct {
	Year       int
	FolderPath string
	Receiver   string
}

func RenderUI() {

	formContent := FormContent{
		FolderPath: helper.Configuration.Filesystem.Path,
		Year:       time.Now().Year(),
		Receiver:   helper.Configuration.Email.Receiver,
	}
	a := app.New()
	w := a.NewWindow("E-Mail Tool")

	folderCanvas := canvas.NewText("Ordner: "+formContent.FolderPath, color.Black)
	yearCanvas := canvas.NewText("Jahr: "+fmt.Sprint(formContent.Year), color.Black)
	receiverCanvas := canvas.NewText("Empfänger: "+formContent.Receiver, color.Black)

	dataContainer := container.New(layout.NewGridLayout(1), folderCanvas, yearCanvas, receiverCanvas)

	sendButton := widget.NewButton("Senden", func() {
		folderPath := helper.BuildFolderPath(formContent.FolderPath, formContent.Year)

		helper.CheckFilePath(folderPath)

		err := mail.Send_mail(folderPath, formContent.Receiver)

		if err != nil {
			wErr := a.NewWindow("Fehler")
			wErr.SetContent(widget.NewLabel(err.Error()))
			wErr.Show()
			return
		}

		wSucc := a.NewWindow("Erfolg")

		closeButton := widget.NewButton("Schliessen", func() {
			wSucc.Close()
		})
		closeButton.Importance = widget.HighImportance

		successContainer := container.New(layout.NewGridLayout(1), canvas.NewText("E-Mail wurde erfolgreich versendet.", color.Black), closeButton)

		wSucc.SetContent(successContainer)

		w.Close()
		wSucc.Show()
	})
	sendButton.Importance = widget.HighImportance

	w.SetContent(container.New(layout.NewGridLayout(1), dataContainer, sendButton))

	w.ShowAndRun()
}
