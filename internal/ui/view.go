package ui

import (
	"fyne.io/fyne/v2"
)

type viewID uint

const (
	SEND_TICKET viewID = iota
	LIST_TICKETS
	SNED_CREDENTIALS
)

type view struct {
	//grpcClient gen.AxoneClient
	Win       fyne.Window
	App       fyne.App
	Notif     *fyne.Notification
	AuthToken string
}
