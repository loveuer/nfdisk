package main

import (
	"embed"
	"fmt"
	"github.com/sirupsen/logrus"
	"nfdisk/internal/controller"
	"nfdisk/internal/model"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	logrus.SetLevel(logrus.TraceLevel)

	model.Init()

	app := controller.NewApp()

	if err := wails.Run(&options.App{
		Title:  "💾 NFDisk",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
			//Handler: handler.NewHandler(),
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 1},
		OnStartup:        app.Startup,
		Bind: []interface{}{
			app,
		},
	}); err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
}
