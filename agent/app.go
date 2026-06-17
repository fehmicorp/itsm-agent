package main

import (
	"agent/internal/config"
	"context"
	"embed"
	"fmt"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist/*
var Assets embed.FS

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context, id int) {
	a.ctx = ctx
	config.SetContext(id, ctx)
	a.runBackgroundWatcher(ctx)
}

func (a *App) runBackgroundWatcher(ctx context.Context) {
	config.SetContext(5001, ctx)
	TriggerNotification(ctx, "Alert", "Service is running in background!")
}

func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

var cfg *config.AppConfig

func InitializeApp(appCfg *config.AppConfig) {
	cfg = appCfg
}

func QuitApp(id int) {
	ctx := config.GetContext(id)
	if ctx != nil {
		config.DeleteContext(id)
		runtime.Quit(ctx)
	}
}

func RunApp(id int, title string, hide bool) error {
	app := NewApp()
	path := Register(id)
	js := fmt.Sprintf("window.location.hash = '%s';", path)

	return wails.Run(&options.App{
		Title:       fmt.Sprintf("%s - %s", cfg.Title, title),
		Width:       cfg.Dimension.Width,
		Height:      cfg.Dimension.Height,
		StartHidden: hide,
		AssetServer: &assetserver.Options{Assets: Assets},
		OnStartup: func(ctx context.Context) {
			app.startup(ctx, id)
			runtime.WindowExecJS(ctx, js)
		},
		OnBeforeClose: func(ctx context.Context) bool {
			QuitApp(id)
			return false
		},
		Bind: []interface{}{app},
	})
}
