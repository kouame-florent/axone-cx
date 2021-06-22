package ui

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/kouame-florent/axone-cx/api/grpc/gen"
	"github.com/kouame-florent/axone-cx/internal/svc"
)

type AuthView struct {
	view
	loginEntry    widget.Entry
	passwordEntry widget.Entry

	//grpcConn *grpc.ClientConn
}

func NewAuth(app fyne.App, win fyne.Window) *AuthView {
	return &AuthView{
		view: view{
			Win: win,
		},
	}
}

func (s *AuthView) MakeUI() {

	s.loginEntry = widget.Entry{
		PlaceHolder: "Entrez votre identifieant",
	}

	s.passwordEntry = widget.Entry{
		PlaceHolder: "Entrez votre mot de passe",
		Password:    true,
	}

	dialog.ShowForm("Authentifcation", "Se connecter", "Abandonner",
		[]*widget.FormItem{
			widget.NewFormItem("Identifiant", &s.loginEntry),
			widget.NewFormItem("Mot de passe", &s.passwordEntry),
		},

		func(r bool) {
			if r {
				cli, conn, err := svc.Dial(s.loginEntry.Text, s.passwordEntry.Text)
				if err != nil {
					dialog.ShowError(err, s.Win)
					return
				}
				defer conn.Close()
				resp, err := cli.Login(context.Background(), &gen.LoginRequest{Username: "homer", Password: "homer"})
				if err != nil {
					dialog.ShowError(err, s.Win)
					return
				}
				setEnvVariables(resp)
				next, err := nextUI(SEND_TICKET_VIEW, s.Win)
				if err != nil {
					dialog.ShowError(err, s.Win)
					return
				}
				s.Win.SetContent(next)
			} else {
				s.Win.Close()
			}

		}, s.Win)

}

func setEnvVariables(resp *gen.LoginResponse) {
	log.Printf("AUTH: %s", resp.AuthToken)
	creds := strings.Split(resp.AuthToken, ";")
	if len(creds) == 2 {
		log.Printf("USER: %s", creds[0])
		log.Printf("PASSWD: %s", creds[1])
		os.Setenv("AXONE_USERNAME", creds[0])
		os.Setenv("AXONE_PASSWORD", creds[1])
	}
}

func nextUI(id viewID, win fyne.Window) (fyne.CanvasObject, error) {

	log.Printf("AXONE USERNAME: %s", os.Getenv("AXONE_USERNAME"))
	log.Printf("AXONE PASSWORD: %s", os.Getenv("AXONE_PASSWORD"))

	if id == SEND_TICKET_VIEW {
		sendTicket := NewSendTicket(win)

		return sendTicket.MakeUI(), nil
	}

	return nil, fmt.Errorf("%s", "cannot build 'SEND TICKET' view")

}
