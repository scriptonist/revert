package daemon

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/scriptonist/gowebapp/model"

	"github.com/scriptonist/gowebapp/db"
	"github.com/scriptonist/gowebapp/ui"
)

type Config struct {
	ListenSpec string
	Db         db.Config
	Ui         ui.Config
}

func Run(cfg Config) error {
	log.Printf("Starting HTTP Server on : %s\n", cfg.ListenSpec)
	db, err := db.InitDB(cfg.Db)
	if err != nil {
		log.Printf("Error initializing database - %v", err)
	}

	m := model.New(db)

	l, err := net.Listen("tcp", cfg.ListenSpec)
	if err != nil {
		log.Printf("Error creating listener ! %v\n", err)
		return err
	}

	ui.Start(cfg.Ui, m, l)

	waitForSignal()
	return nil

}

func waitForSignal() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	s := <-ch
	log.Printf("Got Signal: %v, Exiting", s)
}
