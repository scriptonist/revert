package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/scriptonist/gowebapp/daemon"
)

var assetsPath string

func processFlags() *daemon.Config {
	cfg := &daemon.Config{}
	flag.StringVar(&cfg.ListenSpec, "listen", "localhost:3000", "HTTP listen spec")
	flag.StringVar(&cfg.Db.ConnectionString, "db-connect", "mongodb://gowebapp:gowebapp@localhost/gowebapp", "Mongo Connection String")
	flag.StringVar(&cfg.Db.DBName, "db-name", "gowebapp", "Database Name")
	flag.StringVar(&assetsPath, "assets-path", "assets", "path to assets dir")

	flag.Parse()

	return cfg
}

func setupHttpAssets(cfg *daemon.Config) {
	log.Printf("Assets served from %q", assetsPath)
	cfg.Ui.Assets = http.Dir(assetsPath)
}
func main() {
	cfg := processFlags()
	setupHttpAssets(cfg)
	if err := daemon.Run(*cfg); err != nil {
		log.Printf("Error in main!")
	}
}
