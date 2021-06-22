package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type ListTicketView struct {
	view
}

func NewListTicketView(win fyne.Window) *ListTicketView {
	return &ListTicketView{
		view: view{
			Win: win,
		},
	}
}

func (l *ListTicketView) MakeUI() fyne.CanvasObject {
	menu := NewLeftMenuView(l.Win)

	label := widget.NewLabel("list of tickets")

	return container.NewBorder(nil, nil, menu.MakeUI(), nil, label)

}
