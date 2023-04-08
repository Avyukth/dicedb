package main

import (
	"flag"
	"log"
	"github.com/Avyukth/dicedb/config"
	"github.com/Avyukth/dicedb/server"
)

func setupFlags(){
	flag.StringVar(&config.Host, "host", "0.0.0.0", "Host for the dice server")
	flag.StringVar(&config.Port, "port", 7379, "port for the dice server")
	flag.Parse()
}

func main() {
	setupFlags()
	log.Println("Starting dice server ")
	server.RunSyncTCPServer()
}
