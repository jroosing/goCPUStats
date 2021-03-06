package main

import (
  "github.com/leaanthony/mewn"
  "github.com/wailsapp/wails"

  "cpustats/pkg/sys"
)

func main() {

  js := mewn.String("./frontend/dist/app.js")
  css := mewn.String("./frontend/dist/app.css")

  stats := &sys.Stats{}

  app := wails.CreateApp(&wails.AppConfig{
    Width:  1024,
    Height: 752,
    Resizable: true,
    DisableInspector: true,
    Title:  "CPU Usage",
    JS:     js,
    CSS:    css,
    Colour: "#131313",
  })
  app.Bind(stats)
  app.Run()
}
