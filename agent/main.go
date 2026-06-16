package main

import (
	"agent/internal/config"
	"context"
	"embed"
	"log"

	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

var appCtx context.Context
var appCfg *config.Config

func main() {
	appCfg = config.InitialLoad()
	systray.Run(onReady, onExit)
}

func onExit() {
	activeContexts := config.GetAllContexts()
	for _, ctx := range activeContexts {
		if ctx != nil {
			runtime.Quit(ctx)
		}
	}
}

func onReady() {
	systray.SetTitle(appCfg.App.Title)
	systray.SetTooltip(appCfg.App.Tooltip)
	for _, t := range appCfg.App.Tray {
		item := systray.AddMenuItem(t.Title, t.Tooltip)
		go func(item *systray.MenuItem, id int) {
			for {
				<-item.ClickedCh
				handleTrayClick(id)
			}
		}(item, t.FuncId)
	}
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quit the application")
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()
}

func handleTrayClick(id int) {
	ctx := config.GetContext(id)
	if ctx != nil {
		runtime.WindowShow(ctx)
		runtime.WindowCenter(ctx)
	} else {
		go CreateAppV2(id, NewApp())
	}
}

func Register(id int) string {
	pathMap := map[int]string{
		1: "/scan",
		2: "/update",
		3: "/settings",
	}
	if path, ok := pathMap[id]; ok {
		return path
	}
	return "/"
}

func CreateAppV2(id int, app *App) {
	path := Register(id)
	err := wails.Run(&options.App{
		Title:            appCfg.App.Title,
		Width:            appCfg.App.Dimension.Width,
		Height:           appCfg.App.Dimension.Height,
		StartHidden:      true,
		WindowStartState: 3,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: func(ctx context.Context) {
			config.SetContext(id, ctx)
			app.startup(ctx)
			runtime.EventsEmit(ctx, "navigate", path)
			runtime.WindowShow(ctx)
			runtime.WindowCenter(ctx)
		},
		OnBeforeClose: func(ctx context.Context) bool {
			runtime.WindowHide(ctx)
			return true // Prevent closing
		},
		Bind: []interface{}{app},
	})

	if err != nil {
		log.Printf("App instance %d failed: %v", id, err)
	}
}
