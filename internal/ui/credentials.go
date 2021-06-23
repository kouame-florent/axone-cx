package ui

import (
	"context"
	"fmt"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/kouame-florent/axone-cx/api/grpc/gen"
	"github.com/kouame-florent/axone-cx/internal/axone"
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
				resp, err := cli.Login(context.Background(), &gen.LoginRequest{Login: "homer", Password: "homer"})
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
	os.Setenv(string(axone.ENVIRONMENT_VARIABLE_KEY_USER_ID), resp.UserID)
	os.Setenv(string(axone.ENVIRONMENT_VARIABLE_KEY_USER_LOGIN), resp.Login)
	os.Setenv(string(axone.ENVIRONMENT_VARIABLE_KEY_USER_PASSWORD), resp.Password)
	os.Setenv(string(axone.ENVIRONMENT_VARIABLE_KEY_USER_EMAIL), resp.Email)
	os.Setenv(string(axone.ENVIRONMENT_VARIABLE_KEY_USER_FIRST_NAME), resp.FirstName)
	os.Setenv(string(axone.ENVIRONMENT_VARIABLE_KEY_USER_LAST_NAME), resp.LastName)

}

func nextUI(id viewID, win fyne.Window) (fyne.CanvasObject, error) {

	log.Printf("AXONE LOGIN: %s", os.Getenv(string(axone.ENVIRONMENT_VARIABLE_KEY_USER_LOGIN)))
	log.Printf("AXONE PASSWORD: %s", os.Getenv(string(axone.ENVIRONMENT_VARIABLE_KEY_USER_PASSWORD)))
	log.Printf("AXONE USER ID: %s", os.Getenv(string(axone.ENVIRONMENT_VARIABLE_KEY_USER_ID)))

	if id == SEND_TICKET_VIEW {
		sendTicket := NewSendTicket(win)

		return sendTicket.MakeUI(), nil
	}

	return nil, fmt.Errorf("%s", "cannot build 'SEND TICKET' view")

}
