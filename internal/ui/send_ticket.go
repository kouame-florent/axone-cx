package ui

import (
	"context"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/google/uuid"
	"github.com/kouame-florent/axone-cx/api/grpc/gen"
)

type Ticket struct {
	grpcClient gen.AxoneClient
	Win        fyne.Window
	//formW        *widget.Form
	subjectW     *widget.FormItem
	requestW     *widget.FormItem
	requestTypeW *widget.FormItem

	//subjectB     binding.ExternalString
	//requestB     binding.ExternalString
	//requestTypeB binding.ExternalString
}

func NewTicket(cli gen.AxoneClient, w fyne.Window) *Ticket {
	return &Ticket{
		grpcClient: cli,
		Win:        w,
	}
}

type ticketTypeKey string

const (
	TICKET_TYPE_KEY_QUESTION ticketTypeKey = "Question"
	TICKET_TYPE__KEY_PROBLEM ticketTypeKey = "Problème"
	TICKET_TYPE_KEY_TASK     ticketTypeKey = "Tâche"
)

type TicketType string

const (
	TICKET_TYPE_QUESTION TicketType = "question"
	TICKET_TYPE_PROBLEM  TicketType = "problem"
	TICKET_TYPE_TASK     TicketType = "task"
)

var TicketTypeMap = map[ticketTypeKey]TicketType{
	TICKET_TYPE_KEY_QUESTION: TICKET_TYPE_QUESTION,
	TICKET_TYPE__KEY_PROBLEM: TICKET_TYPE_PROBLEM,
	TICKET_TYPE_KEY_TASK:     TICKET_TYPE_TASK,
}

func (t *Ticket) MakeUI() fyne.CanvasObject {

	menu := widget.NewList(
		func() int {
			return 2
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("ticket numero 0001")
		},
		func(i int, c fyne.CanvasObject) {})

	subjectEntry := widget.NewEntry()
	t.subjectW = widget.NewFormItem("Sujet", subjectEntry)
	requestTypeSelect := widget.NewSelect([]string{string(TICKET_TYPE_KEY_QUESTION), string(TICKET_TYPE__KEY_PROBLEM),
		string(TICKET_TYPE_KEY_TASK)}, func(s string) {})
	t.requestTypeW = widget.NewFormItem("Type", requestTypeSelect)
	requestEntry := widget.NewMultiLineEntry()
	t.requestW = widget.NewFormItem("Requête", requestEntry)

	formItems := []*widget.FormItem{
		t.subjectW,
		t.requestTypeW,
		t.requestW,
	}
	TicketForm := widget.Form{
		SubmitText: "Envoyer",
		CancelText: "Annuler",
		Items:      formItems,
	}

	//newResp := &gen.NewTicketResponse{}
	TicketForm.OnSubmit = func() {
		log.Print("sending ticket")
		tID := uuid.New()

		requesterID := uuid.MustParse("4a2bfb72-94ab-4fb2-b195-52dc1a12ffdb")
		tType := TicketTypeMap[ticketTypeKey(requestTypeSelect.Selected)]

		req := &gen.NewTicketRequest{
			TicketID:    tID.String(),
			Subject:     subjectEntry.Text,
			Request:     requestEntry.Text,
			Type:        string(tType),
			RequesterID: requesterID.String(),
		}

		res, err := t.grpcClient.SendNewTicket(context.TODO(), req)
		if err != nil {
			dialog.ShowError(err, t.Win)
		}
		log.Printf("ticket id: %s", res.GetID())

	}

	content := container.NewMax(&TicketForm)

	return container.NewBorder(nil, nil, menu, nil, content)
}
