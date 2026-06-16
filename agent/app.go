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
var assets embed.FS

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
func (a *App) startup(ctx context.Context, id int) {
	a.ctx = ctx
	config.SetContext(id, ctx)
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

//go:embed all:frontend/dist/*
var Assets embed.FS

var cfg *config.AppConfig

// 2. Store the passed-in values into the package variables
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

func RunApp(id int, hide bool) error {
	app := NewApp()
	path := Register(id)
	js := fmt.Sprintf("window.location.hash = '%s';", path)
	err := wails.Run(&options.App{
		Title:       fmt.Sprintf("%s - %d", cfg.Title, id),
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
	return err
}
