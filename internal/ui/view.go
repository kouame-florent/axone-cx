package ui

import (
	"fyne.io/fyne/v2"
)

type viewID uint

const (
	LIST_TICKETS_VIEW viewID = iota
	SEND_TICKET_VIEW
	CREDENTIALS_VIEW
)

type view struct {
	Win fyne.Window
}
