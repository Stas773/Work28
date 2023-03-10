package main

import (
	"D/Works/GO/27work/pkg/app"
	"D/Works/GO/27work/pkg/memstorage"
)

func main() {
	repository := memstorage.NewMemStore()
	app := &app.App{Repository: repository}
	app.Run()
}
