package ui

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
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
	attachmentW  *widget.Button

	AttachmentURIs []fyne.URI

	ID uuid.UUID

	//subjectB     binding.ExternalString
	//requestB     binding.ExternalString
	//requestTypeB binding.ExternalString
}

func NewTicket(cli gen.AxoneClient, id uuid.UUID, w fyne.Window) *Ticket {
	return &Ticket{
		grpcClient: cli,
		Win:        w,
		ID:         id,
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
	ticketForm := widget.Form{
		//SubmitText: "Envoyer",
		//CancelText: "Annuler",
		Items: formItems,
	}

	sendFunc := func() {
		log.Print("sending ticket")
		//tID := uuid.New()

		requesterID := uuid.MustParse("4a2bfb72-94ab-4fb2-b195-52dc1a12ffdb")
		tType := TicketTypeMap[ticketTypeKey(requestTypeSelect.Selected)]

		ticketReq := &gen.NewTicketRequest{
			TicketID:    t.ID.String(),
			Subject:     subjectEntry.Text,
			Request:     requestEntry.Text,
			Type:        string(tType),
			RequesterID: requesterID.String(),
		}

		res, err := t.grpcClient.SendNewTicket(context.TODO(), ticketReq)
		if err != nil {
			dialog.ShowError(err, t.Win)
		}
		log.Printf("ticket id: %s", res.GetID())

		stream, err := t.grpcClient.SendAttachment(context.TODO())
		if err != nil {
			dialog.ShowError(err, t.Win)
		}
		sendAttachment(t.AttachmentURIs, t.ID.String(), t.Win, stream)

	}

	//var attachmentURI fyne.URI

	t.attachmentW = widget.NewButtonWithIcon("Ajouter une pièce jointe", theme.MailAttachmentIcon(), func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, t.Win)
				return
			}
			if reader == nil { // user cancelled
				return
			}

			t.AttachmentURIs = append(t.AttachmentURIs, reader.URI())
			log.Printf("Attachment uri: %s", t.AttachmentURIs)

		}, t.Win)
	})
	attchmentWContainer := container.NewBorder(nil, nil, t.attachmentW, nil)

	sendButton := widget.Button{
		Text:       "Envoyer",
		Importance: widget.HighImportance,
		OnTapped:   sendFunc,
		Icon:       theme.MailSendIcon(),
	}
	sendButtonContainer := container.NewBorder(nil, nil, nil, &sendButton)

	contentGrid := container.NewVBox(&ticketForm, attchmentWContainer, sendButtonContainer)

	return container.NewBorder(nil, nil, menu, nil, contentGrid)
}

func attachmentReqMeta(uri fyne.URI, ticketID string) (*gen.AttachmentRequest, error) {
	i, err := os.Stat(uri.Path())
	if err != nil {
		return &gen.AttachmentRequest{}, err
	}
	attReq := &gen.AttachmentRequest{
		Data: &gen.AttachmentRequest_Info{
			Info: &gen.AttachmentInfo{
				UploadedName: uri.Name(),
				MimeType:     uri.MimeType(),
				StorageName:  uuid.New().String(),
				Size:         uint32(i.Size()),
				TicketID:     ticketID,
			},
		},
	}
	return attReq, nil
}

func sendAttachment(uris []fyne.URI, ticketID string, win fyne.Window, stream gen.Axone_SendAttachmentClient) {
	for _, uri := range uris {
		sendAttacmentMeta(uri, ticketID, win, stream)
		sendAtttachmentChunk(uri, win, stream)
	}

}

func sendAttacmentMeta(uri fyne.URI, ticketID string, win fyne.Window, stream gen.Axone_SendAttachmentClient) {
	req, err := attachmentReqMeta(uri, ticketID)
	if err != nil {
		dialog.ShowError(err, win)
	}
	stream.Send(req)
}

func sendAtttachmentChunk(uri fyne.URI, win fyne.Window, stream gen.Axone_SendAttachmentClient) {
	reader, err := storage.Reader(uri)
	if err != nil {
		dialog.ShowError(err, win)
	}

	bufferedReader := bufio.NewReader(reader)
	buffer := make([]byte, 1024)

	for {
		n, err := bufferedReader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("cannot read chunk to buffer: ", err)
		}

		req := &gen.AttachmentRequest{
			Data: &gen.AttachmentRequest_ChunkData{
				ChunkData: buffer[:n],
			},
		}

		stream.Send(req)
		if err != nil {
			dialog.ShowError(err, win)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		dialog.ShowError(err, win)
	}
	log.Printf("file uploaded with id: %s", res.GetTicketID())

}
