package ui

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/google/uuid"
	"github.com/kouame-florent/axone-cx/api/grpc/gen"
	"github.com/kouame-florent/axone-cx/internal/axone"
	"github.com/kouame-florent/axone-cx/internal/svc"
)

type SendTicketView struct {
	view

	sendTicketForm      widget.Form
	subjectFormItem     *widget.FormItem
	subjectEntry        *widget.Entry
	requestFormItem     *widget.FormItem
	requestEntry        *widget.Entry
	requestTypeFormItem *widget.FormItem
	requestTypeSelect   *widget.Select
	sendButton          *widget.Button
	attachmentButton    *widget.Button

	attachmentVbox *fyne.Container

	AttachmentURIs []fyne.URI

	subjectBinding       binding.ExternalString
	subjectBindingString *string
	requestBinding       binding.ExternalString
	requestBindingString *string
}

func NewSendTicket(win fyne.Window) *SendTicketView {
	return &SendTicketView{
		view: view{
			Win: win,
		},
	}
}

type ticketTypeKey string

const (
	TICKET_TYPE_KEY_QUESTION ticketTypeKey = "Question"
	TICKET_TYPE__KEY_PROBLEM ticketTypeKey = "Problème"
	TICKET_TYPE_KEY_TASK     ticketTypeKey = "Tâche"
)

var TicketTypeMap = map[ticketTypeKey]axone.TicketType{
	TICKET_TYPE_KEY_QUESTION: axone.TICKET_TYPE_QUESTION,
	TICKET_TYPE__KEY_PROBLEM: axone.TICKET_TYPE_PROBLEM,
	TICKET_TYPE_KEY_TASK:     axone.TICKET_TYPE_TASK,
}

func (st *SendTicketView) MakeUI() fyne.CanvasObject {

	menu := NewLeftMenuView(st.Win)

	st.subjectBinding = binding.BindString(st.subjectBindingString)
	st.subjectEntry = widget.NewEntryWithData(st.subjectBinding)
	st.subjectEntry.PlaceHolder = "Sujet de la Requête"
	st.subjectFormItem = widget.NewFormItem("", st.subjectEntry)

	//st.requestTypeBinding = binding.BindString(st.requestTypeBindingString)
	st.requestTypeSelect = widget.NewSelect([]string{string(TICKET_TYPE_KEY_QUESTION), string(TICKET_TYPE__KEY_PROBLEM),
		string(TICKET_TYPE_KEY_TASK)}, func(s string) {})
	st.requestTypeSelect.PlaceHolder = "Type de la Requête"
	st.requestTypeFormItem = widget.NewFormItem("", st.requestTypeSelect)

	st.requestBinding = binding.BindString(st.requestBindingString)
	st.requestEntry = widget.NewEntryWithData(st.requestBinding)
	st.requestEntry.MultiLine = true
	st.requestEntry.PlaceHolder = "Description de la Requête"
	st.requestFormItem = widget.NewFormItem("", st.requestEntry)

	formItems := []*widget.FormItem{
		st.subjectFormItem,
		st.requestTypeFormItem,
		st.requestFormItem,
	}
	st.sendTicketForm = widget.Form{
		Items: formItems,
	}

	st.attachmentVbox = container.NewVBox()

	st.attachmentButton = widget.NewButtonWithIcon("Ajouter une pièce jointe", theme.MailAttachmentIcon(), func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, st.Win)
				return
			}
			if reader == nil { // user cancelled
				return
			}

			st.AttachmentURIs = append(st.AttachmentURIs, reader.URI())
			log.Printf("Attachment uri: %s", st.AttachmentURIs)

			button := widget.NewButtonWithIcon("", theme.CancelIcon(), func() {})
			label := widget.NewLabel(reader.URI().Name())
			hbox := container.NewHBox(label, button)

			//remove attachment when x button is tapped
			button.OnTapped = func() {
				st.attachmentVbox.Remove(hbox)
				st.attachmentVbox.Refresh()
			}

			st.attachmentVbox.Add(hbox)
			st.attachmentVbox.Refresh()

		}, st.Win)
	})

	attchmentBorder := container.NewBorder(nil, nil, st.attachmentButton, nil)

	st.sendButton = &widget.Button{
		Text:       "Envoyer",
		Importance: widget.HighImportance,
		OnTapped:   sendCallBack(st),
		Icon:       theme.MailSendIcon(),
	}
	sendButtonBorder := container.NewBorder(nil, nil, nil, st.sendButton)

	contentVBox := container.NewVBox(&st.sendTicketForm, attchmentBorder, st.attachmentVbox, sendButtonBorder)

	return container.NewBorder(nil, nil, menu.MakeUI(), nil, contentVBox)
}

