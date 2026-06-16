package main

import (
	"agent/internal/config"
	"log"
	"os"
)

func setupLogger() {
	// Open (or create) the log file.
	// O_APPEND: append to file, O_CREATE: create if not exists, O_WRONLY: write-only
	f, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	// Set log output to the file
	log.SetOutput(f)
	// Optional: add flags for date and time
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	for _, arg := range os.Args {
		if arg == "-binding-check" {
			os.Exit(0)
		}
	}
	setupLogger()

	log.Println("Engine Starting...")
	cfg := config.InitialLoad()

	log.Println("Initializing App...")
	// Start App in a goroutine
	go InitializeApp(&cfg.App)

	log.Println("Initializing System Tray...")
	InitializeTray(&cfg.App)
}
