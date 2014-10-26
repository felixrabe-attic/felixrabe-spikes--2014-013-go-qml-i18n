package main

import (
	"log"

	"gopkg.in/qml.v1"
)

func main() {
	if err := qml.Run(run); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	engine := qml.NewEngine()
	com, err := engine.LoadFile("ui.qml")
	if err != nil {
		return err
	}
	win := com.CreateWindow(nil)
	win.Show()
	win.Wait()
	return nil
}
