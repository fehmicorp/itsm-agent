package tray

import (
	"agent/internal/config"

	"github.com/getlantern/systray"
)

// Global variables to hold the data
var (
	icon   []byte
	appCfg *config.AppConfig
)

func InitializeTray(cfg *config.AppConfig, iconData []byte) {
	appCfg = cfg
	icon = iconData

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
			if t.FuncId == 100 {
				go func() {
					<-item.ClickedCh
					systray.Quit()
				}()
			} else {
				go func(item *systray.MenuItem, id int) {
					for {
						<-item.ClickedCh
						println(t.Title)
					}
				}(item, t.FuncId)
			}
		}
	}
}

func onExit() {

}
