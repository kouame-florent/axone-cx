package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type LeftMenuView struct {
	view
}

const (
	LEFT_MENU_ITEM_MES_REQUETES    = "Mes requetes"
	LEFT_MENU_ITEM_MON_COMPTE      = "Mon compte"
	LEFT_MENU_ITEM_MES_PREFERENCES = "Mes preferences"
)

func NewLeftMenuView(win fyne.Window) *LeftMenuView {
	return &LeftMenuView{
		view: view{
			Win: win,
		},
	}
}

var menuItems = []string{"Mes requetes", "Mon compte", "Mes preferences"}

func (l *LeftMenuView) MakeUI() fyne.CanvasObject {
	menu := widget.NewList(
		func() int {
			return len(menuItems)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("ticket numero 0001")
		},
		func(i int, template fyne.CanvasObject) {
			label := template.(*widget.Label)
			label.Text = menuItems[i]
		})

	menu.OnSelected = l.onSelect

	return menu
}

func (l *LeftMenuView) onSelect(id int) {
	view := menuItems[id]

	switch view {
	case LEFT_MENU_ITEM_MES_REQUETES:
		r := NewListTicketView(l.Win)
		r.Win.SetContent(r.MakeUI())
	case LEFT_MENU_ITEM_MON_COMPTE:
	}

}