func (st *SendTicketView) reset() {
	st.subjectBinding.Set("")
	st.requestBinding.Set("")
	st.requestTypeSelect.Selected = ""
	st.requestTypeSelect.Refresh()
	st.AttachmentURIs = []fyne.URI{}
	st.attachmentVbox.Refresh()
}

//callback to send tickets
func sendCallBack(st *SendTicketView) func() {
	log.Print("sending ticket")

	f := func() {
		cli, conn, err := svc.Dial(os.Getenv(string(axone.ENVIRONMENT_VARIABLE_KEY_USER_LOGIN)),
			os.Getenv(string(axone.ENVIRONMENT_VARIABLE_KEY_USER_PASSWORD)))
		if err != nil {
			dialog.ShowError(err, st.Win)
			return
		}
		defer conn.Close()

		requesterID := uuid.MustParse(string(axone.ENVIRONMENT_VARIABLE_KEY_USER_ID))
		tType := TicketTypeMap[ticketTypeKey(st.requestTypeSelect.Selected)]

		ticketID := uuid.New().String()

		res, err := sendNewTicket(ticketID, st.subjectEntry.Text, st.requestEntry.Text, string(tType), requesterID.String(), cli)
		if err != nil {
			dialog.ShowError(err, st.Win)
			return
		}
		log.Printf("ticket id: %s", res.GetID())

		for _, uri := range st.AttachmentURIs {
			//get client stream
			stream, err := cli.SendAttachment(context.TODO())
			if err != nil {
				dialog.ShowError(err, st.Win)
				return
			}
			sendAttachmentMeta(uri, ticketID, st.Win, stream)
			sendAtttachmentChunk(uri, st.Win, stream)
		}

		st.reset()
		sendNotification(fyne.CurrentApp(), ticketID)

	}
	return f
}

func sendAttachmentMeta(uri fyne.URI, ticketID string, win fyne.Window, stream gen.Axone_SendAttachmentClient) {
	log.Printf("Send meta for attachment: %s", uri.Name())
	req, err := attachmentMeta(uri, ticketID)
	if err != nil {
		dialog.ShowError(err, win)
	}
	stream.Send(req)
}

func sendAtttachmentChunk(uri fyne.URI, win fyne.Window, stream gen.Axone_SendAttachmentClient) {
	log.Printf("Send chunck for attachment: %s", uri.Name())
	reader, err := storage.Reader(uri)
	if err != nil {
		dialog.ShowError(err, win)
	}

	buffSize := 1 << 20
	bufferedReader := bufio.NewReader(reader)
	buffer := make([]byte, buffSize)

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

func sendNotification(app fyne.App, ticketID string) *fyne.Notification {
	notif := fyne.NewNotification("Axone", fmt.Sprintf("Msg: %s sent", ticketID))
	app.SendNotification(notif)
	return notif

}

func sendNewTicket(ticketID, subject, request, ticketType, requesterID string, cli gen.AxoneClient) (*gen.NewTicketResponse, error) {
	req := &gen.NewTicketRequest{
		TicketID:    ticketID,
		Subject:     subject,
		Request:     request,
		Type:        string(ticketType),
		RequesterID: requesterID,
	}

	res, err := cli.SendNewTicket(context.TODO(), req)
	if err != nil {
		return &gen.NewTicketResponse{}, err
	}
	return res, nil

}

func attachmentMeta(uri fyne.URI, ticketID string) (*gen.AttachmentRequest, error) {
	i, err := os.Stat(uri.Path())
	if err != nil {
		return &gen.AttachmentRequest{}, err
	}
	attReq := &gen.AttachmentRequest{
		Data: &gen.AttachmentRequest_Info{
			Info: &gen.AttachmentInfo{
				UploadedName: uri.Name(),
				MimeType:     uri.MimeType(),
				Size:         uint32(i.Size()),
				TicketID:     ticketID,
			},
		},
	}
	return attReq, nil
}
