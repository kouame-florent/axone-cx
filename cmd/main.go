package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/kouame-florent/axone-cx/api/grpc/gen"
	"github.com/kouame-florent/axone-cx/internal/ui"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {

	application := app.New()

	win := application.NewWindow("Axone")

	win.Resize(fyne.NewSize(1280, 480))

	grpcCli, conn := grpcClient()
	defer conn.Close()

	ticketUI := ui.NewTicket(grpcCli, win)
	win.SetContent(ticketUI.MakeUI())
	win.ShowAndRun()
}

func grpcClient() (gen.AxoneClient, *grpc.ClientConn) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return gen.NewAxoneClient(conn), conn

}
