package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type toolbarView struct {
	view
}

func NewToolBarView(win fyne.Window) *toolbarView {
	return &toolbarView{
		view: view{
			Win: win,
		},
	}

}

func (t *toolbarView) MakeUI() fyne.CanvasObject {
	newTicketButton := widget.NewButtonWithIcon("Nouvelle requete", theme.DocumentCreateIcon(), func() {})
	//axoneLabel := widget.NewLabel("Axone")
	//axoneLabel.TextStyle = fyne.TextStyle{Bold: true, Italic: true}
	border := container.NewBorder(nil, nil, newTicketButton, nil)
	return border
}
