package main

import "github.com/Binod210/gomongoCRUD/app"

func main() {
	app := &app.App{}
	app.Initialize()
	app.Run(":9090")
}
