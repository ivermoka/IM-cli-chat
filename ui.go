package main

import (
	"github.com/awesome-gocui/gocui"
)

func quit(*gocui.Gui, *gocui.View) error {
	return gocui.ErrQuit
}

// selve oppsettet til hvordan GUIet ser ut
func layout(gui *gocui.Gui) error {
	maxX, maxY := gui.Size()

	if v, err := gui.SetView("messages", 0, 0, maxX-1, maxY-3, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Messages"
		v.Wrap = true
		v.Autoscroll = true

		// kanskje ta bort ???
		v.FgColor = gocui.ColorGreen 
		v.BgColor = gocui.ColorDefault 
	}

	if v, err := gui.SetView("input", 0, maxY-3, maxX-1, maxY-1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Input"
		v.Editable = true
		v.Wrap = true
		if _, err := gui.SetCurrentView("input"); err != nil {
			return err
		}
	}

	return nil
}