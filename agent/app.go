package main

import (
	"agent/internal/config"
	"context"
	"embed"
	"fmt"
	"log"

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

	a.runBackgroundWatcher(ctx)
}

func notifyUser(ctx context.Context, title, body string) {
	if runtime.IsNotificationAvailable(ctx) {
		err := runtime.SendNotification(ctx, runtime.NotificationOptions{
			ID:    "tray-notification",
			Title: title,
			Body:  body,
		})
		if err != nil {
			log.Printf("Failed to send notification: %v", err)
		}
	}
}

func (a *App) runBackgroundWatcher(ctx context.Context) {
	// Whenever you need to notify:
	SendAgentNotification(ctx, "Alert", "Service is running in background!")
}

func SendAgentNotification(ctx context.Context, title string, message string) {
	// Check if notifications are supported on the current platform
	if runtime.IsNotificationAvailable(ctx) {
		err := runtime.SendNotification(ctx, runtime.NotificationOptions{
			ID:    "agent-notification-001", // Use a unique ID
			Title: title,
			Body:  message,
		})
		if err != nil {
			// Log error if notification fails
		}
	}
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

func RunApp(id int, title string, hide bool) error {
	app := NewApp()
	path := Register(id)
	js := fmt.Sprintf("window.location.hash = '%s';", path)
	err := wails.Run(&options.App{
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
	return err
}
