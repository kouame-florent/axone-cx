package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/kouame-florent/axone-cx/internal/ui"
)

func main() {

	app := app.NewWithID("axone-cx")

	win := app.NewWindow("Axone")
	win.Resize(fyne.NewSize(1280, 480))

	//grpcCli, conn := svc.GrpcClient()

	auth := ui.NewAuth(app, win)
	auth.MakeUI()

	//sendTicketUI, err := ui.NewSendTicket(app, win)
	//sendTicketCanvasObject := sendTicketUI.MakeUI()

	win.ShowAndRun()
}
