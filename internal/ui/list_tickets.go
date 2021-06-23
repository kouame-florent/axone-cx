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
	/*
		cli, conn, err := svc.DialWithEnvVariables()
		if err != nil {
			dialog.ShowError(err, l.Win)
		}
		defer conn.Close()

		cli.ListRequesterTickets(context.Background(),&gen.ListRequesterTicketsRequest{})
	*/

	tickets := widget.NewList(
		func() int {
			return 4
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("ticket numero 0001")
		},
		func(i int, template fyne.CanvasObject) {

		})

	toolBarView := NewToolBarView(l.Win)

	content := container.NewBorder(toolBarView.MakeUI(), nil, menu.MakeUI(), nil, tickets)

	return content

}
