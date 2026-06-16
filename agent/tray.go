package main

import (
	"agent/internal/config"
	_ "embed"
	"log"
	"strconv"

	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var appCfg *config.AppConfig

//go:embed internal/assets/fav.ico
var icon []byte

func InitializeTray(cfg *config.AppConfig) {

	appCfg = cfg
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(icon)
	systray.SetTitle(appCfg.Title)
	systray.SetTooltip(appCfg.Tooltip)
	PopulateMenu()
}

func PopulateMenu() {
	for _, t := range appCfg.Tray {
		if t.Sperator == true {
			systray.AddSeparator()
		} else {
			item := systray.AddMenuItem(t.Title, t.Tooltip)
			go func(item *systray.MenuItem, id int) {
				for {
					<-item.ClickedCh
					handleTrayClick(t.FuncId, t.Title)
				}
			}(item, t.FuncId)
		}
	}
}

func onExit() {
	log.Println("Exiting application...")
	contexts := config.GetAllContexts()
	for _, ctx := range contexts {
		if ctx != nil {
			runtime.Quit(ctx)
		}
	}
}

func handleTrayClick(id int, title string) {
	if id == 100 {
		log.Println("Quit Action Triggerd")
		systray.Quit()
		return
	}

	ctx := config.GetContext(id)
	if ctx != nil {
		notifyUser(ctx, "Agent Action", "Performing action ID: "+strconv.Itoa(id))
		runtime.WindowShow(ctx)
		runtime.WindowCenter(ctx)
	} else {
		// Run in a new goroutine so the tray loop doesn't hang
		go func(windowID int) {
			log.Printf("Launching new window for ID: %d", windowID)
			RunApp(windowID, title, false)
		}(id)
	}
}

func Register(id int) string {
	pathMap := map[int]string{
		0: "/",
		1: "/scan",
		2: "/conn",
	}
	if path, ok := pathMap[id]; ok {
		return path
	}
	return "/"
}
